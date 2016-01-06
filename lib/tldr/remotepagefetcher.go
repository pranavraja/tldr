package tldr

import (
	"fmt"
	"net/http"
)

func NewRemotePageFetcher(remote string) PageFetcher {
	return &RemotePageFetcher{
		remote: remote,
	}
}

type RemotePageFetcher struct {
	remote string
}

// Caller must close the response body after reading.
func (f *RemotePageFetcher) Fetch(page, platform string) (Page, error) {
	resp, err := http.Get(f.remote + "/" + platform + "/" + page + ".md")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("Not found.\nTo add this command, send Romain a pull request at:\n  https://github.com/rprieto/tldr")
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	return NewRemotePage(resp.Body), nil
}
