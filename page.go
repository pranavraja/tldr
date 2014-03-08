package main

import (
	"fmt"
	"io"
	"net/http"
)

var remote string = "http://raw.github.com/rprieto/tldr/master/pages"

// Caller must close the response body after reading.
func GetPageForPlatform(page, platform string) (io.ReadCloser, error) {
	resp, err := http.Get(remote + "/" + platform + "/" + page + ".md")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("Not found.\nTo add this command, send Romain a pull request at:\n  https://github.com/rprieto/tldr")
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	return resp.Body, nil
}
