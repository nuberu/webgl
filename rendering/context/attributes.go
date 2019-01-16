package context

import "github.com/nuberu/webgl"

type Attributes struct {
	Alpha                        bool
	Antialias                    bool
	Depth                        bool
	PremultipliedAlpha           bool
	PreserveDrawingBuffer        bool
	Stencil                      bool
	PowerPreference              webgl.PowerPreference
	FailIfMajorPerformanceCaveat bool
}
