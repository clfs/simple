package epd

import (
	"errors"
	"testing"
)

func TestOp_MarshalText(t *testing.T) {
	cases := []struct {
		op      Op
		want    string
		wantErr error
	}{
		{
			op:   Op{Opcode: "noop"},
			want: "noop;",
		},
	}

	for i, tc := range cases {
		got, err := tc.op.MarshalText()
		if !errors.Is(err, tc.wantErr) {
			t.Errorf("#%d: MarshalText() error mismatch: want %v, got %v", i, tc.wantErr, err)
		}
		if string(got) != tc.want {
			t.Errorf("#%d: MarshalText() mismatch: want %q, got %q", i, tc.want, got)
		}
	}
}
