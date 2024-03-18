package epd

import (
	"fmt"
	"testing"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
	"github.com/google/go-cmp/cmp"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		in      string
		want    ExtendedPosition
		wantErr bool
	}{
		{
			in: Starting,
			want: ExtendedPosition{
				Position: core.NewPosition(),
			},
		},
		{
			in: Starting + " hmvc 4; fmvn 3;",
			want: ExtendedPosition{
				Position: fen.MustDecode("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 4 3"),
			},
		},
		{
			in:      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq",
			wantErr: true,
		},
		{
			in:      "rnbqkbnr/pppppppp/8/8/8/8/RNBQKBNR w KQkq - ",
			wantErr: true,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := Decode(tc.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("Decode(%q) returned error: %v", tc.in, err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Decode(%q) returned unexpected result (-want +got):\n%s", tc.in, diff)
			}
		})
	}
}
