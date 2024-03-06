package reference

import (
	"slices"
	"testing"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
	"github.com/clfs/simple/encoding/pcn"
	"github.com/google/go-cmp/cmp"
)

// encodeMoves encodes moves as a sorted slice of PCN strings.
func encodeMoves(t *testing.T, moves []core.Move) []string {
	t.Helper()
	res := make([]string, len(moves))
	for i, m := range moves {
		res[i] = pcn.Encode(m)
	}
	slices.Sort(res)
	return res
}

func TestLegalMoves(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want []string
	}{
		{
			name: "starting position",
			in:   fen.Starting,
			want: []string{
				"a2a3", "a2a4", "b1a3", "b1c3", "b2b3", "b2b4", "c2c3", "c2c4",
				"d2d3", "d2d4", "e2e3", "e2e4", "f2f3", "f2f4", "g1f3", "g1h3",
				"g2g3", "g2g4", "h2h3", "h2h4",
			},
		},
		{
			name: "center bishop",
			in:   "k7/1R4p1/2B3P1/4b3/8/8/1P6/1K6 b - - 0 1",
			want: []string{
				"e5b2", "e5b8", "e5c3", "e5c7", "e5d4", "e5d6", "e5f4", "e5f6",
				"e5g3", "e5h2",
			},
		},
		{
			name: "castling",
			in:   "4k3/8/8/8/4p3/p2pPp1p/P2P1P1P/R3K2R w KQ - 0 1",
			want: []string{
				"a1b1", "a1c1", "a1d1", "e1c1", "e1d1", "e1f1", "e1g1", "h1f1",
				"h1g1",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			moves := LegalMoves(fen.MustDecode(tc.in))
			got := encodeMoves(t, moves)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("%q: mismatch (-want +got):\n%s", tc.in, diff)
			}
		})
	}
}
