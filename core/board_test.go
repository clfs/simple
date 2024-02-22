package core

import "testing"

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

func TestBoard_Set(t *testing.T) {
	var b Board

	var (
		p = BlackRook
		s = C6
	)

	b.Set(p, s)

	got, ok := b.Get(s)
	if !ok {
		t.Errorf("%s empty", s)
	}
	if got != p {
		t.Errorf("want %s, got %s", p, got)
	}
}

func TestBoard_Move(t *testing.T) {
	var b Board

	var (
		p    = BlackRook
		from = A2
		to   = A4
	)

	b.Set(p, from)

	b.Move(p, from, to)

	_, ok := b.Get(from)
	if ok {
		t.Errorf("%s not empty", from)
	}

	got, ok := b.Get(to)
	if !ok {
		t.Errorf("%s empty", to)
	}
	if got != p {
		t.Errorf("want %s, got %s", p, got)
	}
}

func TestBoard_MoveToEmpty(t *testing.T) {
	var b Board

	var (
		p    = BlackRook
		from = A2
		to   = A4
	)

	b.Set(p, from)

	b.MoveToEmpty(p, from, to)

	_, ok := b.Get(from)
	if ok {
		t.Errorf("%s not empty", from)
	}

	got, ok := b.Get(to)
	if !ok {
		t.Errorf("%s empty", to)
	}
	if got != p {
		t.Errorf("want %s, got %s", p, got)
	}
}

func TestBoard_Promote_White(t *testing.T) {
	var b Board

	var (
		from = A7
		to   = A8
		p    = Queen
	)

	b.Set(WhitePawn, from)
	b.Promote(from, to, p)

	_, ok := b.Get(from)
	if ok {
		t.Errorf("%s not empty", from)
	}

	got, ok := b.Get(to)
	if !ok {
		t.Errorf("%s empty", to)
	}
	if got != NewPiece(White, p) {
		t.Errorf("want %s, got %s", NewPiece(White, p), got)
	}
}

func TestBoard_Promote_Black(t *testing.T) {
	var b Board

	var (
		from = A2
		to   = A1
		p    = Queen
	)

	b.Set(BlackPawn, from)
	b.Promote(from, to, p)

	_, ok := b.Get(from)
	if ok {
		t.Errorf("%s not empty", from)
	}

	got, ok := b.Get(to)
	if !ok {
		t.Errorf("%s empty", to)
	}
	if got != NewPiece(Black, p) {
		t.Errorf("want %s, got %s", NewPiece(Black, p), got)
	}
}
