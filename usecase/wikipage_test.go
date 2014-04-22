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
    Repo := repository.MemoryWikipageRepository{}
    Repo.Pages = make(map[string]*entity.Page)

    SaveWikipageUseCase := usecase.SaveWikipage{Repo}
    SaveWikipageUseCase.Execute("A page", []byte("body"))

    page, err := Repo.LoadPage("A page")


    Expect(err).To(BeNil())
    Expect(page.Title).To(Equal("A page"))
  })
})
