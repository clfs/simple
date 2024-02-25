package reference

import (
	"testing"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
	"github.com/clfs/simple/encoding/pcn"
	"github.com/google/go-cmp/cmp"
)

func encodeMoves(t *testing.T, moves []core.Move) []string {
	t.Helper()
	res := make([]string, len(moves))
	for i, m := range moves {
		res[i] = pcn.Encode(m)
	}
	return res
}

func TestLegalMoves(t *testing.T) {
	t.Skip("not implemented")

	cases := []struct {
		name string
		in   string
		want []string
	}{
		{
			name: "starting position",
			in:   fen.Starting,
			want: []string{
				"a2a3", "a2a4", "b2b3", "b2b4", "c2c3", "c2c4", "d2d3", "d2d4",
				"e2e3", "e2e4", "f2f3", "f2f4", "g2g3", "g2g4", "h2h3", "h2h4",
				"b1a3", "b1c3", "g1f3", "g1h3",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := encodeMoves(t, LegalMoves(fen.MustDecode(tc.in)))
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("%q: mismatch (-want +got):\n%s", tc.in, diff)
			}
		})
	}
}
