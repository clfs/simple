package epd

import (
	"fmt"
	"testing"

	"github.com/clfs/simple/encoding/fen"
	"github.com/google/go-cmp/cmp"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		in      string
		want    string
		wantOps []Operation
		wantErr string
	}{
		{
			in:   `r1bqk2r/p1pp1ppp/2p2n2/8/1b2P3/2N5/PPP2PPP/R1BQKB1R w KQkq - bm Bd3; id "Crafty Test Pos.28"; c0 "DB/GK Philadelphia 1996, Game 5, move 7W (Bd3)";`,
			want: "r1bqk2r/p1pp1ppp/2p2n2/8/1b2P3/2N5/PPP2PPP/R1BQKB1R w KQkq - 0 1",
			wantOps: []Operation{
				BM{Moves: []string{"Bd3"}},
				ID{ID: "Crafty Test Pos.28"},
				Comment{Comment: "DB/GK Philadelphia 1996, Game 5, move 7W (Bd3)"},
			},
		},
		{
			in:   `8/3r4/pr1Pk1p1/8/7P/6P1/3R3K/5R2 w - - bm Re2+; id "arasan21.16"; c0 "Aldiga (Brainfish 091016)-Knight-king (Komodo 10 64-bit), playchess.com 2016";`,
			want: "8/3r4/pr1Pk1p1/8/7P/6P1/3R3K/5R2 w - - 0 1",
			wantOps: []Operation{
				BM{Moves: []string{"Re2+"}},
				ID{ID: "arasan21.16"},
				Comment{Comment: "Aldiga (Brainfish 091016)-Knight-king (Komodo 10 64-bit), playchess.com 2016"},
			},
		},
		{
			in:   `3r1rk1/1p3pnp/p3pBp1/1qPpP3/1P1P2R1/P2Q3R/6PP/6K1 w - - bm Rxh7;c0 "Mate in 7 moves";id "BT2630-14";`,
			want: "3r1rk1/1p3pnp/p3pBp1/1qPpP3/1P1P2R1/P2Q3R/6PP/6K1 w - - 0 1",
			wantOps: []Operation{
				BM{Moves: []string{"Rxh7"}},
				Comment{Comment: "Mate in 7 moves"},
				ID{ID: "BT2630-14"},
			},
		},
		{
			in:   `4k3/8/P7/8/8/8/8/4K3 w - - bm a7; c0 "semicolon ; inside comment";`,
			want: "4k3/8/P7/8/8/8/8/4K3 w - - 0 1",
			wantOps: []Operation{
				BM{Moves: []string{"a7"}},
				Comment{Comment: "semicolon ; inside comment"},
			},
		},
		{
			in:   `4k3/8/P7/8/8/8/8/4K3 w - -`,
			want: `4k3/8/P7/8/8/8/8/4K3 w - - 0 1`,
		},
		{
			in:      `4k3/8/P7/8/8/8/8/4K3 w -`,
			wantErr: "too few fields: 3",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			p, ops, err := Decode(tc.in)

			// Early exit if an error is expected.
			if tc.wantErr != "" {
				if err == nil {
					t.Errorf("#%d: wrong error: want %q, got <nil>", i, tc.wantErr)
				} else if err.Error() != tc.wantErr {
					t.Errorf("#%d: wrong error: want %q, got %q", i, tc.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("#%d: wrong error: want <nil>, got %v", i, err)
			}

			got := fen.Encode(p)
			if got != tc.want {
				t.Errorf("#%d: wrong position: want %q, got %q", i, tc.want, got)
			}

			if diff := cmp.Diff(tc.wantOps, ops); diff != "" {
				t.Errorf("#%d: wrong ops (-want +got):\n%s", i, diff)
			}
		})
	}
}
