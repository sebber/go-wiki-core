package repository

import(
  "github.com/sebber/go-wiki-core/entity"
)

type MemoryWikipageRepository struct {
  Pages map[string]*entity.Page
}

func (repository MemoryWikipageRepository) LoadPage(title string) (*entity.Page, error) {
  return repository.Pages[title], nil
}

func (repository MemoryWikipageRepository) SavePage(p *entity.Page) error {
  repository.Pages[p.Title] = p

  return nil
}