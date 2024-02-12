package core

// Board contains piece placements.
type Board struct {
	whitePawns   Bitboard
	whiteKnights Bitboard
	whiteBishops Bitboard
	whiteRooks   Bitboard
	whiteQueens  Bitboard
	whiteKing    Bitboard
	blackPawns   Bitboard
	blackKnights Bitboard
	blackBishops Bitboard
	blackRooks   Bitboard
	blackQueens  Bitboard
	blackKing    Bitboard
}

// Set sets a piece down on the board.
func (b *Board) Set(p Piece, s Square) {
	if p.Color == White {
		switch p.Type {
		case Pawn:
			b.whitePawns.Set(s)
		case Knight:
			b.whiteKnights.Set(s)
		case Bishop:
			b.whiteBishops.Set(s)
		case Rook:
			b.whiteRooks.Set(s)
		case Queen:
			b.whiteQueens.Set(s)
		case King:
			b.whiteKing.Set(s)
		}
	} else {
		switch p.Type {
		case Pawn:
			b.blackPawns.Set(s)
		case Knight:
			b.blackKnights.Set(s)
		case Bishop:
			b.blackBishops.Set(s)
		case Rook:
			b.blackRooks.Set(s)
		case Queen:
			b.blackQueens.Set(s)
		case King:
			b.blackKing.Set(s)
		}
	}
}

// Get returns the piece on the given square, if any.
func (b *Board) Get(s Square) (Piece, bool) {
	switch {
	case b.whitePawns.Get(s):
		return Piece{Color: White, Type: Pawn}, true
	case b.whiteKnights.Get(s):
		return Piece{Color: White, Type: Knight}, true
	case b.whiteBishops.Get(s):
		return Piece{Color: White, Type: Bishop}, true
	case b.whiteRooks.Get(s):
		return Piece{Color: White, Type: Rook}, true
	case b.whiteQueens.Get(s):
		return Piece{Color: White, Type: Queen}, true
	case b.whiteKing.Get(s):
		return Piece{Color: White, Type: King}, true
	case b.blackPawns.Get(s):
		return Piece{Color: Black, Type: Pawn}, true
	case b.blackKnights.Get(s):
		return Piece{Color: Black, Type: Knight}, true
	case b.blackBishops.Get(s):
		return Piece{Color: Black, Type: Bishop}, true
	case b.blackRooks.Get(s):
		return Piece{Color: Black, Type: Rook}, true
	case b.blackQueens.Get(s):
		return Piece{Color: Black, Type: Queen}, true
	case b.blackKing.Get(s):
		return Piece{Color: Black, Type: King}, true
	default:
		return Piece{}, false
	}
}
