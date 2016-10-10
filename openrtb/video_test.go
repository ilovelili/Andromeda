package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Video", func() {
	var subject *Video

	BeforeEach(func() {
		err := fixture("video", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Video{
			Mimes: []string{
				"video/x-flv",
				"video/mp4",
				"application/x-shockwave-flash",
				"application/javascript",
			},
			MinDuration:    5,
			MaxDuration:    30,
			Protocols:      []int{VideoProtoVAST2, VideoProtoVAST3},
			Width:          640,
			Height:         480,
			Linearity:      VideoLinearityLinear,
			Sequence:       1,
			BAttr:          []int{13, 14},
			MaxExtended:    30,
			MinBitrate:     300,
			MaxBitrate:     1500,
			BoxingAllowed:  iptr(1),
			PlaybackMethod: []int{VideoPlaybackAutoSoundOn, VideoPlaybackClickToPlay},
			Delivery:       []int{2},
			Pos:            AdPosAboveFold,
			CompanionAd: []Banner{
				{Width: 300, Height: 250, ID: "1234567893-1", Position: 1, BAttr: []int{13, 14}, ExpDir: []int{ExpDirRight, ExpDirDown}},
				{Width: 728, Height: 90, ID: "1234567893-2", Position: 1, BAttr: []int{13, 14}},
			},
			API:           []int{1, 2},
			CompanionType: []int{1, 2},
		}))
	})

	It("should validate", func() {
		Expect((&Video{}).Validate()).To(Equal(ErrInvalidVideoNoMimes))
	})
})
