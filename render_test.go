package main

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "asd\nsdfghi\njk", "asd\nsdfghi\njk\n"},
		{"simple utf8 string", "Hello, 世界\n\ni like chips", "Hello, 世界\n\ni like chips\n"},
		{"headings", "# Hello, 世界\n\ni like chips", "\033[44m\033[1m\033[36mHello, 世界\033[0m\n\ni like chips\n"},
		{"quotation", "Title\n> Raja", "Title\n  \033[37mRaja\033[0m\n"},
		{"inline list", "Title\n- Raja\n- Pranav", "Title\n\033[32m- Raja\033[0m\n\033[32m- Pranav\033[0m\n"},
		{"code", "Title\n`go build`\n`go test`", "Title\n  \033[40m\033[37mgo build\033[0m\n  \033[40m\033[37mgo test\033[0m\n"},
	}
	for _, test := range tests {
		rendered := Render(strings.NewReader(test.input))
		if rendered != test.expected {
			t.Errorf("Incorrect render of %s: got '%s', want '%s'", test.name, rendered, test.expected)
		}
	}
}
