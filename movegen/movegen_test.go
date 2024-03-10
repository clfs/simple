package movegen

import (
	"testing"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
	"github.com/clfs/simple/encoding/pcn"
	"github.com/google/go-cmp/cmp"
)

func TestPerft(t *testing.T) {
	cases := []struct {
		in   string
		want []int
	}{
		{
			fen.Starting,
			[]int{1, 20, 400, 8902, 197281},
		},
		{
			// https://www.chessprogramming.org/Perft_Results#Position_2
			"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
			[]int{1, 48, 2039, 97862},
		},
		{
			// https://www.chessprogramming.org/Perft_Results#Position_3
			"8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1",
			[]int{1, 14, 191, 2812, 43238},
		},
		{
			// https://www.chessprogramming.org/Perft_Results#Position_4
			"r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1",
			[]int{1, 6, 264, 9467, 422333},
		},
		{
			// https://www.chessprogramming.org/Perft_Results#Position_5
			"rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8",
			[]int{1, 44, 1486, 62379},
		},
		{
			// https://www.chessprogramming.org/Perft_Results#Position_6
			"r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10",
			[]int{1, 46, 2079, 89890},
		},
	}

	for _, tc := range cases {
		p := fen.MustDecode(tc.in)
		for i, want := range tc.want {
			if got := Perft(p, i); got != want {
				t.Errorf("%q at depth %d: got %d, want %d", tc.in, i, got, want)
			}
		}
	}
}

// encodeMoves encodes moves as a sorted slice of PCN strings.
func encodeMoveMap(t *testing.T, m map[core.Move]int) map[string]int {
	t.Helper()
	res := make(map[string]int, len(m))
	for move, count := range m {
		res[pcn.Encode(move)] = count
	}
	return res
}

func TestDivide(t *testing.T) {
	cases := []struct {
		in    string
		depth int
		want  map[string]int
	}{
		{
			"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
			1,
			map[string]int{
				"a1b1": 1,
				"a1c1": 1,
				"a1d1": 1,
				"a2a3": 1,
				"a2a4": 1,
				"b2b3": 1,
				"c3a4": 1,
				"c3b1": 1,
				"c3b5": 1,
				"c3d1": 1,
				"d2c1": 1,
				"d2e3": 1,
				"d2f4": 1,
				"d2g5": 1,
				"d2h6": 1,
				"d5d6": 1,
				"d5e6": 1,
				"e1c1": 1,
				"e1d1": 1,
				"e1f1": 1,
				"e1g1": 1,
				"e2a6": 1,
				"e2b5": 1,
				"e2c4": 1,
				"e2d1": 1,
				"e2d3": 1,
				"e2f1": 1,
				"e5c4": 1,
				"e5c6": 1,
				"e5d3": 1,
				"e5d7": 1,
				"e5f7": 1,
				"e5g4": 1,
				"e5g6": 1,
				"f3d3": 1,
				"f3e3": 1,
				"f3f4": 1,
				"f3f5": 1,
				"f3f6": 1,
				"f3g3": 1,
				"f3g4": 1,
				"f3h3": 1,
				"f3h5": 1,
				"g2g3": 1,
				"g2g4": 1,
				"g2h3": 1,
				"h1f1": 1,
				"h1g1": 1,
			},
		},
		{
			"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
			2,
			map[string]int{
				"a1b1": 43,
				"a1c1": 43,
				"a1d1": 43,
				"a2a3": 44,
				"a2a4": 44,
				"b2b3": 42,
				"c3a4": 42,
				"c3b1": 42,
				"c3b5": 39,
				"c3d1": 42,
				"d2c1": 43,
				"d2e3": 43,
				"d2f4": 43,
				"d2g5": 42,
				"d2h6": 41,
				"d5d6": 41,
				"d5e6": 46,
				"e1c1": 43,
				"e1d1": 43,
				"e1f1": 43,
				"e1g1": 43,
				"e2a6": 36,
				"e2b5": 39,
				"e2c4": 41,
				"e2d1": 44,
				"e2d3": 42,
				"e2f1": 44,
				"e5c4": 42,
				"e5c6": 41,
				"e5d3": 43,
				"e5d7": 45,
				"e5f7": 44,
				"e5g4": 44,
				"e5g6": 42,
				"f3d3": 42,
				"f3e3": 43,
				"f3f4": 43,
				"f3f5": 45,
				"f3f6": 39,
				"f3g3": 43,
				"f3g4": 43,
				"f3h3": 43,
				"f3h5": 43,
				"g2g3": 42,
				"g2g4": 42,
				"g2h3": 43,
				"h1f1": 43,
				"h1g1": 43,
			},
		},
		{
			"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
			3,
			map[string]int{
				"a2a3": 2186,
				"b2b3": 1964,
				"g2g3": 1882,
				"d5d6": 1991,
				"a2a4": 2149,
				"g2g4": 1843,
				"g2h3": 1970,
				"d5e6": 2241,
				"c3b1": 2038,
				"c3d1": 2040,
				"c3a4": 2203,
				"c3b5": 2138,
				"e5d3": 1803,
				"e5c4": 1880,
				"e5g4": 1878,
				"e5c6": 2027,
				"e5g6": 1997,
				"e5d7": 2124,
				"e5f7": 2080,
				"d2c1": 1963,
				"d2e3": 2136,
				"d2f4": 2000,
				"d2g5": 2134,
				"d2h6": 2019,
				"e2d1": 1733,
				"e2f1": 2060,
				"e2d3": 2050,
				"e2c4": 2082,
				"e2b5": 2057,
				"e2a6": 1907,
				"a1b1": 1969,
				"a1c1": 1968,
				"a1d1": 1885,
				"h1f1": 1929,
				"h1g1": 2013,
				"f3d3": 2005,
				"f3e3": 2174,
				"f3g3": 2214,
				"f3h3": 2360,
				"f3f4": 2132,
				"f3g4": 2169,
				"f3f5": 2396,
				"f3h5": 2267,
				"f3f6": 2111,
				"e1d1": 1894,
				"e1f1": 1855,
				"e1g1": 2059,
				"e1c1": 1887,
			},
		},
		{
			"r3k2r/p1pNqpb1/bn2pnp1/3P4/1p2P3/2N2Q1p/PPPBBPPP/R3K2R b KQkq - 0 1",
			1,
			map[string]int{
				"a6b5": 1,
				"a6b7": 1,
				"a6c4": 1,
				"a6c8": 1,
				"a6d3": 1,
				"a6e2": 1,
				"a8b8": 1,
				"a8c8": 1,
				"a8d8": 1,
				"b4b3": 1,
				"b4c3": 1,
				"b6a4": 1,
				"b6c4": 1,
				"b6c8": 1,
				"b6d5": 1,
				"b6d7": 1,
				"c7c5": 1,
				"c7c6": 1,
				"e6d5": 1,
				"e6e5": 1,
				"e7c5": 1,
				"e7d6": 1,
				"e7d7": 1,
				"e7d8": 1,
				"e7f8": 1,
				"e8c8": 1,
				"e8d7": 1,
				"e8d8": 1,
				"f6d5": 1,
				"f6d7": 1,
				"f6e4": 1,
				"f6g4": 1,
				"f6g8": 1,
				"f6h5": 1,
				"f6h7": 1,
				"g6g5": 1,
				"g7f8": 1,
				"g7h6": 1,
				"h3g2": 1,
				"h8f8": 1,
				"h8g8": 1,
				"h8h4": 1,
				"h8h5": 1,
				"h8h6": 1,
				"h8h7": 1,
			},
		},
		{
			"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/1R2K2R b Kkq - 1 1",
			1,
			map[string]int{
				"a6b5": 1,
				"a6b7": 1,
				"a6c4": 1,
				"a6c8": 1,
				"a6d3": 1,
				"a6e2": 1,
				"a8b8": 1,
				"a8c8": 1,
				"a8d8": 1,
				"b4b3": 1,
				"b4c3": 1,
				"b6a4": 1,
				"b6c4": 1,
				"b6c8": 1,
				"b6d5": 1,
				"c7c5": 1,
				"c7c6": 1,
				"d7d6": 1,
				"e6d5": 1,
				"e7c5": 1,
				"e7d6": 1,
				"e7d8": 1,
				"e7f8": 1,
				"e8c8": 1,
				"e8d8": 1,
				"e8f8": 1,
				"e8g8": 1,
				"f6d5": 1,
				"f6e4": 1,
				"f6g4": 1,
				"f6g8": 1,
				"f6h5": 1,
				"f6h7": 1,
				"g6g5": 1,
				"g7f8": 1,
				"g7h6": 1,
				"h3g2": 1,
				"h8f8": 1,
				"h8g8": 1,
				"h8h4": 1,
				"h8h5": 1,
				"h8h6": 1,
				"h8h7": 1,
			},
		},
	}

	for _, tc := range cases {
		p := fen.MustDecode(tc.in)
		got := encodeMoveMap(t, Divide(p, tc.depth))
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("%q at depth %d: (-want +got)\n%s", tc.in, tc.depth, diff)
		}
	}
}
