package fen

import (
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
