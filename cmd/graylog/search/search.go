package search

import (
	"time"

	runtime "github.com/go-openapi/runtime"
	client "github.com/martinbaillie/go-graylog/pkg/client"
	clientabs "github.com/martinbaillie/go-graylog/pkg/client/search_decorators"
	clientrel "github.com/martinbaillie/go-graylog/pkg/client/search_saved"
	clientkey "github.com/martinbaillie/go-graylog/pkg/client/search_universal_relative"

	"github.com/martinbaillie/go-graylog/cmd/graylog/logging"
	"github.com/martinbaillie/go-graylog/pkg/models"
)

// Graylog models required fields for dealing with the Graylog search APIs.
type Graylog struct {
	Client *client.Graylog
	Logger logging.Logger

	// For relative search follow mode.
	LastCalled time.Time
}

func (g *Graylog) String() string {
	if g.Logger != nil {
		return g.Logger.String()
	}
	return ""
}

// ByAbsolute searches the provided Graylog using the keyword API and the
// provided parameters and auth. Messages are sent on the provided channel.
// nolint: dupl
func ByAbsolute(
	ch chan logging.LoggableMessages,
	errc chan error,
	g *Graylog,
	params *clientabs.SearchAbsoluteParams,
	auth runtime.ClientAuthInfoWriter) {

	if res, err := g.Client.SearchDecorators.SearchAbsolute(params, auth); err != nil {
		errc <- err
	} else {
		ch <- searchResponseToLoggableMsgs(res.Payload.Messages, g.Logger)
	}
}

// ByKeyword searches the provided Graylog using the keyword API and the
// provided parameters and auth. Messages are sent on the provided channel.
// nolint: dupl
func ByKeyword(
	ch chan logging.LoggableMessages,
	errc chan error,
	g *Graylog,
	params *clientkey.SearchKeywordParams,
	auth runtime.ClientAuthInfoWriter) {

	if res, err := g.Client.SearchUniversalRelative.SearchKeyword(params, auth); err != nil {
		errc <- err
	} else {
		ch <- searchResponseToLoggableMsgs(res.Payload.Messages, g.Logger)
	}
}

// ByRelative searches the provided Graylog using the relative API and the
// provided parameters and auth. Messages are sent on the provided channel.
func ByRelative(
	ch chan logging.LoggableMessages,
	errc chan error,
	g *Graylog,
	params *clientrel.SearchRelativeParams,
	auth runtime.ClientAuthInfoWriter) {

	// Update the relative range.
	now := time.Now()
	params.Range = int64(now.Sub(g.LastCalled).Seconds())
	g.LastCalled = now

	if res, err := g.Client.SearchSaved.SearchRelative(params, auth); err != nil {
		errc <- err
	} else {
		ch <- searchResponseToLoggableMsgs(res.Payload.Messages, g.Logger)
	}
}

// ByRelativeWithPoll performs a relative search every `interval`.
func ByRelativeWithPoll(
	interval time.Duration,
	ch chan logging.LoggableMessages,
	errc chan error,
	g *Graylog,
	params *clientrel.SearchRelativeParams,
	auth runtime.ClientAuthInfoWriter) {

	// Perform a relative search every interval.
	for range time.NewTicker(interval).C {
		ByRelative(ch, errc, g, params, auth)
	}
}

// searchResponseToLoggableMsgs converts Graylog search responses into the
// logging packages LoggableMessages struct.
func searchResponseToLoggableMsgs(
	m models.SearchResponseMessages,
	l logging.Logger,
) logging.LoggableMessages {
	var lms = make(logging.LoggableMessages, len(m))

	for i := range m {
		lms[i] = &logging.LoggableMessage{
			Logger:  l,
			Message: m[i].Message.(map[string]interface{}),
		}
	}

	return lms
}
