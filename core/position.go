package core

// Position represents a chess position.
type Position struct {
	Board      Board
	SideToMove Color
	EnPassant  Square // The zero value indicates no en passant square.

	WhiteOO, WhiteOOO bool
	BlackOO, BlackOOO bool

	HalfMoveClock   int
	FullMoveCounter int // Starts at 1.
}

// NewPosition returns the starting position.
func NewPosition() *Position {
	var p Position

	p.Board.Set(WhiteRook, A1)
	p.Board.Set(WhiteKnight, B1)
	p.Board.Set(WhiteBishop, C1)
	p.Board.Set(WhiteQueen, D1)
	p.Board.Set(WhiteKing, E1)
	p.Board.Set(WhiteBishop, F1)
	p.Board.Set(WhiteKnight, G1)
	p.Board.Set(WhiteRook, H1)

	for sq := A2; sq <= H2; sq++ {
		p.Board.Set(WhitePawn, sq)
	}

	for sq := A7; sq <= H7; sq++ {
		p.Board.Set(BlackPawn, sq)
	}

	p.Board.Set(BlackRook, A8)
	p.Board.Set(BlackKnight, B8)
	p.Board.Set(BlackBishop, C8)
	p.Board.Set(BlackQueen, D8)
	p.Board.Set(BlackKing, E8)
	p.Board.Set(BlackBishop, F8)
	p.Board.Set(BlackKnight, G8)
	p.Board.Set(BlackRook, H8)

	p.WhiteOO = true
	p.WhiteOOO = true
	p.BlackOO = true
	p.BlackOOO = true

	p.FullMoveCounter = 1

	return &p
}
