package epd

import (
	"fmt"

	"github.com/clfs/simple/core"
)

// Encode encodes a position and EPD operations into an EPD string.
//
// Encode ignores p's half move clock and full move number. To provide those
// values, use [OpcodeHalfMoveClock] or [OpcodeFullMoveNumber] operations
// instead.
func Encode(p core.Position, ops []Op) string {
	return ""
}

func (u Unknown) EncodeOp() (string, error) {
	return fmt.Sprintf("%s %s;", u.Opcode, u.Input), nil
}
