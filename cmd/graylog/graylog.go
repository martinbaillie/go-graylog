package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	graylogCmd = &cobra.Command{
		Use: "graylog",
		Short: "Graylog is reference CLI for Graylog's API (tested against v" +
			APIVersion + ")",
		Long: `This CLI is a (self-serving) reference implementation for a very
small subset of the Graylog API.

Golang library bindings for the entire Graylog API (v` + APIVersion + `) can
be found at https://github.com/martinbaillie/go-graylog/tree/master/pkg and
Godoc at https://godoc.org/github.com/martinbaillie/go-graylog/pkg.

This tool showcases some of these bindings. If you find it useful but buggy,
or otherwise would like a new feature, then please raise an issue at
https://github.com/martinbaillie/go-graylog/issues.

PRs are of course welcome.`,
		Args:         mandatoryArgs(nil, 0),
		SilenceUsage: true,
	}
	user       string
	pass       string
	debug      bool
	serverURLs []*url.URL
)

func init() {
	graylogCmd.PersistentFlags().StringVarP(
		&user,
		"user",
		"u",
		"admin",
		"graylog user",
	)
	graylogCmd.PersistentFlags().StringVarP(
		&pass,
		"pass",
		"p",
		"admin",
		"graylog pass (use \"-\" for masked prompt)",
	)
	graylogCmd.PersistentFlags().BoolVarP(
		&debug,
		"debug",
		"d",
		false,
		"print debug information",
	)
	graylogCmd.PersistentFlags().StringSliceP(
		"servers",
		"s",
		[]string{"localhost:9000"},
		"graylog server(s) to query",
	)

	cobra.OnInitialize(func() {
		if debug {
			log.SetLevel(log.DebugLevel)
		}
	})
}

// mandatoryArgs makes sure we have enough flags and args to perform commands
// currently in the tool. It is tailored towards the `search` command for now.
func mandatoryArgs(flags []string, minArgs int) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var missing []string
		for _, flag := range flags {
			val, err := cmd.Flags().GetString(flag)
			if err != nil || val == "" {
				missing = append(missing, flag)
			}
		}

		if len(missing) > 0 {
			return fmt.Errorf("%s subcommand missing mandatory flags: %s",
				cmd.Name(), strings.Join(missing, ", "))
		}

		if len(args) < minArgs {
			return fmt.Errorf("requires at least %d arg(s)", minArgs)
		}

		servers, _ := cmd.Flags().GetStringSlice("servers")
		serverURLs = make([]*url.URL, len(servers))
		for i := range servers {
			server, err := parseGraylogURL(servers[i])
			if err != nil {
				return err
			}
			serverURLs[i] = server
		}

		if pass == "-" {
			var err error
			if pass, err = passFromUser(); err != nil {
				return err
			}
		}

		return nil
	}
}

// parseGraylogURL parses a URL from string and enables sane Graylog defaults
// if missing (scheme, path).
func parseGraylogURL(u string) (*url.URL, error) {
	url, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	if url.Scheme != "http" && url.Scheme != "https" {
		url.Scheme = "http"
	}

	if url.Path == "" {
		url.Path = "api"
	}

	return url, nil
}

// passFromUser prompts the user to enter their Graylog password.
func passFromUser() (pass string, err error) {
	fmt.Println("Enter your Graylog password (will be hidden): ")
	oldState, err := terminal.MakeRaw(int(syscall.Stdin)) // nolint: unconvert
	if err != nil {
		return
	}
	defer func() { _ = terminal.Restore(int(syscall.Stdin), oldState) }() // nolint: unconvert

	passBytes, err := terminal.ReadPassword(int(syscall.Stdin)) // nolint: unconvert

	if err != nil {
		return
	}
	pass = string(passBytes)

	return
}

func main() {
	if err := graylogCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
