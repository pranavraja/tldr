package main

import (
	"fmt"
	"net/http"
	"bytes"
	"bufio"
	"os"
//	"os/user"
)

var remote string = "http://raw.github.com/rprieto/tldr/master/pages"

// Caller must close the response body after reading.
func GetPageForPlatform(page, platform string) (string, error) {
	//usr, _ := user.Current()
	pagePath := CfgParams.CacheDirectory + platform + "/" + page + ".md"
	result := ""

	if _, err := os.Stat(CfgParams.CacheDirectory + platform); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(CfgParams.CacheDirectory + platform + "/", 0755)
		}
	}


	if _, err := os.Stat(pagePath); err != nil {
		if os.IsNotExist(err) {
			resp, err := http.Get(remote + "/" + platform + "/" + page + ".md")
			if err != nil {
				return "", err
			}
			if resp.StatusCode == 404 {
				return "", fmt.Errorf("Not found.\nTo add this command, send Romain a pull request at:\n  https://github.com/rprieto/tldr")
			}
			if resp.StatusCode != 200 {
				return "", fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			result = buf.String()

			cached, err := os.Create(pagePath)
			if err != nil {
				return "", err
			}
			cached.WriteString(result)
			cached.Close()
		}
	} else {
		cached, err := os.Open(pagePath)
		if err != nil {
			return "", err
		}
		defer cached.Close()
		scanner := bufio.NewScanner(cached)
		for scanner.Scan() {
			result = result + scanner.Text() + "\n"
		}
	}
	return result, nil
}
