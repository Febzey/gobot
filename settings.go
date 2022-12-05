package main

type Settings struct {
	Locale             string
	ViewDistance       int
	ChatMode           int
	ChatColors         bool
	DisplayedSkinParts uint8
	MainHand           int

	EnableTextFiltering bool
	AllowListing        bool

	Brand string
}

const (
	_ = 1 << iota
	Jacket
	LeftSleeve
	RightSleeve
	LeftPantsLeg
	RightPantsLeg
	Hat
)

var DefaultSettings = Settings{
	Locale:             "en_US",
	ViewDistance:       4,
	ChatMode:           0,
	DisplayedSkinParts: Jacket | LeftSleeve | RightSleeve | LeftPantsLeg | RightPantsLeg | Hat,
	MainHand:           1,

	EnableTextFiltering: false,
	AllowListing:        true,

	Brand: "ForestBot-1.0",
}
