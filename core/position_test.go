package core

import "testing"

func TestBoard(t *testing.T) {
	var b Board

	b.Set(WhitePawn, A2)
	if !b.IsSet(A2) {
		t.Error("A2 not set")
	}

	p, ok := b.Get(A2)
	if !ok {
		t.Error("A2 not gotten")
	}
	if p != WhitePawn {
		t.Errorf("A2 should be %s, but got %s", WhitePawn, p)
	}

	p, ok = b.Get(A3)
	if ok {
		t.Error("A3 not empty")
	}

	b.Clear(A2)
	if !b.IsEmpty(A2) {
		t.Error("A2 not empty")
	}
}
