[![License](https://img.shields.io/badge/license-BSD-brightgreen.svg?style=flat-square)](/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/martinbaillie/go-graylog?style=flat-square)](https://goreportcard.com/report/github.com/martinbaillie/go-graylog)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/martinbaillie/go-graylog)
[![Build](https://travis-ci.org/martinbaillie/go-graylog.svg?branch=master&style=flat-square)](https://travis-ci.org/martinbaillie/go-graylog)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/martinbaillie/go-graylog/releases/latest)

# go-graylog

This repository contains Golang bindings for the complete Graylog API and a reference CLI that implements a very small subset of the API.

The CLI is self-serving at this stage i.e. I needed the current subset of commands implemented. However, if you find it useful and but lacking a certain command you need then don't be shy about raising an issue. Likewise, PRs are most welcome.

### API
The library is currently coded to Graylog `v2.4.0-beta.3+a6b18a2`. YMMV if you're not on the 2.4 stream of Graylog (latest at the time of writing this).

It starts here: `go get -u github.com/martinbaillie/go-graylog/pkg` [[Godoc]](https://godoc.org/github.com/martinbaillie/go-graylog/pkg).

If you're developing with the library I would start by referencing your Graylog server's API browser at `<graylog server>/api/api-browser` but then searching through `pkg` or the Godoc for the name of the endpoint—it's often not in the package you would expect due to some [awkward codegen](./api/README.md).

### CLI

At this time, only search functionality is implemented in the CLI.

Features of note:

- Ability to search and merge messages from *multiple* Graylog servers with (`-s`) with colour coded results to differentiate (if `$TERM` supports).
- Follow mode (`-f`) in `search relative` command for tail-like behaviour.
- Use `-p, --pass -` to keep passwords out of your shell history.
- Show/hide Graylog's own timestamp (in UTC or Local) with `-t, --timestamps`.

##### Search Options
```
Usage:
  graylog search [command]

Available Commands:
  absolute    Search for messages using an absolute timerange
  keyword     Search for messages in a natural language timerange
  relative    Search for messages starting from a relative timestamp

Flags:
  -h, --help         help for search
  -t, --timestamps   print message timestamps (local timezone)
      --utc -t       print UTC message timestamps (has no effect without -t)

Global Flags:
  -p, --pass string       graylog pass (use "-" for masked prompt) (default "admin")
  -s, --servers strings   graylog server(s) to query (default [localhost:9000])
  -u, --user string       graylog user (default "admin")

Use "graylog search [command] --help" for more information about a command.
```

##### Search: Absolute
```
Usage:
  graylog search [command]

Available Commands:
  absolute    Search for messages using an absolute timerange
  keyword     Search for messages using a natural language timerange
  relative    Search for messages starting from a relative timestamp

Flags:
  -h, --help         help for search
  -t, --timestamps   print message timestamps (local timezone)
      --utc -t       print UTC message timestamps (has no effect without -t)

Global Flags:
  -p, --pass string       graylog pass (use "-" for masked prompt) (default "admin")
  -s, --servers strings   graylog server(s) to query (default [localhost:9000])
  -u, --user string       graylog user (default "admin")

Use "graylog search [command] --help" for more information about a command.
```

##### Search: Keyword
```
Search for messages in a natural language timerange e.g. "yesterday" or "2 weeks ago to wednesday".

Usage:
  graylog search keyword [flags]

Flags:
  -h, --help             help for keyword
      --keyword string   natural language timerange e.g. "yesterday" or "2 weeks ago to wednesday"

Global Flags:
  -p, --pass string       graylog pass (use "-" for masked prompt) (default "admin")
  -s, --servers strings   graylog server(s) to query (default [localhost:9000])
  -t, --timestamps        print message timestamps (local timezone)
  -u, --user string       graylog user (default "admin")
      --utc -t            print UTC message timestamps (has no effect without -t)
```

##### Search: Relative
```
Search for messages starting from a relative timestamp until now.

The timestamp must be a valid Golang duration e.g. "2h".

Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

Usage:
  graylog search relative [flags]

Flags:
  -f, --follow              follow mode (tail-like functionality)
  -h, --help                help for relative
  -i, --interval duration   follow mode polling interval e.g. "5m" (default 5s)
      --since duration      from this relative Golang timestamp until now e.g. "2h"

Global Flags:
  -p, --pass string       graylog pass (use "-" for masked prompt) (default "admin")
  -s, --servers strings   graylog server(s) to query (default [localhost:9000])
  -t, --timestamps        print message timestamps (local timezone)
  -u, --user string       graylog user (default "admin")
      --utc -t            print UTC message timestamps (has no effect without -t)
```

