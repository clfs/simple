// Package core implements basic chess functionality.
package core

import "fmt"

// A Color is white or black.
type Color bool

// Color constants.
const (
	White Color = false
	Black Color = true
)

func (c Color) String() string {
	if c {
		return "Black"
	}
	return "White"
}

// Other returns the other color.
func (c Color) Other() Color {
	return !c
}

// Uint64 returns 0 for white and 1 for black.
func (c Color) Uint64() uint64 {
	if c {
		return 1
	}
	return 0
}

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

func (p PieceType) String() string {
	switch p {
	case Pawn:
		return "Pawn"
	case Knight:
		return "Knight"
	case Bishop:
		return "Bishop"
	case Rook:
		return "Rook"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return fmt.Sprintf("PieceType(%d)", p)
	}
}

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

var pieceNames = [12]string{
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

// A File is a column on the chess board.
type File uint64

// File constants.
const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

func (f File) String() string {
	switch f {
	case FileA:
		return "FileA"
	case FileB:
		return "FileB"
	case FileC:
		return "FileC"
	case FileD:
		return "FileD"
	case FileE:
		return "FileE"
	case FileF:
		return "FileF"
	case FileG:
		return "FileG"
	case FileH:
		return "FileH"
	default:
		return fmt.Sprintf("File(%d)", f)
	}
}

func (f File) Valid() bool {
	return f >= FileA && f <= FileH
}

// A Rank is a row on the chess board.
type Rank uint64

// Rank constants.
const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

func (r Rank) String() string {
	switch r {
	case Rank1:
		return "Rank1"
	case Rank2:
		return "Rank2"
	case Rank3:
		return "Rank3"
	case Rank4:
		return "Rank4"
	case Rank5:
		return "Rank5"
	case Rank6:
		return "Rank6"
	case Rank7:
		return "Rank7"
	case Rank8:
		return "Rank8"
	default:
		return fmt.Sprintf("Rank(%d)", r)
	}
}

func (r Rank) Valid() bool {
	return r >= Rank1 && r <= Rank8
}

// A Square is a location on the chess board.
type Square uint64

// Square constants.
const (
	A1 Square = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

// NewSquare returns a new square.
func NewSquare(f File, r Rank) Square {
	return Square(int(f) + int(r)*8)
}

func (s Square) String() string {
	if s < A1 || s > H8 {
		return fmt.Sprintf("Square(%d)", s)
	}

	file := 'A' + s%8
	rank := '1' + s/8

	return fmt.Sprintf("%c%c", file, rank)
}

// Bitboard returns a bitboard with only this square set.
func (s Square) Bitboard() Bitboard {
	return 1 << s
}

// File returns a square's file.
func (s Square) File() File {
	return File(s % 8)
}

// Rank returns a square's rank.
func (s Square) Rank() Rank {
	return Rank(s / 8)
}

// Above returns the square above s.
func (s Square) Above() Square {
	return s + 8
}

// Below returns the square below s.
func (s Square) Below() Square {
	return s - 8
}

// Left returns the square to the left of s.
func (s Square) Left() Square {
	return s - 1
}

// Right returns the square to the right of s.
func (s Square) Right() Square {
	return s + 1
}

// A Move represents a chess move.
// For castling moves, From and To are the king's squares.
type Move struct {
	From, To  Square
	Promotion PieceType // The zero value indicates no promotion.
}

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
