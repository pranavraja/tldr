package main

import (
	"errors"
	"os"
	"os/user"
	"time"

	"github.com/pranavraja/tldr/lib/tldr"
	"github.com/pranavraja/tldr/lib/tldr/cache"
	"github.com/pranavraja/tldr/lib/tldr/entity"
	"github.com/pranavraja/tldr/lib/tldr/remote"
)

var remoteAddress string = "https://raw.github.com/tldr-pages/tldr/master/pages"

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

	var repository entity.Repository
	repository = remote.NewRemoteRepository(remoteAddress)
	repository = cache.NewFileSystemCacheRepository(repository, "/Users/txgruppi/.tldr", time.Hour*24)
	repository = tldr.NewIndexCheckerRepository(repository)

	cmd := os.Args[1]

	page, err := repository.Page(cmd, "common")
	if err != nil {
		return err
	}
	defer page.Close()
	println(tldr.Render(page.Reader()))
	return nil
}
