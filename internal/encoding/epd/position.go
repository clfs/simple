package epd

import "github.com/clfs/simple/core"

// ExtendedPosition represents an EPD position.
type ExtendedPosition struct {
	Position core.Position
	Comment  string
}
