package usecase_test

import (
  "github.com/sebber/go-wiki-core/entity"
  "github.com/sebber/go-wiki-core/repository"
  "github.com/sebber/go-wiki-core/usecase"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Wikipage", func() {

  var (
    repo repository.MemoryWikipageRepository

    saveWikipage usecase.SaveWikipage
    loadWikipage usecase.LoadWikipage
    getAllWikipages usecase.GetAllWikipages
  )

  BeforeEach(func() {
    repo = repository.MemoryWikipageRepository{}
    repo.Pages = make(map[string]*entity.Page)

    saveWikipage = usecase.SaveWikipage{repo}
    loadWikipage = usecase.LoadWikipage{repo}
    getAllWikipages = usecase.GetAllWikipages{repo}
  })

  Context("Usecase API", func() {
    It("should return 0 pages if trying to find all when none exist", func() {
      pages, err := getAllWikipages.Execute()
      Expect(err).To(BeNil())
      Expect(len(pages)).To(Equal(0))
    })

    It("Saving page with title 'A page' should be available to find", func() {
      saveWikipage.Execute("A page", []byte("body"))
      page, err := loadWikipage.Execute("A page")
      Expect(err).To(BeNil())
      Expect(page.Title).To(Equal("A page"))
    })

    It("should return 1 page of trying to find all when 1 has been created", func() {
      saveWikipage.Execute("A page", []byte("body"))
      pages, err := getAllWikipages.Execute()
      Expect(err).To(BeNil())
      Expect(len(pages)).To(Equal(1))
    })
  })
  
})
