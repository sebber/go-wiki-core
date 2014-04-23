package usecase_test

import (
  "github.com/sebber/go-wiki-core/entity"
  "github.com/sebber/go-wiki-core/repository"
  "github.com/sebber/go-wiki-core/usecase"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Wikipage", func() {

  It("Saving page with title 'A page' should get saved", func() {
    repo := repository.MemoryWikipageRepository{}
    repo.Pages = make(map[string]*entity.Page)

    SaveWiki := usecase.SaveWikipage{repo}
    SaveWiki.Execute("A page", []byte("body"))

    LoadWiki := usecase.LoadWikipage{repo}
    page, err := LoadWiki.Execute("A page")

    Expect(err).To(BeNil())
    Expect(page.Title).To(Equal("A page"))
  })
})
