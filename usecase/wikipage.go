package usecase

import (
  "github.com/sebber/go-wiki-core/repository"
  "github.com/sebber/go-wiki-core/entity"
)

type SaveWikipage struct {
  PageRepository repository.WikipageRepository
}

func NewSaveWikipage(repository repository.WikipageRepository) (SaveWikipage) {
  return SaveWikipage{PageRepository: repository}
}

func (u *SaveWikipage) Execute(title string, body []byte) error {
  p := &entity.Page{Title: title, Body: body}
  return u.PageRepository.Add(p)
}

type LoadWikipage struct {
  PageRepository repository.WikipageRepository
}

func NewLoadWikipage(repository repository.WikipageRepository) (LoadWikipage) {
  return LoadWikipage{PageRepository: repository}
}

func (u *LoadWikipage) Execute(title string) (*entity.Page, error) {
  page, err := u.PageRepository.GetByTitle(title)

  return page, err
}