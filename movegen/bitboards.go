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
		for _, d := range knightDeltas {
			f := s.File() + d.f
			r := s.Rank() + d.r
			if f.Valid() && r.Valid() {
				knightAttacks[s].Set(core.NewSquare(f, r))
			}
		}
	}

	// Bishop attacks.
	bishopDeltas := []struct {
		f core.File
		r core.Rank
	}{
		{1, 1}, {-1, 1}, {-1, -1}, {1, -1},
	}
	for s := core.A1; s <= core.H8; s++ {
		for _, d := range bishopDeltas {
			f := s.File() + d.f
			r := s.Rank() + d.r
			for f.Valid() && r.Valid() {
				bishopAttacks[s].Set(core.NewSquare(f, r))
				f += d.f
				r += d.r
			}
		}
	}

	// Rook attacks.
	rookDeltas := []struct {
		f core.File
		r core.Rank
	}{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	for s := core.A1; s <= core.H8; s++ {
		for _, d := range rookDeltas {
			f := s.File() + d.f
			r := s.Rank() + d.r
			for f.Valid() && r.Valid() {
				rookAttacks[s].Set(core.NewSquare(f, r))
				f += d.f
				r += d.r
			}
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
		for _, delta := range kingDeltas {
			f := s.File() + delta.f
			r := s.Rank() + delta.r
			if f.Valid() && r.Valid() {
				kingAttacks[s].Set(core.NewSquare(f, r))
			}
		}
	}
}
