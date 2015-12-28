package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-colorable"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Usage: tldr <command>")
		os.Exit(1)
	}
	cmd := os.Args[1]
	platform := "common"
	page, err := GetPageForPlatform(cmd, platform)
	if err != nil {
		fmt.Fprintln(os.Stderr, os.Args[0]+":", err.Error())
		os.Exit(1)
	}
	fmt.Fprintln(colorable.NewColorableStdout(), Render(page))
}
