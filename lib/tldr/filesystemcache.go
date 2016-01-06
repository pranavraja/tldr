package tldr

import (
	"io/ioutil"
	"os"
	"path"
	"time"
)

func NewFileSystemCache(decorated PageFetcher, path string, ttl time.Duration) PageFetcher {
	return &FileSystemCache{
		decorated: decorated,
		path:      path,
		ttl:       ttl,
	}
}

type FileSystemCache struct {
	decorated PageFetcher
	path      string
	ttl       time.Duration
}

func (f *FileSystemCache) Fetch(page, platform string) (Page, error) {
	filePath := path.Join(f.path, platform, page+".md")
	cachedPage, err := f.loadFromCache(filePath)
	if err != nil {
		return nil, err
	}
	if cachedPage != nil {
		return cachedPage, nil
	}

	remotePage, err := f.decorated.Fetch(page, platform)
	if err != nil {
		return nil, err
	}
	defer remotePage.Close()
	contents, err := ioutil.ReadAll(remotePage.Reader())
	if err != nil {
		return nil, err
	}
	os.MkdirAll(path.Dir(filePath), 0744)
	err = ioutil.WriteFile(filePath, contents, 0644)
	if err != nil {
		return nil, err
	}

	return NewCachedPage(contents), nil
}

func (f *FileSystemCache) loadFromCache(filePath string) (Page, error) {
	stat, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if stat.ModTime().Before(time.Now().Add(-f.ttl)) {
		return nil, nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return NewCachedPage(contents), nil
}
