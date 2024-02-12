package core

import (
	"testing"
)

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

func TestBitboard_Clear(t *testing.T) {
	cases := []struct {
		bb   Bitboard
		sq   Square
		want Bitboard
	}{
		{0x1, A1, 0x0},
		{0x1, A2, 0x1},
		{0xffff_ffff_ffff_ffff, H8, 0x7fff_ffff_ffff_ffff},
	}
	for i, c := range cases {
		got := c.bb
		got.Clear(c.sq)
		if got != c.want {
			t.Errorf("%d: got %#x, want %#x", i, got, c.want)
		}
	}
}

func TestBitboard_Set(t *testing.T) {
	cases := []struct {
		bb   Bitboard
		sq   Square
		want Bitboard
	}{
		{0x0, A1, 0x1},
		{0x100, A2, 0x100},
		{0x7fff_ffff_ffff_ffff, H8, 0xffff_ffff_ffff_ffff},
	}
	for i, c := range cases {
		got := c.bb
		got.Set(c.sq)
		if got != c.want {
			t.Errorf("%d: got %#x, want %#x", i, got, c.want)
		}
	}
}

func TestBitboard_Get(t *testing.T) {
	cases := []struct {
		bb   Bitboard
		sq   Square
		want bool
	}{
		{0x1, A1, true},
		{0x2, A1, false},
		{0x100, A2, true},
		{0x7fff_ffff_ffff_ffff, H8, false},
	}
	for i, c := range cases {
		got := c.bb.Get(c.sq)
		if got != c.want {
			t.Errorf("%d: got %t, want %t", i, got, c.want)
		}
	}
}
