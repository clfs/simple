package eval

import (
	"testing"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
)

func TestDoubledPawns(t *testing.T) {
	cases := []struct {
		in   string
		want core.Bitboard
	}{
		{
			in: fen.Starting,
		},
		{
			in:   "3qr1k1/5pp1/2pp1nnp/1b1Pp3/2N1P3/2P1P2P/1PBN2P1/Q4RK1 w - - 0 23",
			want: core.NewBitboard(core.E3, core.E4),
		},
	}

	for _, tc := range cases {
		got := doubledPawns(fen.MustDecode(tc.in))
		if tc.want != got {
			t.Errorf("%s:\nwant:\n%s\ngot:\n%s", tc.in, tc.want.Debug(), got.Debug())
		}
	}
}
