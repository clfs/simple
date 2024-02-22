package core

import "fmt"

// A Piece represents a chess piece.
type Piece uint64

// White piece constants.
const (
	WhitePawn Piece = iota
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

// NewPiece returns a new piece.
func NewPiece(c Color, pt PieceType) Piece {
	return Piece(c.Uint64()*6 + uint64(pt))
}

func (p Piece) Valid() bool {
	return p <= BlackKing
}

var pieceNames = [...]string{
	"WhitePawn",
	"WhiteKnight",
	"WhiteBishop",
	"WhiteRook",
	"WhiteQueen",
	"WhiteKing",
	"BlackPawn",
	"BlackKnight",
	"BlackBishop",
	"BlackRook",
	"BlackQueen",
	"BlackKing",
}

func (p Piece) String() string {
	if p.Valid() {
		return pieceNames[p]
	}
	return fmt.Sprintf("Piece(%d)", p)
}

// Type returns a piece's type.
func (p Piece) Type() PieceType {
	return PieceType(p % 6)
}

// Color returns a piece's color.
func (p Piece) Color() Color {
	return p >= BlackPawn
}
