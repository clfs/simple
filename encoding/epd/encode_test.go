package epd

import (
	"testing"

	"github.com/clfs/simple/encoding/fen"
	"github.com/google/go-cmp/cmp"
)

type encodeTestCase struct {
	in   string
	ops  []Op
	want string
}

var encodeTestCases = []encodeTestCase{
	{
		fen.Starting,
		[]Op{
			{"foo", "bar"},
			{"c0", `"my comment"`},
			{"bm", "c4 Nf3"},
		},
		`rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - foo bar; c0 "my comment"; bm c4 Nf3;`,
	},
}

func TestEncode(t *testing.T) {
	for i, tc := range encodeTestCases {
		got := Encode(fen.MustDecode(tc.in), tc.ops)
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("#%d: Encode() mismatch: (-want +got):\n%s", i, diff)
		}
	}
}
