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
