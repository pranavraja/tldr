package main

import (
	"os"

	"github.com/pranavraja/tldr/lib/tldr"
)

func main() {
	if len(os.Args) <= 1 {
		println("Usage: tldr <command>")
		os.Exit(1)
	}
	cmd := os.Args[1]
	platform := "common"
	page, err := tldr.GetPageForPlatform(cmd, platform)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(tldr.Render(page))
}
