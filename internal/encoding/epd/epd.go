// Package epd implements decoding for the Extended Position Description format.
package epd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
)

// Starting is the starting position in EPD format.
const Starting = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -"

// ExtendedPosition represents an EPD position.
type ExtendedPosition struct {
	Position core.Position
	Comment  string
	Perft    []int
}

// MustDecode is like Decode but panics if the EPD is invalid.
func MustDecode(s string) ExtendedPosition {
	res, err := Decode(s)
	if err != nil {
		panic(err)
	}
	return res
}

// Decode decodes an EPD string.
func Decode(s string) (ExtendedPosition, error) {
	var res ExtendedPosition

	fields := strings.SplitN(s, " ", 5)
	if n := len(fields); n < 4 {
		return ExtendedPosition{}, fmt.Errorf("too few fields: %d", n)
	}

	// Decode the position.
	pseudoFEN := strings.Join(fields[:4], " ") + " 0 1"
	p, err := fen.Decode(pseudoFEN)
	if err != nil {
		return ExtendedPosition{}, err
	}

	res.Position = p

	// Return early if there are no operations.
	if len(fields) == 4 {
		return res, nil
	}

	// Decode the operations.
	for ops := fields[4]; len(ops) > 0; {
		var (
			op string
			ok bool
		)

		op, ops, ok = strings.Cut(ops, ";")
		if !ok {
			return ExtendedPosition{}, fmt.Errorf("missing semicolon")
		}

		op = strings.TrimSpace(op)

		opcode, args, _ := strings.Cut(op, " ")

		switch opcode {
		case "fmvn":
			n, err := strconv.Atoi(args)
			if err != nil {
				return ExtendedPosition{}, fmt.Errorf("invalid fmvn")
			}
			res.Position.FullMoveNumber = n
		case "hmvc":
			n, err := strconv.Atoi(args)
			if err != nil {
				return ExtendedPosition{}, fmt.Errorf("invalid hmvc")
			}
			res.Position.HalfMoveClock = n
		case "c0":
			args, ok = strings.CutPrefix(args, `"`)
			if !ok {
				return ExtendedPosition{}, fmt.Errorf("invalid c0")
			}
			args, ok = strings.CutSuffix(args, `"`)
			if !ok {
				return ExtendedPosition{}, fmt.Errorf("invalid c0")
			}
			res.Comment = args
		case "Perft":
			depths := strings.Split(args, " ")
			for _, d := range depths {
				n, err := strconv.Atoi(d)
				if err != nil {
					return ExtendedPosition{}, fmt.Errorf("invalid Perft")
				}
				res.Perft = append(res.Perft, n)
			}
		default:
			return ExtendedPosition{}, fmt.Errorf("unknown opcode: %s", opcode)
		}
	}

	return res, nil
}
