package core

import "fmt"

// A Square is a location on the chess board.
type Square uint64

// Square constants.
const (
	A1 Square = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

// NewSquare returns a new square.
func NewSquare(f File, r Rank) Square {
	return Square(int(f) + int(r)*8)
}

func (s Square) String() string {
	if s < A1 || s > H8 {
		return fmt.Sprintf("Square(%d)", s)
	}

	file := 'A' + s%8
	rank := '1' + s/8

	return fmt.Sprintf("%c%c", file, rank)
}

// Bitboard returns a bitboard with only this square set.
func (s Square) Bitboard() Bitboard {
	return 1 << s
}

// File returns a square's file.
func (s Square) File() File {
	return File(s % 8)
}

// Rank returns a square's rank.
func (s Square) Rank() Rank {
	return Rank(s / 8)
}

// Above returns the square above s, wrapping around if necessary.
func (s Square) Above() Square {
	return NewSquare(s.File(), s.Rank().Above())
}

// Below returns the square below s, wrapping around if necessary.
func (s Square) Below() Square {
	return NewSquare(s.File(), s.Rank().Below())
}

// Left returns the square to the left of s, wrapping around if necessary.
func (s Square) Left() Square {
	return NewSquare(s.File().Left(), s.Rank())
}

// Right returns the square to the right of s, wrapping around if necessary.
func (s Square) Right() Square {
	return NewSquare(s.File().Right(), s.Rank())
}
