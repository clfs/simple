package epd

import (
	"github.com/clfs/simple/core"
)

// Encode encodes a position and EPD operations into an EPD string.
//
// Encode ignores p's half move clock and full move number. To include them in
// the EPD string, use [OpcodeFullMoveNumber] or [OpcodeHalfMoveClock]
// operations.
func Encode(p core.Position, ops []Op) string {
	return ""
}
