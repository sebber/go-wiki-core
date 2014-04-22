package usecase

import (
  "github.com/sebber/go-wiki-core/repository"
  "github.com/sebber/go-wiki-core/entity"
)

type SaveWikipage struct {
  PageRepository repository.WikipageRepository
}

func NewSaveWikipage(repository repository.MemoryWikipageRepository) (SaveWikipage) {
  return SaveWikipage{PageRepository: repository}
}

func (u *SaveWikipage) Execute(title string, body []byte) error {
  p := &entity.Page{Title: title, Body: body}
  return u.PageRepository.SavePage(p)
}

type LoadWikipage struct {
  PageRepository repository.WikipageRepository
}

func NewLoadWikipage(repository repository.MemoryWikipageRepository) (LoadWikipage) {
  return LoadWikipage{PageRepository: repository}
}

func (u *LoadWikipage) Execute(title string) (*entity.Page, error) {
  page, err := u.PageRepository.LoadPage(title)

  return page, err
}