package core

import "fmt"

// A PieceType is a type of piece.
type PieceType uint64

// Piece type constants.
const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

var pieceTypeNames = [...]string{
	"Pawn",
	"Knight",
	"Bishop",
	"Rook",
	"Queen",
	"King",
}

func (p PieceType) Valid() bool {
	return p <= King
}

func (p PieceType) String() string {
	if p.Valid() {
		return pieceTypeNames[p]
	}
	return fmt.Sprintf("PieceType(%d)", p)
}
