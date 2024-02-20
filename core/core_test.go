package core

import "testing"

func TestColor_String(t *testing.T) {
	cases := []struct {
		color Color
		want  string
	}{
		{White, "White"},
		{Black, "Black"},
	}
	for i, c := range cases {
		got := c.color.String()
		if got != c.want {
			t.Errorf("%d: got %q, want %q", i, got, c.want)
		}
	}
}

func TestColor_Other(t *testing.T) {
	cases := []struct {
		color Color
		want  Color
	}{
		{White, Black},
		{Black, White},
	}
	for _, c := range cases {
		got := c.color.Other()
		if got != c.want {
			t.Errorf("%s.Other() = %s, want %s", c.color, got, c.want)
		}
	}
}

func TestPieceType_String(t *testing.T) {
	cases := []struct {
		pt   PieceType
		want string
	}{
		{Pawn, "Pawn"},
		{Knight, "Knight"},
		{Bishop, "Bishop"},
		{Rook, "Rook"},
		{Queen, "Queen"},
		{King, "King"},
		{PieceType(42), "PieceType(42)"},
	}
	for i, c := range cases {
		got := c.pt.String()
		if got != c.want {
			t.Errorf("%d: got %q, want %q", i, got, c.want)
		}
	}
}

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

func TestFile_String(t *testing.T) {
	cases := []struct {
		file File
		want string
	}{
		{FileA, "FileA"},
		{FileB, "FileB"},
		{FileC, "FileC"},
		{FileD, "FileD"},
		{FileE, "FileE"},
		{FileF, "FileF"},
		{FileG, "FileG"},
		{FileH, "FileH"},
		{File(42), "File(42)"},
	}
	for i, c := range cases {
		got := c.file.String()
		if got != c.want {
			t.Errorf("%d: got %q, want %q", i, got, c.want)
		}
	}
}

func TestRank_String(t *testing.T) {
	cases := []struct {
		rank Rank
		want string
	}{
		{Rank1, "Rank1"},
		{Rank2, "Rank2"},
		{Rank3, "Rank3"},
		{Rank4, "Rank4"},
		{Rank5, "Rank5"},
		{Rank6, "Rank6"},
		{Rank7, "Rank7"},
		{Rank8, "Rank8"},
		{Rank(42), "Rank(42)"},
	}
	for i, c := range cases {
		got := c.rank.String()
		if got != c.want {
			t.Errorf("%d: got %q, want %q", i, got, c.want)
		}
	}
}

func TestNewSquare(t *testing.T) {
	cases := []struct {
		file File
		rank Rank
		want Square
	}{
		{FileA, Rank1, A1},
		{FileD, Rank4, D4},
		{FileH, Rank8, H8},
	}
	for i, c := range cases {
		got := NewSquare(c.file, c.rank)
		if got != c.want {
			t.Errorf("%d: got %d, want %d", i, got, c.want)
		}
	}
}

func TestSquare_String(t *testing.T) {
	cases := []struct {
		sq   Square
		want string
	}{
		{A1, "A1"},
		{D4, "D4"},
		{H8, "H8"},
		{Square(1000), "Square(1000)"},
	}
	for i, c := range cases {
		got := c.sq.String()
		if got != c.want {
			t.Errorf("%d: got %q, want %q", i, got, c.want)
		}
	}
}

func TestSquare_Bitboard(t *testing.T) {
	cases := []struct {
		sq   Square
		want Bitboard
	}{
		{A1, 0x1},
		{B1, 0x2},
		{H8, 0x8000_0000_0000_0000},
	}
	for _, c := range cases {
		got := c.sq.Bitboard()
		if got != c.want {
			t.Errorf("%s.Bitboard() == %#x, want %#x", c.sq, got, c.want)
		}
	}
}

func TestSquare_File(t *testing.T) {
	cases := []struct {
		sq   Square
		want File
	}{
		{A1, FileA},
		{D4, FileD},
		{H8, FileH},
	}
	for i, c := range cases {
		got := c.sq.File()
		if got != c.want {
			t.Errorf("%d: got %d, want %d", i, got, c.want)
		}
	}
}

func TestSquare_Rank(t *testing.T) {
	cases := []struct {
		sq   Square
		want Rank
	}{
		{A1, Rank1},
		{D4, Rank4},
		{H8, Rank8},
	}
	for i, c := range cases {
		got := c.sq.Rank()
		if got != c.want {
			t.Errorf("%d: got %d, want %d", i, got, c.want)
		}
	}
}

func TestBoard(t *testing.T) {
	var b Board

	b.SetOnEmpty(WhitePawn, A2)

	p, ok := b.Get(A2)
	if !ok {
		t.Error("A2 not gotten")
	}
	if p != WhitePawn {
		t.Errorf("A2 should be %s, but got %s", WhitePawn, p)
	}

	_, ok = b.Get(A3)
	if ok {
		t.Error("A3 not empty")
	}
}

func TestSquare_Above(t *testing.T) {
	cases := []struct {
		sq   Square
		want Square
	}{
		{A1, A2},
		{D4, D5},
	}
	for i, c := range cases {
		got := c.sq.Above()
		if got != c.want {
			t.Errorf("%d: got %s, want %s", i, got, c.want)
		}
	}
}

func TestSquare_Below(t *testing.T) {
	cases := []struct {
		sq   Square
		want Square
	}{
		{A2, A1},
		{D5, D4},
	}
	for i, c := range cases {
		got := c.sq.Below()
		if got != c.want {
			t.Errorf("%d: got %s, want %s", i, got, c.want)
		}
	}
}

func TestSquare_Left(t *testing.T) {
	cases := []struct {
		sq   Square
		want Square
	}{
		{B2, A2},
		{D5, C5},
	}
	for i, c := range cases {
		got := c.sq.Left()
		if got != c.want {
			t.Errorf("%d: got %s, want %s", i, got, c.want)
		}
	}
}

func TestSquare_Right(t *testing.T) {
	cases := []struct {
		sq   Square
		want Square
	}{
		{A1, B1},
		{D4, E4},
	}
	for i, c := range cases {
		got := c.sq.Right()
		if got != c.want {
			t.Errorf("%d: got %s, want %s", i, got, c.want)
		}
	}
}
