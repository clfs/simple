package epd

import (
	"testing"

	"github.com/clfs/simple/encoding/fen"
	"github.com/google/go-cmp/cmp"
)

type encodeTestCase struct {
	in   string
	ops  []Op3
	want string
}

var encodeTestCases = []encodeTestCase{
	{
		fen.Starting,
		[]Op3{
			{Opcode: "noop"},
			{"c0", `"my comment"`},
			{"bm", "c4 Nf3"},
		},
		`rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - noop; c0 "my comment"; bm c4 Nf3;`,
	},
}

func TestEncode(t *testing.T) {
	for i, tc := range encodeTestCases {
		got, err := Encode(fen.MustDecode(tc.in), tc.ops)
		if err != nil {
			t.Errorf("#%d: Encode() failed: %v", i, err)
		}
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("#%d: Encode() mismatch: (-want +got):\n%s", i, diff)
		}
	}
}
