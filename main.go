package main

import (
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		println("Usage: tldr <command>")
		os.Exit(1)
	}
	cmd := os.Args[1]
	platform := "common"
	page, err := GetPageForPlatform(cmd, platform)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(Render(page))
}
