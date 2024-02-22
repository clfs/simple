package core

// A Move represents a chess move.
// For castling moves, From and To are the king's squares.
type Move struct {
	From, To  Square
	Promotion PieceType // The zero value indicates no promotion.
}
