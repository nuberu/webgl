package types

type Attributes struct {
	Alpha                        bool
	Antialias                    bool
	Depth                        bool
	FailIfMajorPerformanceCaveat bool
	PowerPreference              PowerPreference
	PremultipliedAlpha           bool
	PreserveDrawingBuffer        bool
	Stencil                      bool
	Storage                      string // Only Chromium
	WillReadFrequently           bool   // Only Firefox
}
