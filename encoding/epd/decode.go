package epd

import (
	"fmt"
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
		applyOp(&p, op)
	}

	return p, ops, nil
}

func decodeOps(_ string) ([]Op, error) {
	// TODO: Implement.
	return nil, nil
}

func applyOp(p *core.Position, op Op) {
	// TODO: Implement.
}
