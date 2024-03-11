package movegen

import (
	"slices"
	"strings"
	"testing"

	"github.com/clfs/simple/encoding/fen"
	"github.com/clfs/simple/encoding/pcn"
)

func FuzzLegalMoves(f *testing.F) {
	f.Add(fen.Starting, "e2e4 d7d5 e4d5")
	f.Fuzz(func(t *testing.T, pos, moves string) {
		p, err := fen.Decode(pos)
		if err != nil {
			t.Skip() // invalid FEN
		}

		for _, m := range strings.Split(moves, " ") {
			move, err := pcn.Decode(m)
			if err != nil {
				t.Skip() // invalid move
			}

			if !slices.Contains(LegalMoves(p), move) {
				t.Skip() // illegal move
			}

			p.Make(move)
		}
	})
}
