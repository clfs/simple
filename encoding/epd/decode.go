package epd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
)

// Decode decodes an EPD string into a position and operations.
//
// The returned position has a half move clock of 0 and a full move number of 1,
// unless the EPD string contains operations that specify them.
func Decode(s string) (core.Position, []Op, error) {
	fields := strings.SplitN(s, " ", 5)

	if n := len(fields); n < 4 {
		return core.Position{}, nil, fmt.Errorf("too few fields: %d", n)
	}

	pseudoFEN := fmt.Sprintf("%s 0 1", strings.Join(fields[:4], " "))

	p, err := fen.Decode(pseudoFEN)
	if err != nil {
		return core.Position{}, nil, err
	}

	// Return early if there are no operations.
	if len(fields) == 4 {
		return p, nil, nil
	}

	ops, err := decodeOps(fields[4])
	if err != nil {
		return core.Position{}, nil, err
	}

	for _, op := range ops {
		applyOp(&p, op)
	}

	return p, ops, nil
}

func decodeOps(_ string) ([]Op, error) {
	// TODO: Implement.
	return nil, nil
}

func applyOp(p *core.Position, op Op) error {
	switch op.Opcode {
	case OpcodeFullMoveNumber:
		n, err := strconv.Atoi(op.Args)
		if err != nil {
			return err
		}
		p.FullMoveNumber = n
	case OpcodeHalfMoveClock:
		n, err := strconv.Atoi(op.Args)
		if err != nil {
			return err
		}
		p.HalfMoveClock = n
	}
	return nil
}
