package fen

import (
	"bufio"
	"os"
	"testing"

	"github.com/clfs/simple/core"
	"github.com/google/go-cmp/cmp"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		moves []core.Move
		in    string
	}{
		{
			in: Starting,
		},
		{
			moves: []core.Move{
				{From: core.E2, To: core.E4},
			},
			in: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
		},
		{
			moves: []core.Move{
				{From: core.E2, To: core.E4},
				{From: core.E7, To: core.E5},
			},
			in: "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2",
		},
		{
			moves: []core.Move{
				{From: core.E2, To: core.E4},
				{From: core.E7, To: core.E5},
				{From: core.E1, To: core.E2},
			},
			in: "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPPKPPP/RNBQ1BNR b kq - 1 2",
		},
		{
			moves: []core.Move{
				{From: core.E2, To: core.E4},
				{From: core.E7, To: core.E5},
				{From: core.E1, To: core.E2},
				{From: core.E8, To: core.E7},
			},
			in: "rnbq1bnr/ppppkppp/8/4p3/4P3/8/PPPPKPPP/RNBQ1BNR w - - 2 3",
		},
	}

	for i, tc := range cases {
		want := core.NewPosition()
		for _, m := range tc.moves {
			want.Make(m)
		}

		got, err := Decode(tc.in)
		if err != nil {
			t.Errorf("#%d: Decode() error: %v", i, err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("#%d: Decode() mismatch (-want +got):\n%s", i, diff)
		}
	}
}

func readFENs(tb testing.TB, name string) []string {
	tb.Helper()

	f, err := os.Open(name)
	if err != nil {
		tb.Fatal(err)
	}
	defer f.Close()

	var lines []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); err != nil {
		tb.Fatal(err)
	}

	return lines
}

func TestDecode_Valid(t *testing.T) {
	for _, fen := range readFENs(t, "testdata/valid.fen") {
		if _, err := Decode(fen); err != nil {
			t.Errorf("Decode(%q) error: %v", fen, err)
		}
	}
}

func piecePlacements(t *testing.T, b core.Board) map[core.Square]core.Piece {
	t.Helper()

	m := make(map[core.Square]core.Piece)
	for s := core.A1; s <= core.H8; s++ {
		p, ok := b.Get(s)
		if ok {
			m[s] = p
		}
	}
	return m
}

func TestDecode_PiecePlacements(t *testing.T) {
	cases := []struct {
		in   string
		want map[core.Square]core.Piece
	}{
		{
			"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/1R2K2R b Kkq - 1 1",
			map[core.Square]core.Piece{
				core.A2: core.WhitePawn,
				core.A6: core.BlackBishop,
				core.A7: core.BlackPawn,
				core.A8: core.BlackRook,
				core.B1: core.WhiteRook,
				core.B2: core.WhitePawn,
				core.B4: core.BlackPawn,
				core.B6: core.BlackKnight,
				core.C2: core.WhitePawn,
				core.C3: core.WhiteKnight,
				core.C7: core.BlackPawn,
				core.D2: core.WhiteBishop,
				core.D5: core.WhitePawn,
				core.D7: core.BlackPawn,
				core.E1: core.WhiteKing,
				core.E2: core.WhiteBishop,
				core.E4: core.WhitePawn,
				core.E5: core.WhiteKnight,
				core.E6: core.BlackPawn,
				core.E7: core.BlackQueen,
				core.E8: core.BlackKing,
				core.F2: core.WhitePawn,
				core.F3: core.WhiteQueen,
				core.F6: core.BlackKnight,
				core.F7: core.BlackPawn,
				core.G2: core.WhitePawn,
				core.G6: core.BlackPawn,
				core.G7: core.BlackBishop,
				core.H1: core.WhiteRook,
				core.H2: core.WhitePawn,
				core.H3: core.BlackPawn,
				core.H8: core.BlackRook,
			},
		},
	}

	for i, tc := range cases {
		p, err := Decode(tc.in)
		if err != nil {
			t.Errorf("#%d: Decode() error: %v", i, err)
		}

		got := piecePlacements(t, p.Board)

		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("#%d: piecePlacements() mismatch (-want +got):\n%s", i, diff)
		}
	}
}
