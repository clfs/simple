package core

import "testing"

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
