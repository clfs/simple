package eval

import (
	"testing"

	"github.com/clfs/simple/encoding/fen"
)

func TestEval(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{
			fen.Starting,
			0,
		},
		{
			"4k3/8/8/8/8/1R6/8/4K3 w - - 0 1",
			500,
		},
		{
			"4k3/8/8/8/8/1R6/8/4K3 b - - 0 1",
			-500,
		},
	}

	for _, tc := range cases {
		got := Eval(fen.MustDecode(tc.in))
		if tc.want != got {
			t.Errorf("%q: want %d, got %d", tc.in, tc.want, got)
		}
	}
}
