package epd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
)

func Decode(s string) (core.Position, []Op, error) {
	fields := strings.SplitN(s, " ", 5)

	if n := len(fields); n < 4 {
		return core.Position{}, nil, fmt.Errorf("too few fields: %d", n)
	}

	pseudoFEN := strings.Join(fields[:4], " ") + " 0 1"

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
		if err := applyOp(&p, op); err != nil {
			return core.Position{}, nil, err
		}
	}

	return p, ops, nil
}

func decodeOps(s string) ([]Op, error) {
	// TODO: Implement.
	return nil, nil
}

func applyOp(p *core.Position, op Op) error {
	switch op.Opcode {
	case OpcodeFullMoveNumber:
		n, err := strconv.Atoi(op.Operands[0])
		if err != nil {
			return fmt.Errorf("invalid full move number: %v", err)
		}
		p.FullMoveNumber = n
	case OpcodeHalfMoveClock:
		n, err := strconv.Atoi(op.Operands[0])
		if err != nil {
			return fmt.Errorf("invalid half move clock: %v", err)
		}
		p.HalfMoveClock = n
	}

	return nil
}
