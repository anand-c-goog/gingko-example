package utils_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  . "gingko-example/utils"
)

var _ = Describe("Arrays", func() {
  var (
    containsString      []string
    doesntContainString []string
    stringToSearch      string
  )

  BeforeEach(func() {
    containsString = make([]string, 2)
    containsString[0] = "ab"
    containsString[1] = "cd"
    doesntContainString = make([]string, 1)
    doesntContainString[0] = "ab"
    stringToSearch = "cd"
  })

  Describe("Searching for existence of a string in an array", func() {
    Context("With an array containing a string", func() {
      It("string should exist", func() {
        Expect(ContainsString(containsString, stringToSearch)).To(Equal(true))
      })
    })
    Context("With an array not containing a string", func() {
      It("string should not exist", func() {
        Expect(ContainsString(doesntContainString, stringToSearch)).To(Equal(false))
      })
    })
  })
})
