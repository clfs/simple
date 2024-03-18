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
		wantErr string
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
			wantErr: "too few fields: 3",
		},
		{
			in:      "rnbqkbnr/pppppppp/8/8/8/8/RNBQKBNR w KQkq - ",
			wantErr: "invalid number of board rows: 7",
		},
		{
			in:      Starting + " hmvc 4",
			wantErr: "missing semicolon",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := Decode(tc.in)
			if tc.wantErr != "" {
				if err == nil || err.Error() != tc.wantErr {
					t.Errorf("#d: wrong error: got %v, want %v", err, tc.wantErr)
				}
				return // early return if we expected an error
			}
			if err != nil {
				t.Errorf("#d: unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("#d: mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
