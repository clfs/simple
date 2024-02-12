package core

// Board contains piece placements.
type Board struct {
	pieces [12]Bitboard
}

// IsSet returns true if the given square is occupied.
func (b *Board) IsSet(s Square) bool {
	for _, bb := range b.pieces {
		if bb.Get(s) {
			return true
		}
	}
	return false
}

// IsEmpty returns true if the given square is empty.
func (b *Board) IsEmpty(s Square) bool {
	return !b.IsSet(s)
}

// Clear clears a square, removing any piece that may be there.
func (b *Board) Clear(s Square) {
	for i := range b.pieces {
		b.pieces[i].Clear(s)
	}
}

// Set sets a piece down on the board.
func (b *Board) Set(p Piece, s Square) {
	b.pieces[p].Set(s)
}

// Get returns the piece on the given square, if any.
func (b *Board) Get(s Square) (Piece, bool) {
	for i, bb := range b.pieces {
		if bb.Get(s) {
			return Piece(i), true
		}
	}
	return 0, false
}
