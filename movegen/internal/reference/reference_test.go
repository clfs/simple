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
			"starting position",
			fen.Starting,
			[]string{
				"a2a3", "a2a4", "b1a3", "b1c3", "b2b3", "b2b4", "c2c3", "c2c4",
				"d2d3", "d2d4", "e2e3", "e2e4", "f2f3", "f2f4", "g1f3", "g1h3",
				"g2g3", "g2g4", "h2h3", "h2h4",
			},
		},
		{
			"center bishop",
			"k7/1R4p1/2B3P1/4b3/8/8/1P6/1K6 b - - 0 1",
			[]string{
				"e5b2", "e5b8", "e5c3", "e5c7", "e5d4", "e5d6", "e5f4", "e5f6",
				"e5g3", "e5h2",
			},
		},
		{
			"castling",
			"4k3/8/8/8/4p3/p2pPp1p/P2P1P1P/R3K2R w KQ - 0 1",
			[]string{
				"a1b1", "a1c1", "a1d1", "e1c1", "e1d1", "e1f1", "e1g1", "h1f1",
				"h1g1",
			},
		},
		{
			"escape check",
			"r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1",
			[]string{
				"b4c5", "c4c5", "d2d4", "f1f2", "f3d4", "g1h1",
			},
		},
		{
			"castling",
			"rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8",
			[]string{
				"a2a3", "a2a4", "b1a3", "b1c3", "b1d2", "b2b3", "b2b4", "c1d2",
				"c1e3", "c1f4", "c1g5", "c1h6", "c2c3", "c4a6", "c4b3", "c4b5",
				"c4d3", "c4d5", "c4e6", "c4f7", "d1d2", "d1d3", "d1d4", "d1d5",
				"d1d6", "d7c8b", "d7c8n", "d7c8q", "d7c8r", "e1d2", "e1f1",
				"e1f2", "e1g1", "e2c3", "e2d4", "e2f4", "e2g1", "e2g3", "g2g3",
				"g2g4", "h1f1", "h1g1", "h2h3", "h2h4",
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
