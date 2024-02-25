// Package pcn implements encoding and decoding of pure coordinate notation
// (PCN).
package pcn

import (
	"fmt"
	"strings"

	"github.com/clfs/simple/core"
)

var encodePromotion = map[core.PieceType]string{
	core.Knight: "n",
	core.Bishop: "b",
	core.Rook:   "r",
	core.Queen:  "q",
}

// Encode encodes a move as a PCN string.
func Encode(m core.Move) string {
	f := strings.ToLower(m.From.String())
	t := strings.ToLower(m.To.String())
	p := encodePromotion[m.Promotion]
	return f + t + p
}

func decodeSquare(s string) (core.Square, bool) {
	if len(s) != 2 {
		return 0, false
	}

	f := core.File(s[0] - 'a')
	r := core.Rank(s[1] - '1')

	if !f.Valid() || !r.Valid() {
		return 0, false
	}
	return core.NewSquare(f, r), true
}

var decodePromotion = map[byte]core.PieceType{
	'n': core.Knight,
	'b': core.Bishop,
	'r': core.Rook,
	'q': core.Queen,
}

// Decode decodes a PCN string and returns the move it represents.
func Decode(s string) (core.Move, error) {
	if n := len(s); n < 4 || n > 5 {
		return core.Move{}, fmt.Errorf("invalid length: %d", n)
	}

	var m core.Move

	if sq, ok := decodeSquare(s[:2]); !ok {
		return core.Move{}, fmt.Errorf("invalid source square: %s", s[:2])
	} else {
		m.From = sq
	}

	if sq, ok := decodeSquare(s[2:4]); !ok {
		return core.Move{}, fmt.Errorf("invalid target square: %s", s[2:4])
	} else {
		m.To = sq
	}

	if len(s) == 5 {
		if p, ok := decodePromotion[s[4]]; !ok {
			return core.Move{}, fmt.Errorf("invalid promotion: %c", s[4])
		} else {
			m.Promotion = p
		}
	}

	return m, nil
}

// MustDecode is like Decode but panics if the PCN is invalid.
func MustDecode(s string) core.Move {
	m, err := Decode(s)
	if err != nil {
		panic(err)
	}
	return m
}
