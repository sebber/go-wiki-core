package repository

import(
  "github.com/sebber/go-wiki-core/entity"
)

type WikipageRepository interface {
  GetByTitle(title string) (*entity.Page, error)
  Add(p *entity.Page) error
}