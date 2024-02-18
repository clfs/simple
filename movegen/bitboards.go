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

	for s := range core.H8 {
		f, r := s.File(), s.Rank()
		if f >= core.FileC && r <= core.Rank6 {
			knightAttacks[s].Set(s.Above().Above().Left())
		}
		if f <= core.FileG && r <= core.Rank6 {
			knightAttacks[s].Set(s.Above().Above().Right())
		}
		if f >= core.FileB && r <= core.Rank7 {
			knightAttacks[s].Set(s.Above().Left().Left())
		}
		if f <= core.FileG && r <= core.Rank7 {
			knightAttacks[s].Set(s.Above().Right().Right())
		}
		if f >= core.FileB && r >= core.Rank2 {
			knightAttacks[s].Set(s.Above().Left().Left())
		}
		if f <= core.FileG && r >= core.Rank2 {
			knightAttacks[s].Set(s.Below().Right().Right())
		}
		if f >= core.FileC && r >= core.Rank3 {
			knightAttacks[s].Set(s.Below().Below().Left())
		}
		if f <= core.FileG && r >= core.Rank3 {
			knightAttacks[s].Set(s.Below().Below().Right())
		}
	}

	// Bishop attacks.
	for s := range core.H8 {
		valid := func(f core.File, r core.Rank) bool {
			return f >= core.FileA && f <= core.FileH &&
				r >= core.Rank1 && r <= core.Rank8
		}
		f, r := s.File(), s.Rank()
		for f, r := f, r; valid(f, r); f, r = f-1, r-1 {
			bishopAttacks[s].Set(core.NewSquare(f, r))
		}
		for f, r := f, r; valid(f, r); f, r = f+1, r-1 {
			bishopAttacks[s].Set(core.NewSquare(f, r))
		}
		for f, r := f, r; valid(f, r); f, r = f-1, r+1 {
			bishopAttacks[s].Set(core.NewSquare(f, r))
		}
		for f, r := f, r; valid(f, r); f, r = f+1, r+1 {
			bishopAttacks[s].Set(core.NewSquare(f, r))
		}
	}

	// Rook attacks.
	for s := range core.H8 {
		valid := func(f core.File, r core.Rank) bool {
			return f >= core.FileA && f <= core.FileH &&
				r >= core.Rank1 && r <= core.Rank8
		}
		f, r := s.File(), s.Rank()
		for f, r := f, r; valid(f, r); r-- {
			rookAttacks[s].Set(core.NewSquare(f, r))
		}
		for f, r := f, r; valid(f, r); f++ {
			rookAttacks[s].Set(core.NewSquare(f, r))
		}
		for f, r := f, r; valid(f, r); r-- {
			rookAttacks[s].Set(core.NewSquare(f, r))
		}
		for f, r := f, r; valid(f, r); f-- {
			rookAttacks[s].Set(core.NewSquare(f, r))
		}
	}

	// Queen attacks.
	queenAttacks = bishopAttacks
	for s := range core.H8 {
		queenAttacks[s].With(rookAttacks[s])
	}

	// King attacks.
	for s := range core.H8 {
		f, r := s.File(), s.Rank()
		if r <= core.Rank7 {
			kingAttacks[s].Set(s.Above())
		}
		if r >= core.Rank2 {
			kingAttacks[s].Set(s.Below())
		}
		if f >= core.FileB {
			kingAttacks[s].Set(s.Left())
		}
		if f <= core.FileG {
			kingAttacks[s].Set(s.Right())
		}
		if f >= core.FileB && r <= core.Rank7 {
			kingAttacks[s].Set(s.Above().Left())
		}
		if f <= core.FileG && r <= core.Rank7 {
			kingAttacks[s].Set(s.Above().Right())
		}
		if f >= core.FileB && r >= core.Rank2 {
			kingAttacks[s].Set(s.Below().Left())
		}
		if f <= core.FileG && r >= core.Rank2 {
			kingAttacks[s].Set(s.Below().Right())
		}
	}
}
