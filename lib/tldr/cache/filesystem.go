package cache

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/pranavraja/tldr/lib/tldr/entity"
)

var ErrTooOld = errors.New("Cached file is too old")

func NewFileSystemCacheRepository(repository entity.Repository, path string, ttl time.Duration) *FileSystemCacheRepository {
	return &FileSystemCacheRepository{
		Repository: repository,
		path:       path,
		ttl:        ttl,
	}
}

type FileSystemCacheRepository struct {
	entity.Repository
	path string
	ttl  time.Duration
}

func (f *FileSystemCacheRepository) loadFromCache(filePath string) (*os.File, error) {
	stat, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if stat.ModTime().Before(time.Now().Add(-f.ttl)) {
		return nil, ErrTooOld
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (f *FileSystemCacheRepository) loadPageFromCache(filePath string) (entity.Page, error) {
	file, err := f.loadFromCache(filePath)
	if err != nil {
		return nil, err
	}
	if file == nil {
		return nil, nil
	}
	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return NewCachedPage(contents), nil
}

func (f *FileSystemCacheRepository) loadIndexFromCache(filePath string) (entity.Index, error) {
	file, err := f.loadFromCache(filePath)
	if err != nil {
		return nil, err
	}
	if file == nil {
		return nil, nil
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	jsonMap := map[string][]string{}
	err = decoder.Decode(&jsonMap)
	if err != nil {
		return nil, err
	}
	return entity.NewIndex(jsonMap), nil
}

func (f *FileSystemCacheRepository) removeFromCache(filePath string) error {
	return os.Remove(filePath)
}

func (f *FileSystemCacheRepository) saveToCache(filePath string, data []byte) error {
	os.MkdirAll(path.Dir(filePath), 0744)
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileSystemCacheRepository) Page(page, platform string) (entity.Page, error) {
	filePath := path.Join(f.path, platform, page+".md")
	cachedPage, err := f.loadPageFromCache(filePath)
	if err == ErrTooOld {
		err = nil
		f.removeFromCache(filePath)
	}
	if err != nil {
		return nil, err
	}
	if cachedPage != nil {
		return cachedPage, nil
	}

	remotePage, err := f.Repository.Page(page, platform)
	if err != nil {
		return nil, err
	}
	defer remotePage.Close()
	contents, err := ioutil.ReadAll(remotePage.Reader())
	if err != nil {
		return nil, err
	}
	err = f.saveToCache(filePath, contents)
	if err != nil {
		return nil, err
	}

	return NewCachedPage(contents), nil
}

func (f *FileSystemCacheRepository) Index() (entity.Index, error) {
	filePath := path.Join(f.path, "index.json")
	cachedIndex, err := f.loadIndexFromCache(filePath)
	if err == ErrTooOld {
		err = nil
		f.removeFromCache(filePath)
	}
	if err != nil {
		return nil, err
	}
	if cachedIndex != nil {
		return cachedIndex, nil
	}

	remoteIndex, err := f.Repository.Index()
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(remoteIndex.Commands())
	if err != nil {
		return nil, err
	}
	err = f.saveToCache(filePath, js)
	if err != nil {
		return nil, err
	}

	return remoteIndex, nil
}
