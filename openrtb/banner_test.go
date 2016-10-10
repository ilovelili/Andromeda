package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Banner", func() {
	var subject *Banner

	BeforeEach(func() {
		err := fixture("banner", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Banner{
			Width:    728,
			Height:   90,
			Position: AdPosAboveFold,
			BType:    []int{4},
			BAttr:    []int{14},
			API:      []int{3},
		}))
	})

})
