package core

import "math/bits"

// A Bitboard represents each square on the board as a bit.
// The LSB is a1, and the MSB is h8.
type Bitboard uint64

// Bitboard returns a bitboard with only this square set.
func (s Square) Bitboard() Bitboard {
	return 1 << s
}

// Clear clears the given square to 0.
func (b *Bitboard) Clear(s Square) {
	*b &= ^(1 << s)
}

// Set sets the given square to 1.
func (b *Bitboard) Set(s Square) {
	*b |= 1 << s
}

// Get returns true if the given square is set to 1.
func (b *Bitboard) Get(s Square) bool {
	return *b&(1<<s) != 0
}

// FlipV flips the bitboard horizontally.
func (b *Bitboard) FlipV() {
	*b = Bitboard(bits.ReverseBytes64(uint64(*b)))
}

// With sets all squares set in the other bitboard.
func (b *Bitboard) With(other Bitboard) {
	*b |= other
}
