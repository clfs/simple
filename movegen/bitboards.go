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
	// White pawn pushes.
	for s := core.A2; s <= core.H7; s++ {
		whitePawnPushes[s].Set(s.Above())
		if s.Rank() == core.Rank2 {
			whitePawnPushes[s].Set(s.Above().Above())
		}
	}

	// Black pawn pushes.
	blackPawnPushes = whitePawnPushes
	for i := range blackPawnPushes {
		blackPawnPushes[i].FlipV()
	}

	// White pawn attacks.
	for s := core.A2; s <= core.H7; s++ {
		f := s.File()
		if f >= core.FileB {
			whitePawnAttacks[s].Set(s.Above().Left())
		}
		if f <= core.FileG {
			whitePawnAttacks[s].Set(s.Above().Right())
		}
	}

	// Black pawn attacks.
	blackPawnAttacks = whitePawnAttacks
	for i := range blackPawnAttacks {
		blackPawnAttacks[i].FlipV()
	}

	// Knight attacks.
	knightDeltas := []struct {
		f core.File
		r core.Rank
	}{
		{2, 1}, {1, 2}, {-2, 1}, {-1, 2},
		{2, -1}, {1, -2}, {-2, -1}, {-1, -2},
	}
	for s := core.A1; s <= core.H8; s++ {
		f, r := s.File(), s.Rank()
		for _, delta := range knightDeltas {
			ff, rr := f+delta.f, r+delta.r
			if ff.Valid() && rr.Valid() {
				knightAttacks[s].Set(core.NewSquare(ff, rr))
			}
		}
	}

	// Bishop attacks.
	for s := core.A1; s <= core.H8; s++ {
		f, r := s.File(), s.Rank()
		for f, r := f-1, r-1; f.Valid() && r.Valid(); f, r = f-1, r-1 {
			bishopAttacks[s].Set(core.NewSquare(f, r))
		}
		for f, r := f+1, r-1; f.Valid() && r.Valid(); f, r = f+1, r-1 {
			bishopAttacks[s].Set(core.NewSquare(f, r))
		}
		for f, r := f-1, r+1; f.Valid() && r.Valid(); f, r = f-1, r+1 {
			bishopAttacks[s].Set(core.NewSquare(f, r))
		}
		for f, r := f+1, r+1; f.Valid() && r.Valid(); f, r = f+1, r+1 {
			bishopAttacks[s].Set(core.NewSquare(f, r))
		}
	}

	// Rook attacks.
	for s := core.A1; s <= core.H8; s++ {
		f, r := s.File(), s.Rank()
		for f := f - 1; f.Valid(); f-- {
			rookAttacks[s].Set(core.NewSquare(f, r))
		}
		for f := f + 1; f.Valid(); f++ {
			rookAttacks[s].Set(core.NewSquare(f, r))
		}
		for r := r - 1; r.Valid(); r-- {
			rookAttacks[s].Set(core.NewSquare(f, r))
		}
		for r := r + 1; r.Valid(); r++ {
			rookAttacks[s].Set(core.NewSquare(f, r))
		}
	}

	// Queen attacks.
	queenAttacks = bishopAttacks
	for s := core.A1; s <= core.H8; s++ {
		queenAttacks[s].With(rookAttacks[s])
	}

	// King attacks.
	kingDeltas := []struct {
		f core.File
		r core.Rank
	}{
		{1, 0}, {1, 1}, {0, 1}, {-1, 1},
		{-1, 0}, {-1, -1}, {0, -1}, {1, -1},
	}
	for s := core.A1; s <= core.H8; s++ {
		f, r := s.File(), s.Rank()
		for _, delta := range kingDeltas {
			ff, rr := f+delta.f, r+delta.r
			if ff.Valid() && rr.Valid() {
				kingAttacks[s].Set(core.NewSquare(ff, rr))
			}
		}
	}
}
