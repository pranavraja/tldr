package main

import (
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
func Render(markdown string) (rendered string) {
	lines := strings.Split(markdown, "\n")
	for i := range lines {
		if strings.HasPrefix(lines[i], "#") {
			// Heading
			rendered += OnBlue + Bold + Cyan + strings.TrimLeft(lines[i], "# ") + Reset + "\n"
		} else if strings.HasPrefix(lines[i], ">") {
			// Quotation
			rendered += "  " + White + strings.TrimLeft(lines[i], "> ") + Reset + "\n"
		} else if strings.HasPrefix(lines[i], "-") {
			// Inline list
			rendered += Green + lines[i] + Reset + "\n"
		} else if strings.HasPrefix(lines[i], "`") {
			// Code
			rendered += "  " + OnGrey + White + strings.Trim(lines[i], "`") + Reset + "\n"
		} else {
			rendered += lines[i] + "\n"
		}
	}
	return rendered
}
