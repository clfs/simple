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
// The returned position defaults to a half move clock of 0 and a full move
// number of 1, unless the EPD string contains [OpcodeHalfMoveClock] or
// [OpcodeFullMoveNumber] operations.
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

	ops := parseOps(fields[4])

	for _, op := range ops {
		if err := applyOp(&p, op); err != nil {
			return core.Position{}, nil, err
		}
	}

	return p, ops, nil
}

func parseOps(s string) []Op {
	var (
		ops     []Op
		rawOp   []rune
		inQuote bool
	)

	for _, rn := range s {
		switch {
		case rn == ';':
			if !inQuote {
				ops = append(ops, parseOp(string(rawOp)))
				rawOp = nil
				continue
			}
		}
		if rn == '"' {
			inQuote = !inQuote
		}
		rawOp = append(rawOp, rn)
	}

	return ops
}

func parseOp(s string) Op {
	s = strings.TrimSpace(s)
	opcode, operands, _ := strings.Cut(s, " ")
	return Op{opcode, operands}
}

func applyOp(p *core.Position, op Op) error {
	switch op.Opcode {
	case OpcodeFullMoveNumber:
		n, err := strconv.Atoi(op.Operands)
		if err != nil {
			return fmt.Errorf("invalid full move number: %s", op.Operands)
		}
		p.FullMoveNumber = n
	case OpcodeHalfMoveClock:
		n, err := strconv.Atoi(op.Operands)
		if err != nil {
			return fmt.Errorf("invalid half move clock: %s", op.Operands)
		}
		p.HalfMoveClock = n
	}
	return nil
}
