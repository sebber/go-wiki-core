package repository

import(
  "github.com/sebber/go-wiki-core/entity"
)

type WikipageRepository interface {
  SavePage(p *entity.Page) error
  LoadPage(title string) (*entity.Page, error)
}