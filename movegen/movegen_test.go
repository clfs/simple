package movegen

import (
	"testing"

	"github.com/clfs/simple/encoding/fen"
	"github.com/clfs/simple/encoding/pcn"
	"github.com/google/go-cmp/cmp"
)

func TestDivide(t *testing.T) {
	cases := []struct {
		in    string
		depth int
		want  map[string]int
	}{
		{
			fen.Starting,
			5,
			map[string]int{
				"a2a3": 181046,
				"a2a4": 217832,
				"b1a3": 198572,
				"b1c3": 234656,
				"b2b3": 215255,
				"b2b4": 216145,
				"c2c3": 222861,
				"c2c4": 240082,
				"d2d3": 328511,
				"d2d4": 361790,
				"e2e3": 402988,
				"e2e4": 405385,
				"f2f3": 178889,
				"f2f4": 198473,
				"g1f3": 233491,
				"g1h3": 198502,
				"g2g3": 217210,
				"g2g4": 214048,
				"h2h3": 181044,
				"h2h4": 218829,
			},
		},
	}

	for _, tc := range cases {
		got := make(map[string]int)
		for k, v := range Divide(fen.MustDecode(tc.in), tc.depth) {
			got[pcn.Encode(k)] = v
		}
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("%q: -want +got\n%s", tc.in, diff)
		}
	}
}

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
