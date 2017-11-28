package remote

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pranavraja/tldr/lib/tldr"
	"github.com/pranavraja/tldr/lib/tldr/entity"
)

func NewRemoteRepository(remote string) *Repository {
	return &Repository{
		remote: remote,
	}
}

type Repository struct {
	remote string
}

// Caller must close the response body after reading.
func (f *Repository) Page(page, platform string) (entity.Page, error) {
	resp, err := http.Get(f.remote + "/" + platform + "/" + page + ".md")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		resp.Body.Close()
		return nil, tldr.ErrNotFound
	}
	if resp.StatusCode != 200 {
		resp.Body.Close()
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	return NewRemotePage(resp.Body), nil
}

func (f *Repository) Index() (entity.Index, error) {
	resp, err := http.Get("https://tldr-pages.github.io/assets/index.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	data := struct {
		Commands []struct {
			Name      string   `json:"name"`
			Platforms []string `json:"platform"`
		} `json:"commands"`
	}{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	indexMap := map[string][]string{}
	for _, command := range data.Commands {
		indexMap[command.Name] = command.Platforms
	}
	return entity.NewIndex(indexMap), nil
}
