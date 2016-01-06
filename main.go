package main

import (
	"errors"
	"os"
	"os/user"
	"time"

	"github.com/pranavraja/tldr/lib/tldr"
)

var remote string = "https://raw.github.com/tldr-pages/tldr/master/pages"

func main() {
	if len(os.Args) <= 1 {
		println("Usage: tldr <command>")
		os.Exit(1)
	}

	usr, err := user.Current()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if usr.HomeDir == "" {
		println(errors.New("Can't load user's home folder path"))
		os.Exit(1)
	}

	var fetcher tldr.PageFetcher
	fetcher = tldr.NewRemotePageFetcher(remote)
	fetcher = tldr.NewFileSystemCache(fetcher, "/Users/txgruppi/.tldr", time.Hour*24)

	cmd := os.Args[1]
	platform := "common"
	page, err := fetcher.Fetch(cmd, platform)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer page.Close()
	println(tldr.Render(page.Reader()))
}
