package main

import (
	"os"
	"fmt"
	goopt "github.com/droundy/goopt"
)

var param_platform = goopt.String([]string{"-p", "--platform"}, "common", "specify a platform")

func main() {

	goopt.Description = func() string {
		return "Tldr - simplified and community-driven man pages."
	}
	goopt.Version = "0.1"
	goopt.Summary = "simplified and community-driven man pages"
	goopt.Parse(nil)

	if len(os.Args) <= 1 {
		fmt.Printf(goopt.Usage())
		os.Exit(1)
	}

	cmd := goopt.Args[0]
	platform := param_platform
	page, err := GetPageForPlatform(cmd, *platform)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(Render(page))
}
