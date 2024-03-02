package core

// Position represents a chess position.
type Position struct {
	Board      Board
	SideToMove Color
	EnPassant  Square // The zero value indicates no en passant square.

	WhiteOO, WhiteOOO bool
	BlackOO, BlackOOO bool

	HalfMoveClock  int
	FullMoveNumber int // Starts at 1.
}

// NewPosition returns the starting position.
func NewPosition() Position {
	var p Position

	p.Board.SetOnEmpty(WhiteRook, A1)
	p.Board.SetOnEmpty(WhiteKnight, B1)
	p.Board.SetOnEmpty(WhiteBishop, C1)
	p.Board.SetOnEmpty(WhiteQueen, D1)
	p.Board.SetOnEmpty(WhiteKing, E1)
	p.Board.SetOnEmpty(WhiteBishop, F1)
	p.Board.SetOnEmpty(WhiteKnight, G1)
	p.Board.SetOnEmpty(WhiteRook, H1)

	for sq := A2; sq <= H2; sq++ {
		p.Board.SetOnEmpty(WhitePawn, sq)
	}

	for sq := A7; sq <= H7; sq++ {
		p.Board.SetOnEmpty(BlackPawn, sq)
	}

	p.Board.SetOnEmpty(BlackRook, A8)
	p.Board.SetOnEmpty(BlackKnight, B8)
	p.Board.SetOnEmpty(BlackBishop, C8)
	p.Board.SetOnEmpty(BlackQueen, D8)
	p.Board.SetOnEmpty(BlackKing, E8)
	p.Board.SetOnEmpty(BlackBishop, F8)
	p.Board.SetOnEmpty(BlackKnight, G8)
	p.Board.SetOnEmpty(BlackRook, H8)

	p.WhiteOO = true
	p.WhiteOOO = true
	p.BlackOO = true
	p.BlackOOO = true

	p.FullMoveNumber = 1

	return p
}

// Make makes a move.
// It does not check for invalid moves.
func (p *Position) Make(m Move) {
	// Select the piece that we're going to move.
	heldPiece, _ := p.Board.Get(m.From)

	// Determine if the move is a capture.
	isCapture := p.Board.IsOccupied(m.To) ||
		(heldPiece.Type() == Pawn && m.To == p.EnPassant)

	// Move the piece.
	if m.Promotion == 0 {
		p.Board.Move(heldPiece, m.From, m.To)
	} else {
		p.Board.Promote(m.From, m.To, m.Promotion)
	}

	// Adjust pawn positions if capturing en passant.
	switch {
	case heldPiece == WhitePawn && m.To == p.EnPassant:
		p.Board.Clear(p.EnPassant.Below())
	case heldPiece == BlackPawn && m.To == p.EnPassant:
		p.Board.Clear(p.EnPassant.Above())
	}

	// Update castling rights.
	switch {
	case heldPiece == WhiteKing:
		p.WhiteOO, p.WhiteOOO = false, false
	case heldPiece == WhiteRook && (m.From == A1 || m.From == H1):
		p.WhiteOO, p.WhiteOOO = false, false
	case heldPiece == BlackKing:
		p.BlackOO, p.BlackOOO = false, false
	case heldPiece == BlackRook && (m.From == A8 || m.From == H8):
		p.BlackOO, p.BlackOOO = false, false
	}

	// Update the en passant square.
	switch {
	case heldPiece.Type() == Pawn && m.From.Rank() == Rank2 && m.To.Rank() == Rank4:
		p.EnPassant = m.From.Above()
	case heldPiece.Type() == Pawn && m.From.Rank() == Rank7 && m.To.Rank() == Rank5:
		p.EnPassant = m.From.Below()
	default:
		p.EnPassant = 0
	}

	// Adjust rook positions if castling.
	switch {
	case heldPiece.Type() == King && m.From == E1 && m.To == G1: // WhiteOO
		p.Board.MoveToEmpty(WhiteRook, H1, F1)
	case heldPiece.Type() == King && m.From == E1 && m.To == C1: // WhiteOOO
		p.Board.MoveToEmpty(WhiteRook, A1, D1)
	case heldPiece.Type() == King && m.From == E8 && m.To == G8: // BlackOO
		p.Board.MoveToEmpty(BlackRook, H8, F8)
	case heldPiece.Type() == King && m.From == E8 && m.To == C8: // BlackOOO
		p.Board.MoveToEmpty(BlackRook, A8, D8)
	}

	// Update the half move clock.
	if heldPiece.Type() == Pawn || isCapture {
		p.HalfMoveClock = 0
	} else {
		p.HalfMoveClock++
	}

	// Update the full move counter.
	if p.SideToMove == Black {
		p.FullMoveNumber++
	}

	// Switch sides.
	p.SideToMove = p.SideToMove.Other()
}

// FriendlyKing returns the location of the side to move's king.
func (p *Position) FriendlyKing() Square {
	if p.SideToMove == White {
		return p.Board.WhiteKing()
	}
	return p.Board.BlackKing()
}

// EnemyKing returns the location of the opponent's king.
func (p *Position) EnemyKing() Square {
	if p.SideToMove == White {
		return p.Board.BlackKing()
	}
	return p.Board.WhiteKing()
}
