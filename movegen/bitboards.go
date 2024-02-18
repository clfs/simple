package movegen

import (
	"github.com/clfs/simple/core"
)

var (
	whitePawnPushes  [64]core.Bitboard
	blackPawnPushes  [64]core.Bitboard
	whitePawnAttacks [64]core.Bitboard
	blackPawnAttacks [64]core.Bitboard
	knightAttacks    [64]core.Bitboard
	bishopAttacks    [64]core.Bitboard
	rookAttacks      [64]core.Bitboard
	queenAttacks     [64]core.Bitboard
	kingAttacks      [64]core.Bitboard
)

func init() {
	type delta struct {
		f core.File
		r core.Rank
	}

	var (
		knightDeltas = []delta{{2, 1}, {1, 2}, {-2, 1}, {-1, 2}, {2, -1}, {1, -2}, {-2, -1}, {-1, -2}}
		bishopDeltas = []delta{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
		rookDeltas   = []delta{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		queenDeltas  = []delta{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
		kingDeltas   = []delta{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	)

	for s := core.A1; s <= core.H8; s++ {
		f, r := s.File(), s.Rank()

		// Pawn attacks and pushes.
		if r != core.Rank1 && r != core.Rank8 {
			whitePawnPushes[s].Set(s.Above())
			blackPawnPushes[s].Set(s.Below())
			if f >= core.FileB {
				whitePawnAttacks[s].Set(s.Above().Left())
				blackPawnAttacks[s].Set(s.Below().Left())
			}
			if f <= core.FileG {
				whitePawnAttacks[s].Set(s.Above().Right())
				blackPawnAttacks[s].Set(s.Below().Right())
			}
			if r == core.Rank2 {
				whitePawnPushes[s].Set(s.Above().Above())
			}
			if r == core.Rank7 {
				blackPawnPushes[s].Set(s.Below().Below())
			}
		}

		// Knight attacks.
		for _, d := range knightDeltas {
			f := f + d.f
			r := r + d.r
			if f.Valid() && r.Valid() {
				knightAttacks[s].Set(core.NewSquare(f, r))
			}
		}

		// Bishop attacks.
		for _, d := range bishopDeltas {
			f := f + d.f
			r := r + d.r
			for f.Valid() && r.Valid() {
				bishopAttacks[s].Set(core.NewSquare(f, r))
				f += d.f
				r += d.r
			}
		}

		// Rook attacks.
		for _, d := range rookDeltas {
			f := f + d.f
			r := r + d.r
			for f.Valid() && r.Valid() {
				rookAttacks[s].Set(core.NewSquare(f, r))
				f += d.f
				r += d.r
			}
		}

		// Queen attacks.
		for _, d := range queenDeltas {
			f := f + d.f
			r := r + d.r
			for f.Valid() && r.Valid() {
				queenAttacks[s].Set(core.NewSquare(f, r))
				f += d.f
				r += d.r
			}
		}

		// King attacks.
		for _, d := range kingDeltas {
			f := f + d.f
			r := r + d.r
			if f.Valid() && r.Valid() {
				kingAttacks[s].Set(core.NewSquare(f, r))
			}
		}
	}
}
