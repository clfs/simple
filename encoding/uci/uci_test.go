package uci

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUCI_UnmarshalText(t *testing.T) {
	cases := []struct {
		in      string
		want    UCI
		wantErr error
	}{
		{in: "uci", want: UCI{}},
		{in: "foo", wantErr: ErrUnmarshalWrongPrefix},
		{in: "uci foo", wantErr: ErrUnmarshalInvalidArgs},
	}

	for i, tc := range cases {
		var got UCI

		err := got.UnmarshalText([]byte(tc.in))
		if err != tc.wantErr {
			t.Errorf("#%d: wrong error: want %v, got %v", i, tc.wantErr, err)
		}

		if tc.wantErr != nil {
			continue
		}

		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("#%d: mismatch (-want +got):\n%s", i, diff)
		}
	}
}
