package tldr

import "github.com/pranavraja/tldr/lib/tldr/entity"

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
		return nil, ErrNotFound
	}

	for _, p := range platforms {
		if p == platform {
			return r.Repository.Page(name, platform)
		}
	}

	return nil, ErrNotFound
}
