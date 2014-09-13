[![Build Status](https://drone.io/github.com/Like-all/tldr-client/status.png)](https://drone.io/github.com/Like-all/tldr-client/latest)

A work-in-progress [Go](http://golang.org/) client for [tldr](https://github.com/rprieto/tldr/).

# Setup

If you have Go installed, grab the latest version:

    go get github.com/pranavraja/tldr

Binary releases for common platforms are available under [Releases](https://github.com/pranavraja/tldr/releases/latest). These work without Go installed.

# Prerequisites

If you installed using `go get`, make sure `$GOPATH/bin` is in your `$PATH`.

# Usage

    tldr <command>

Fetch the docs for `command` and render them to the terminal.

# Hacking

Once you have cloned the repo, build using `go build`, run the tests using `go test`.

# TODO (contributions welcome)

- Improve rendering of command placeholders, like in `sed 's/a/b/' {{filename}}`
- ~~Caching of commands~~
- ~~Add a command-line flag to override the platform (currently only "common" is supported)~~
- Improve multi-line command rendering
