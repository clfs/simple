package core

import (
	"bytes"
	"math/bits"
)

// A Bitboard represents each square on the board as a bit.
// The LSB is a1, and the MSB is h8.
type Bitboard uint64

// NewBitboard returns a new bitboard with the given squares set.
func NewBitboard(s ...Square) Bitboard {
	var b Bitboard
	for _, sq := range s {
		b.Set(sq)
	}
	return b
}

// Debug returns an 8x8 representation of the bitboard.
func (b *Bitboard) Debug() string {
	var buf bytes.Buffer
	for r := Rank8; r.Valid(); r-- {
		for f := FileA; f.Valid(); f++ {
			if b.Get(NewSquare(f, r)) {
				buf.WriteByte('X')
			} else {
				buf.WriteByte('.')
			}
		}
		buf.WriteByte('\n')
	}
	return buf.String()
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

// Count returns the number of squares set to 1.
func (b *Bitboard) Count() int {
	return bits.OnesCount64(uint64(*b))
}

// FlipV flips the bitboard horizontally.
func (b *Bitboard) FlipV() {
	*b = Bitboard(bits.ReverseBytes64(uint64(*b)))
}

// With sets all squares set in the other bitboard.
func (b *Bitboard) With(other Bitboard) {
	*b |= other
}
