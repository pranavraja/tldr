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
	err := run()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) <= 1 {
		return errors.New("Usage: tldr <command>")
	}

	usr, err := user.Current()
	if err != nil {
		return err
	}
	if usr.HomeDir == "" {
		return errors.New("Can't load user's home folder path")
	}

	var fetcher tldr.PageFetcher
	fetcher = tldr.NewRemotePageFetcher(remote)
	fetcher = tldr.NewFileSystemCache(fetcher, "/Users/txgruppi/.tldr", time.Hour*24)

	cmd := os.Args[1]
	platform := "common"
	page, err := fetcher.Fetch(cmd, platform)
	if err != nil {
		return err
	}
	defer page.Close()
	println(tldr.Render(page.Reader()))
	return nil
}
