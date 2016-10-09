package openrtb

// 5.2 Banner Ad Types
const (
	BannerTypeXHTMLText int = iota + 1
	BannerTypeXHTML
	BannerTypeJS
	BannerTypeFrame
)

// 5.3 Creative Attributes
const (
	CreativeAttrAudioAdAutoplay int = iota + 1
	CreativeAttrAudioAdUserInitiated
	CreativeAttrExpandableAutomatic
	CreativeAttrExpandableUserInitiatedClick
	CreativeAttrExpandableUserInitiatedRollover
	CreativeAttrInBannerVideoAdAutoPlay
	CreativeAttrInBannerVideoAdUserInitiated
	CreativeAttrPop
	CreativeAttrProvocativeOrSuggestiveImagery
	CreativeAttrShakyFlashingFlickeringExtremeAnimationSmileys
	CreativeAttrSurveys
	CreativeAttrTextOnly
	CreativeAttrUserInteractive
	CreativeAttrWindowsDialogOrAlertStyle
	CreativeAttrHasAudioOnOffButton
	CreativeAttrAdProvidesSkipButton
	CreativeAttrAdobeFlash
)

// 5.4 Ad Position
const (
	AdPosUnknown int = iota
	AdPosAboveFold
	AdPosDEPRECATED
	AdPosBelowFold
	AdPosHeader
	AdPosFooter
	AdPosSidebar
	AdPosFullscreen
)

// 5.5 Expandable Direction
const (
	ExpDirLeft int = iota + 1
	ExpDirRight
	ExpDirUp
	ExpDirDown
	ExpDirFullScreen
)

// 5.6 API Frameworks
const (
	APIFrameworkVPAID1 int = iota + 1
	APIFrameworkVPAID2
	APIFrameworkMRAID1
	APIFrameworkORMMA
	APIFrameworkMRAID2
)

// 5.7 Video Linearity
const (
	VideoLinearityLinear int = iota + 1
	VideoLinearityNonLinear
)

// 5.8 Video Bid Response Protocols
const (
	VideoProtoVAST1 int = iota + 1
	VideoProtoVAST2
	VideoProtoVAST3
	VideoProtoVAST1Wrapper
	VideoProtoVAST2Wrapper
	VideoProtoVAST3Wrapper
	VideoProtoVAST4
	VideoProtoDAAST1
)

// 5.9 Video Playback Methods
const (
	VideoPlaybackAutoSoundOn int = iota + 1
	VideoPlaybackAutoSoundOff
	VideoPlaybackClickToPlay
	VideoPlaybackMouseOver
)

// 5.10 Video Start Delay
const (
	VideoStartDeleyMidRoll         = 1 // actually > 0
	VideoStartDelayPreRoll         = 0
	VideoStartDelayGenericMidRoll  = -1
	VideoStartDelayGenericPostRoll = -2
)

// 5.11 Video Quality
const (
	VideoQualityUnknown int = iota
	VideoQualityProfessional
	VideoQualityProsumer
	VideoQualityUGC
)

// 5.12 VAST Companion Types
const (
	CompanionTypeStatic int = iota + 1
	CompanionTypeHTML
	CompanionTypeIFrame
)

// 5.13 Content Delivery Methods
const (
	ContentDeliveryStreaming int = iota + 1
	ContentDeliveryProgressive
)

// 5.14 Feed Types
const (
	FeedTypeSingleFeed int = iota + 1
	FeedTypeMultipleFeed
	FeedTypeFMAMBroadcast
	FeedTypePodcast
)

// 5.15 Volume Normalization Modes
const (
	VolumeNormModeNone = iota
	VolumeNormAverage
	VolumnNormPeak
	VolumnNormCustom
)

// 5.16 Content Context
const (
	ContextVideo int = iota + 1
	ContextGame
	ContextMusic
	ContextApplication
	ContextText
	ContextOther
	ContextUnknown
)

// 5.17 IQG Media Ratings
const (
	QAGAll int = iota + 1
	QAGOver12
	QAGMature
)

// 5.18 Location Type
const (
	LocationTypeGPS int = iota + 1
	LocationTypeIP
	LocationTypeUser
)

// 5.19 Device Type
const (
	DeviceTypeUnknown int = iota
	DeviceTypeMobile
	DeviceTypePC
	DeviceTypeTV
	DeviceTypePhone
	DeviceTypeTablet
	DeviceTypeConnectedDevice
	DeviceTypeSetTopBox
)

// 5.20 Connection Type
const (
	ConnTypeUnknown int = iota
	ConnTypeEthernet
	ConnTypeWIFI
	ConnTypeCellUnknownGeneration
	ConnTypeCell2G
	ConnTypeCell3G
	ConnTypeCell4G
)

// 5.21 IP Location Service
const (
	IPLocationServiceIP2Location = iota + 1
	IPLocationServiceNeustar
	IPLocationServiceMaxMind
)

// 5.22 No-Bid Reason Codes
const (
	NBRUnknownError int = iota
	NBRTechnicalError
	NBRInvalidRequest
	NBRKnownSpider
	NBRSuspectedNonHuman
	NBRProxyIP
	NBRUnsupportedDevice
	NBRBlockedSite
	NBRUnmatchedUser
)
