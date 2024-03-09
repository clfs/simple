package core

import "testing"

func TestPosition_Make_ShortCastle(t *testing.T) {
	var p Position

	p.Board.Set(WhiteKing, E1)
	p.Board.Set(WhiteRook, H1)
	p.Board.Set(BlackKing, E8)
	p.Board.Set(BlackRook, H8)

	p.WhiteOO = true
	p.WhiteOOO = true
	p.BlackOO = true
	p.BlackOOO = true

	p.Make(Move{From: E1, To: G1})
	p.Make(Move{From: E8, To: G8})

	if p.WhiteOO || p.WhiteOOO || p.BlackOO || p.BlackOOO {
		t.Errorf(
			"expected no castling rights, got WhiteOO=%v WhiteOOO=%v BlackOO=%v BlackOOO=%v",
			p.WhiteOO, p.WhiteOOO, p.BlackOO, p.BlackOOO,
		)
	}

	piece, ok := p.Board.Get(G1)
	if !ok || piece != WhiteKing {
		t.Errorf("expected WhiteKing on G1, got %s, %t", piece, ok)
	}
	piece, ok = p.Board.Get(G8)
	if !ok || piece != BlackKing {
		t.Errorf("expected BlackKing on G8, got %s, %t", piece, ok)
	}
}

func TestPosition_Make_LongCastle(t *testing.T) {
	var p Position

	p.Board.Set(WhiteKing, E1)
	p.Board.Set(WhiteRook, A1)
	p.Board.Set(BlackKing, E8)
	p.Board.Set(BlackRook, A8)

	p.WhiteOO = true
	p.WhiteOOO = true
	p.BlackOO = true
	p.BlackOOO = true

	p.Make(Move{From: E1, To: C1})
	p.Make(Move{From: E8, To: C8})

	if p.WhiteOO || p.WhiteOOO || p.BlackOO || p.BlackOOO {
		t.Errorf(
			"expected no castling rights, got WhiteOO=%v WhiteOOO=%v BlackOO=%v BlackOOO=%v",
			p.WhiteOO, p.WhiteOOO, p.BlackOO, p.BlackOOO,
		)
	}

	piece, ok := p.Board.Get(C1)
	if !ok || piece != WhiteKing {
		t.Errorf("expected WhiteKing on C1, got %s, %t", piece, ok)
	}
	piece, ok = p.Board.Get(C8)
	if !ok || piece != BlackKing {
		t.Errorf("expected BlackKing on C8, got %s, %t", piece, ok)
	}
}

func TestPosition_Make_Promotion(t *testing.T) {
	var p Position

	p.Board.Set(WhitePawn, A7)
	p.Make(Move{From: A7, To: A8, Promotion: Queen})

	piece, ok := p.Board.Get(A8)
	if !ok || piece != WhiteQueen {
		t.Errorf("expected WhiteQueen on A8, got %s, %t", piece, ok)
	}
}

func TestPosition_Make_RookRights(t *testing.T) {
	var p Position

	p.Board.Set(WhiteRook, A1)
	p.Board.Set(BlackRook, A8)

	p.WhiteOO = true
	p.WhiteOOO = true
	p.BlackOO = true
	p.BlackOOO = true

	p.Make(Move{From: A1, To: A3})
	p.Make(Move{From: A8, To: A6})

	if !p.WhiteOO {
		t.Errorf("expected WhiteOO=true, got %t", p.WhiteOO)
	}
	if p.WhiteOOO {
		t.Errorf("expected WhiteOOO=false, got %t", p.WhiteOOO)
	}
	if !p.BlackOO {
		t.Errorf("expected BlackOO=true, got %t", p.BlackOO)
	}
	if p.BlackOOO {
		t.Errorf("expected BlackOOO=false, got %t", p.BlackOOO)
	}
}

func TestPosition_Make_EnPassant(t *testing.T) {
	p := NewPosition()

	p.Make(Move{From: A2, To: A4})
	if p.EnPassant != A3 {
		t.Errorf("expected en passant square A3, got %s", p.EnPassant)
	}

	p.Make(Move{From: A7, To: A5})
	if p.EnPassant != A6 {
		t.Errorf("expected en passant square A6, got %s", p.EnPassant)
	}
}

func TestPosition_Make_Capture(t *testing.T) {
	var p Position

	p.Board.Set(WhitePawn, A2)
	p.Board.Set(BlackPawn, B3)

	p.Make(Move{From: A2, To: B3})
	piece, ok := p.Board.Get(B3)
	if !ok || piece != WhitePawn {
		t.Errorf("expected WhitePawn on B3, got %s, %t", piece, ok)
	}
}

func TestPosition_Make_WhiteEPCapture(t *testing.T) {
	var p Position

	p.Board.Set(WhitePawn, A5)
	p.Board.Set(BlackPawn, B5)
	p.EnPassant = B6

	p.Make(Move{From: A5, To: B6})
	piece, ok := p.Board.Get(B6)
	if !ok || piece != WhitePawn {
		t.Errorf("expected WhitePawn on B6, got %s, %t", piece, ok)
	}
	piece, ok = p.Board.Get(B5)
	if ok {
		t.Errorf("expected no BlackPawn on B5, got %s, %t", piece, ok)
	}
}

func TestPosition_Make_BlackEPCapture(t *testing.T) {
	var p Position

	p.Board.Set(BlackPawn, A4)
	p.Board.Set(WhitePawn, B4)
	p.EnPassant = B3

	p.Make(Move{From: A4, To: B3})
	piece, ok := p.Board.Get(B3)
	if !ok || piece != BlackPawn {
		t.Errorf("expected BlackPawn on B3, got %s, %t", piece, ok)
	}
	piece, ok = p.Board.Get(B4)
	if ok {
		t.Errorf("expected no WhitePawn on B4, got %s, %t", piece, ok)
	}
}
