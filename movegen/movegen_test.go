package movegen

import (
	"testing"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
	"github.com/google/go-cmp/cmp"
)

func TestMoves(t *testing.T) {
	t.Skip("not implemented")
	cases := []struct {
		in   string
		want []core.Move
	}{
		{
			fen.Starting,
			[]core.Move{
				// Pawn moves.
				{From: core.A2, To: core.A3},
				{From: core.A2, To: core.A4},
				{From: core.B2, To: core.B3},
				{From: core.B2, To: core.B4},
				{From: core.C2, To: core.C3},
				{From: core.C2, To: core.C4},
				{From: core.D2, To: core.D3},
				{From: core.D2, To: core.D4},
				{From: core.E2, To: core.E3},
				{From: core.E2, To: core.E4},
				{From: core.F2, To: core.F3},
				{From: core.F2, To: core.F4},
				{From: core.G2, To: core.G3},
				{From: core.G2, To: core.G4},
				{From: core.H2, To: core.H3},
				{From: core.H2, To: core.H4},
				// Knight moves.
				{From: core.B1, To: core.A3},
				{From: core.B1, To: core.C3},
				{From: core.G1, To: core.F3},
				{From: core.G1, To: core.H3},
			},
		},
	}
	for _, tc := range cases {
		got := Moves(fen.MustDecode(tc.in))
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("Moves(%q) mismatch (-want +got):\n%s", tc.in, diff)
		}
	}
}
