package core

import "testing"

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
