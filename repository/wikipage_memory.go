package repository

import(
  "github.com/sebber/go-wiki-core/entity"
)

type MemoryWikipageRepository struct {
  Pages map[string]*entity.Page
}

func (repository MemoryWikipageRepository) All() (map[string]*entity.Page, error) {
  return repository.Pages, nil
}

func (repository MemoryWikipageRepository) GetByTitle(title string) (*entity.Page, error) {
  return repository.Pages[title], nil
}

func (repository MemoryWikipageRepository) Add(p *entity.Page) error {
  repository.Pages[p.Title] = p

  return nil
}