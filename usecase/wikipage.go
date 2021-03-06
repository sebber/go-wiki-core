package usecase

import (
  "github.com/sebber/go-wiki-core/repository"
  "github.com/sebber/go-wiki-core/entity"
)

/* ### Usecase for saving pages ### */

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

/* ### Usecase for finding specific page ### */

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

/* ### Usecase for loading all pages ### */

type GetAllWikipages struct {
  PageRepository repository.WikipageRepository
}

func NewGetAllWikipages(repository repository.WikipageRepository) (GetAllWikipages) {
  return GetAllWikipages{PageRepository: repository}
}

func (r *GetAllWikipages) Execute() (map[string]*entity.Page, error) {
  pages, err := r.PageRepository.All()

  return pages, err
}