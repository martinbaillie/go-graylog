package main

import (
	"net/url"
	"strings"
	"time"

	"github.com/martinbaillie/go-graylog/cmd/graylog/logging"
	"github.com/martinbaillie/go-graylog/cmd/graylog/search"
	"github.com/martinbaillie/go-graylog/pkg/client"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"

	clientrun "github.com/go-openapi/runtime/client"

	clientabs "github.com/martinbaillie/go-graylog/pkg/client/search_decorators"
	clientrel "github.com/martinbaillie/go-graylog/pkg/client/search_saved"
	clientkey "github.com/martinbaillie/go-graylog/pkg/client/search_universal_relative"
)

const (
	errSearchFailed = "failed to perform search"
)

var (
	// Search command flags
	utc        bool
	timestamps bool
	follow     bool
	from       string
	to         string
	keyword    string
	interval   time.Duration
	since      time.Duration

	// Search command defaults
	defaultSort  = "timestamp:asc"
	defaultLimit int64

	// Output logger
	lf = logging.NewColouredLoggerFactory()
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search related commands",
}

var absCmd = &cobra.Command{
	Use:   "absolute",
	Short: "Search for messages using an absolute timerange",
	Long: `Search for messages using an absolute timerange, specified as from/to
with format yyyy-MM-ddTHH:mm:ss.SSSZ (e.g. 2014-01-23T15:34:49.000Z) or
yyyy-MM-dd HH:mm:ss.`,
	Args: mandatoryArgs([]string{"from", "to"}, 1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// Prepare search.
		var (
			graylogs = urlsToSearchableGraylogs(serverURLs)
			auth     = clientrun.BasicAuth(user, pass)
			c        = make(chan logging.LoggableMessages, len(graylogs))
			errc     = make(chan error)
		)
		defer close(c)
		defer close(errc)

		// Concurrently search specified Graylog(s).
		for _, graylog := range graylogs {
			params := clientabs.NewSearchAbsoluteParams().WithSort(&defaultSort).WithLimit(&defaultLimit)
			params.Query = strings.Join(args, " ")
			params.From = from
			params.To = to

			log.WithField("graylog", graylog).Debug("Performing absolute search")
			go search.ByAbsolute(c, errc, graylog, params, auth)
		}

		// Collate result messages.
		var messages logging.LoggableMessages
		messages, err = collateSearchResultMessages(c, errc, graylogs)

		// Sort result messages by timestamp and log them.
		for _, lm := range messages.Sort() {
			lm.Log(timestamps, utc)
		}

		return err
	},
}

var keyCmd = &cobra.Command{
	Use:   "keyword",
	Short: "Search for messages using a natural language timerange",
	Long: `Search for messages using a natural language timerange e.g.
"yesterday" or "2 weeks ago to wednesday".`,
	Args: mandatoryArgs([]string{"keyword"}, 1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// Prepare search.
		var (
			graylogs = urlsToSearchableGraylogs(serverURLs)
			auth     = clientrun.BasicAuth(user, pass)
			c        = make(chan logging.LoggableMessages, len(graylogs))
			errc     = make(chan error)
		)
		defer close(c)
		defer close(errc)

		// Concurrently search specified Graylog(s).
		for _, graylog := range graylogs {
			params := clientkey.NewSearchKeywordParams().WithSort(&defaultSort).WithLimit(&defaultLimit)
			params.Query = strings.Join(args, " ")
			params.Keyword = keyword

			log.WithField("graylog", graylog).Debug("Performing keyword search")
			go search.ByKeyword(c, errc, graylog, params, auth)
		}

		// Collate result messages.
		var messages logging.LoggableMessages
		messages, err = collateSearchResultMessages(c, errc, graylogs)

		// Sort result messages by timestamp and log them.
		for _, lm := range messages.Sort() {
			lm.Log(timestamps, utc)
		}

		return err
	},
}

var relCmd = &cobra.Command{
	Use:   "relative",
	Short: "Search for messages starting from a relative timestamp",
	Long: `Search for messages starting from a relative timestamp until now.
	
The timestamp must be a valid Golang duration e.g. "2h".

Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".`,
	Args: mandatoryArgs(nil, 1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// Prepare search.
		var (
			graylogs = urlsToSearchableGraylogs(serverURLs)
			auth     = clientrun.BasicAuth(user, pass)
			params   = make([]*clientrel.SearchRelativeParams, len(graylogs))
			c        = make(chan logging.LoggableMessages, len(graylogs))
			errc     = make(chan error)
		)
		defer close(c)
		defer close(errc)

		// Concurrently search specified Graylog(s).
		for i, graylog := range graylogs {
			params[i] = clientrel.NewSearchRelativeParams().WithSort(&defaultSort).WithLimit(&defaultLimit)
			params[i].Query = strings.Join(args, " ")

			// Relative to program start minus specified `since` duration (if any).
			graylog.LastCalled = time.Unix(time.Now().Unix()-int64(since.Seconds()), 0)

			log.WithField("graylog", graylog).Debug("Performing relative search")
			go search.ByRelative(c, errc, graylog, params[i], auth)
		}

		// Collate result messages.
		var messages logging.LoggableMessages
		messages, err = collateSearchResultMessages(c, errc, graylogs)

		// Sort result messages by timestamp and log them.
		for _, lm := range messages.Sort() {
			lm.Log(timestamps, utc)
		}

		if follow {
			for i, graylog := range graylogs {
				log.WithFields(log.Fields{
					"graylog":  graylog,
					"interval": interval,
				}).Debug("Repeating relative search on an interval")
				go search.ByRelativeWithPoll(interval, c, errc, graylog, params[i], auth)
			}
			for {
				select {
				case err = <-errc:
					log.WithError(err).Error(errSearchFailed)
				case m := <-c:
					for _, lm := range m {
						lm.Log(timestamps, utc)
					}
				}
			}
		}

		return err
	},
}

func init() {
	relCmd.Flags().DurationVar(
		&since,
		"since",
		0,
		"from this relative Golang timestamp until now e.g. \"2h\"",
	)
	relCmd.Flags().BoolVarP(
		&follow,
		"follow",
		"f",
		false,
		"follow mode (tail-like functionality)",
	)
	relCmd.Flags().DurationVarP(
		&interval,
		"interval",
		"i",
		5*time.Second,
		"follow mode polling interval e.g. \"5m\"",
	)

	absCmd.Flags().StringVar(
		&from,
		"from",
		"",
		"from this absolute timestamp e.g. \"2018-01-01 00:00:00\"",
	)
	absCmd.Flags().StringVar(
		&to,
		"to",
		"",
		"to this absolute timestamp e.g. \"2018-02-01 00:00:00\"",
	)

	keyCmd.Flags().StringVar(
		&keyword,
		"keyword",
		"",
		"natural language timerange e.g. \"yesterday\" or \"2 weeks ago to wednesday\"",
	)

	searchCmd.PersistentFlags().BoolVarP(
		&timestamps,
		"timestamps",
		"t",
		false,
		"print message timestamps (local timezone)",
	)
	searchCmd.PersistentFlags().BoolVar(
		&utc,
		"utc",
		false,
		"print UTC message timestamps (has no effect without `-t`)",
	)
	searchCmd.AddCommand(absCmd)
	searchCmd.AddCommand(relCmd)
	searchCmd.AddCommand(keyCmd)

	graylogCmd.AddCommand(searchCmd)
}

// urlsToSearchableGraylogs transforms a slice of CLI-provided URLs into
// Graylog structs ready to be used by the search package.
func urlsToSearchableGraylogs(urls []*url.URL) []*search.Graylog {
	var graylogs = make([]*search.Graylog, len(urls))

	for i, url := range urls {
		graylogs[i] = &search.Graylog{
			Logger: lf.Create(url.String()),
			Client: client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
				Host:     strings.Join([]string{url.Hostname(), url.Port()}, ":"),
				Schemes:  []string{url.Scheme},
				BasePath: url.Path,
			}),
		}
	}

	return graylogs
}

// collateSearchResultMessages listens on the channels for messages and errors
// for each of the provided Graylogs. It will merge concurrent results together
// before returning the combined slice.
func collateSearchResultMessages(
	c chan logging.LoggableMessages,
	errc chan error,
	gs []*search.Graylog,
) (messages logging.LoggableMessages, err error) {
	for range gs {
		select {
		case err = <-errc:
			log.WithError(err).Error(errSearchFailed)
		case m := <-c:
			messages = append(messages, m...)
		}
	}
	return
}
