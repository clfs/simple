package movegen

import (
	"github.com/clfs/simple/core"
)

var (
	bishopAttacks [64]core.Bitboard
	rookAttacks   [64]core.Bitboard
	queenAttacks  [64]core.Bitboard
)

func init() {
	type delta struct {
		f core.File
		r core.Rank
	}

	var (
		bishopDeltas = []delta{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
		rookDeltas   = []delta{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		queenDeltas  = []delta{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	)

	for s := core.A1; s <= core.H8; s++ {
		f, r := s.File(), s.Rank()

		// Bishop attacks.
		for _, d := range bishopDeltas {
			f, r := f+d.f, r+d.r
			for f.Valid() && r.Valid() {
				bishopAttacks[s].Set(core.NewSquare(f, r))
				f, r = f+d.f, r+d.r
			}
		}

		// Rook attacks.
		for _, d := range rookDeltas {
			f, r := f+d.f, r+d.r
			for f.Valid() && r.Valid() {
				rookAttacks[s].Set(core.NewSquare(f, r))
				f, r = f+d.f, r+d.r
			}
		}

		// Queen attacks.
		for _, d := range queenDeltas {
			f, r := f+d.f, r+d.r
			for f.Valid() && r.Valid() {
				queenAttacks[s].Set(core.NewSquare(f, r))
				f, r = f+d.f, r+d.r
			}
		}
	}
}
