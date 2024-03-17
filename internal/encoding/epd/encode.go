package epd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
)

func encodeOp(op Op3) (string, error) {
	if op.Opcode == "" {
		return "", fmt.Errorf("missing opcode")
	}
	if strings.Contains(op.Opcode, ";") {
		return "", fmt.Errorf("opcode contains semicolon")
	}
	if strings.Contains(op.Operands, ";") {
		return "", fmt.Errorf("operands contain semicolon")
	}
	if op.Operands == "" {
		return fmt.Sprintf("%s;", op.Opcode), nil
	}
	return fmt.Sprintf("%s %s;", op.Opcode, op.Operands), nil
}

// Encode encodes a position and EPD operations into an EPD string.
//
// Encode ignores p's half move clock and full move number. To specify them in
// the EPD string, use [OpcodeHalfMoveClock] or [OpcodeFullMoveNumber]
// operations.
func Encode(p core.Position, ops []Op3) (string, error) {
	var buf bytes.Buffer

	fenFields := strings.Fields(fen.Encode(p))

	// Ignore the half move clock and full move number.
	fmt.Fprint(&buf, strings.Join(fenFields[:4], " "))

	for _, op := range ops {
		s, err := encodeOp(op)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&buf, " %s", s)
	}

	return buf.String(), nil
}
