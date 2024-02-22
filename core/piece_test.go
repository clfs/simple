package core

import "testing"

func TestNewPiece(t *testing.T) {
	cases := []struct {
		color Color
		pt    PieceType
		want  Piece
	}{
		{White, Pawn, WhitePawn},
		{Black, Knight, BlackKnight},
	}
	for _, c := range cases {
		got := NewPiece(c.color, c.pt)
		if got != c.want {
			t.Errorf("NewPiece(%s, %s) = %s, want %s", c.color, c.pt, got, c.want)
		}
	}
}

func TestPiece_String(t *testing.T) {
	cases := []struct {
		p    Piece
		want string
	}{
		{WhitePawn, "WhitePawn"},
		{WhiteKnight, "WhiteKnight"},
		{Piece(42), "Piece(42)"},
	}
	for i, c := range cases {
		got := c.p.String()
		if got != c.want {
			t.Errorf("%d: got %q, want %q", i, got, c.want)
		}
	}
}

func TestPiece_Type(t *testing.T) {
	cases := []struct {
		p    Piece
		want PieceType
	}{
		{WhitePawn, Pawn},
		{WhiteKnight, Knight},
		{WhiteBishop, Bishop},
		{WhiteRook, Rook},
		{WhiteQueen, Queen},
		{WhiteKing, King},
		{BlackPawn, Pawn},
		{BlackKnight, Knight},
		{BlackBishop, Bishop},
		{BlackRook, Rook},
		{BlackQueen, Queen},
		{BlackKing, King},
	}
	for _, c := range cases {
		got := c.p.Type()
		if got != c.want {
			t.Errorf("%s.Type() = %s, want %s", c.p, got, c.want)
		}
	}
}

func TestPiece_Color(t *testing.T) {
	cases := []struct {
		p    Piece
		want Color
	}{
		{WhitePawn, White},
		{WhiteKnight, White},
		{WhiteBishop, White},
		{WhiteRook, White},
		{WhiteQueen, White},
		{WhiteKing, White},
		{BlackPawn, Black},
		{BlackKnight, Black},
		{BlackBishop, Black},
		{BlackRook, Black},
		{BlackQueen, Black},
		{BlackKing, Black},
	}
	for _, c := range cases {
		got := c.p.Color()
		if got != c.want {
			t.Errorf("%s.Color() = %s, want %s", c.p, got, c.want)
		}
	}
}
