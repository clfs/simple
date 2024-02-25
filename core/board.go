package core

// Board contains piece placements.
type Board [12]Bitboard

// Set places a piece on a square.
// If there is already a piece on the square, that piece is removed.
// If the square is known to be empty, SetOnEmpty is faster.
func (b *Board) Set(p Piece, s Square) {
	for i := range b {
		b[i].Clear(s)
	}
	b[p].Set(s)
}

// SetOnEmpty places a piece on an empty square.
func (b *Board) SetOnEmpty(p Piece, s Square) {
	b[p].Set(s)
}

// Get returns the piece on the given square, if any.
func (b *Board) Get(s Square) (Piece, bool) {
	for i, bb := range b {
		if bb.Get(s) {
			return Piece(i), true
		}
	}
	return 0, false
}

// IsEmpty returns true if the square is empty.
func (b *Board) IsEmpty(s Square) bool {
	for _, bb := range b {
		if bb.Get(s) {
			return false
		}
	}
	return true
}

// Move moves a piece to a square.
// If there is already a piece at the destination, that piece is removed.
// If the destination is known to be empty, MoveToEmpty is faster.
func (b *Board) Move(p Piece, from, to Square) {
	b[p].Clear(from)
	for i := range b {
		b[i].Clear(to)
	}
	b[p].Set(to)
}

// MoveToEmpty moves a piece to an empty square.
func (b *Board) MoveToEmpty(p Piece, from, to Square) {
	b[p].Clear(from)
	b[p].Set(to)
}

// Promote moves a pawn to a square, promoting it.
func (b *Board) Promote(from, to Square, p PieceType) {
	if to.Rank() == Rank8 { // White
		b[WhitePawn].Clear(from)
		b[NewPiece(White, p)].Set(to)
	} else { // Black
		b[BlackPawn].Clear(from)
		b[NewPiece(Black, p)].Set(to)
	}
}

// WhitePieces returns the location of all white pieces.
func (b *Board) WhitePieces() Bitboard {
	var bb Bitboard
	for i := WhitePawn; i <= WhiteKing; i++ {
		bb.With(b[i])
	}
	return bb
}

// BlackPieces returns the location of all black pieces.
func (b *Board) BlackPieces() Bitboard {
	var bb Bitboard
	for i := BlackPawn; i <= BlackKing; i++ {
		bb.With(b[i])
	}
	return bb
}
