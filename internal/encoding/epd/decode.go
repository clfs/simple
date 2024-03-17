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
func Decode(s string) (core.Position, []Op3, error) {
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

	ops, err := parseOps(fields[4])
	if err != nil {
		return core.Position{}, nil, err
	}

	for _, op := range ops {
		if err := applyOp(&p, op); err != nil {
			return core.Position{}, nil, err
		}
	}

	return p, ops, nil
}

func parseOps(s string) ([]Op3, error) {
	var (
		ops     []Op3
		inQuote bool
		head    int
	)

	for i, rn := range s {
		switch {
		case rn == '"':
			inQuote = !inQuote
		case rn == ';' && !inQuote:
			op, err := parseOp(s[head:i])
			if err != nil {
				return nil, err
			}
			ops = append(ops, op)
			head = i + 1
		}
	}

	return ops, nil
}

func parseOp(s string) (Op3, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return Op3{}, fmt.Errorf("operation has no opcode")
	}
	opcode, operands, _ := strings.Cut(s, " ")
	return Op3{opcode, operands}, nil
}

func applyOp(p *core.Position, op Op3) error {
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
