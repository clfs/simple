package epd

import (
	"fmt"
	"strings"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
)

func Decode(s string) (core.Position, []Operation, error) {
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

	ops, err := decodeOperations(fields[4])
	if err != nil {
		return core.Position{}, nil, err
	}

	for _, op := range ops {
		applyOperation(&p, op)
	}

	return p, ops, nil
}

func decodeOperations(_ string) ([]Operation, error) {
	// TODO: Implement.
	return nil, nil
}

func applyOperation(p *core.Position, op Operation) {
	switch v := op.(type) {
	case FMVN:
		p.FullMoveNumber = v.Number
	case HMVC:
		p.HalfMoveClock = v.Clock
	}
}
