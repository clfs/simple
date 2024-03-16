package epd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
)

// Encode encodes a position and EPD operations into an EPD string.
//
// Encode ignores p's half move clock and full move number. To specify them in
// the EPD string, use [OpcodeHalfMoveClock] or [OpcodeFullMoveNumber]
// operations.
func Encode(p core.Position, ops []Op) string {
	var buf bytes.Buffer

	fenFields := strings.Fields(fen.Encode(p))

	// Ignore the half move clock and full move number.
	fmt.Fprint(&buf, strings.Join(fenFields[:4], " "))

	for _, op := range ops {
		fmt.Fprintf(&buf, " %s %s;", op.Opcode, op.Operand)
	}

	return buf.String()
}
