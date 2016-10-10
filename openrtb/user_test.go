package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var subject *User

	BeforeEach(func() {
		err := fixture("user", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&User{
			ID:          "ffffffd5135596709273b3a1a07e466ea2bf4fff",
			YearOfBirth: 1984,
			Gender:      "M",
		}))
	})
})
