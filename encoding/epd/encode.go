package epd

import (
	"github.com/clfs/simple/core"
)

// Op represents an EPD operation.
type Op struct {
	Opcode   string
	Operands []string
}

// EPD opcode constants.
const (
	OpcodeFullMoveNumber = "fmvn"
	OpcodeHalfMoveClock  = "hmvc"
)

// Encode encodes a position and EPD operations into an EPD string.
//
// Encode ignores p's half move clock and full move number. To provide those
// values, use [OpcodeHalfMoveClock] or [OpcodeFullMoveNumber] operations
// instead.
func Encode(p core.Position, ops []Op) string {
	return ""
}
