package tldr

import (
	"errors"

	"github.com/pranavraja/tldr/lib/tldr/entity"
)

func NewIndexCheckerRepository(repository entity.Repository) *IndexCheckerRepository {
	return &IndexCheckerRepository{
		Repository: repository,
	}
}

type IndexCheckerRepository struct {
	entity.Repository
}

func (r *IndexCheckerRepository) Page(name, platform string) (entity.Page, error) {
	index, err := r.Repository.Index()
	if err != nil {
		return nil, err
	}
	platforms := index.PlatformsFor(name)
	if platforms == nil || len(platforms) == 0 {
		return nil, errors.New("Not found.\nTo add this command, send a pull request at:\n  https://github.com/tldr-pages/tldr")
	}

	// Ignore requested platform and use the first platform from index
	return r.Repository.Page(name, platforms[0])
}
