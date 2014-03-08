package main

import (
	"bufio"
	"io"
	"strings"
)

const (
	TermEscapePrefix = "\033["
	TermEscapeSuffix = "m"

	Reset = TermEscapePrefix + "0" + TermEscapeSuffix

	Bold = TermEscapePrefix + "1" + TermEscapeSuffix

	Green = TermEscapePrefix + "32" + TermEscapeSuffix
	Cyan  = TermEscapePrefix + "36" + TermEscapeSuffix
	White = TermEscapePrefix + "37" + TermEscapeSuffix

	OnBlue = TermEscapePrefix + "44" + TermEscapeSuffix
	OnGrey = TermEscapePrefix + "40" + TermEscapeSuffix
)

// Render pretties up some markdown for terminal display.
// Only a small subset of markdown is actually supported,
// as described in https://github.com/rprieto/tldr/blob/master/CONTRIBUTING.md#markdown-format
func Render(markdown io.Reader) (rendered string) {
	scanner := bufio.NewScanner(markdown)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			// Heading
			rendered += OnBlue + Bold + Cyan + strings.TrimLeft(line, "# ") + Reset + "\n"
		} else if strings.HasPrefix(line, ">") {
			// Quotation
			rendered += "  " + White + strings.TrimLeft(line, "> ") + Reset + "\n"
		} else if strings.HasPrefix(line, "-") {
			// Inline list
			rendered += Green + line + Reset + "\n"
		} else if strings.HasPrefix(line, "`") {
			// Code
			rendered += "  " + OnGrey + White + strings.Trim(line, "`") + Reset + "\n"
		} else {
			rendered += line + "\n"
		}
	}
	return rendered
}
