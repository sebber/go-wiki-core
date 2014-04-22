package repository

import(
  "wiki/entity"
)

type MemoryWikipageRepository struct {
  Pages map[string]*entity.Page
}

func (repository MemoryWikiPageRepository) LoadPage(title string) (*entity.Page, error) {
  return repository.Pages[title], nil
}

func (repository MemoryWikiPageRepository) SavePage(p *entity.Page) error {
  repository.Pages[p.Title] = p

  return nil
}