package fen

import (
	"testing"

	"github.com/clfs/simple/core"
	"github.com/google/go-cmp/cmp"
)

func TestEncode(t *testing.T) {
	cases := []struct {
		p     core.Position
		moves []core.Move
		want  string
	}{
		{
			p:    core.NewPosition(),
			want: Starting,
		},
		{
			p: core.NewPosition(),
			moves: []core.Move{
				{From: core.E2, To: core.E4},
			},
			want: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
		},
		{
			p: core.NewPosition(),
			moves: []core.Move{
				{From: core.E2, To: core.E4},
				{From: core.E7, To: core.E5},
			},
			want: "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2",
		},
		{
			p: core.NewPosition(),
			moves: []core.Move{
				{From: core.E2, To: core.E4},
				{From: core.E7, To: core.E5},
				{From: core.E1, To: core.E2},
			},
			want: "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPPKPPP/RNBQ1BNR b kq - 1 2",
		},
		{
			p: core.NewPosition(),
			moves: []core.Move{
				{From: core.E2, To: core.E4},
				{From: core.E7, To: core.E5},
				{From: core.E1, To: core.E2},
				{From: core.E8, To: core.E7},
			},
			want: "rnbq1bnr/ppppkppp/8/4p3/4P3/8/PPPPKPPP/RNBQ1BNR w - - 2 3",
		},
		{
			p: MustDecode("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"),
			moves: []core.Move{
				{From: core.A1, To: core.B1},
			},
			want: "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/1R2K2R b Kkq - 1 1",
		},
		{
			p: MustDecode("r3k2r/Pppp1ppp/1b3nbN/nPB5/B1P1P3/q4N2/Pp1P2PP/R2Q1RK1 b kq - 1 1"),
			moves: []core.Move{
				{From: core.B2, To: core.A1, Promotion: core.Bishop},
			},
			want: "r3k2r/Pppp1ppp/1b3nbN/nPB5/B1P1P3/q4N2/P2P2PP/b2Q1RK1 w kq - 0 2",
		},
		{
			p: MustDecode("rnbq1k1r/pp1Pbppp/2p5/8/2B5/P7/1PP1NnPP/RNBQK2R b KQ - 0 8"),
			moves: []core.Move{
				{From: core.F2, To: core.H1},
			},
			want: "rnbq1k1r/pp1Pbppp/2p5/8/2B5/P7/1PP1N1PP/RNBQK2n w Q - 0 9",
		},
	}
	for i, tc := range cases {
		for _, m := range tc.moves {
			tc.p.Make(m)
		}
		got := Encode(tc.p)
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("#%d: (-want +got):\n%s", i, diff)
		}
	}
}
