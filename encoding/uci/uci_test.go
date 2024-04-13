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
		{in: " ", wantErr: ErrUnmarshalEmptyMessage},
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

func TestUCI_MarshalText(t *testing.T) {
	cases := []struct {
		in      UCI
		want    string
		wantErr error
	}{
		{in: UCI{}, want: "uci"},
	}

	for i, tc := range cases {
		got, err := tc.in.MarshalText()
		if err != tc.wantErr {
			t.Errorf("#%d: wrong error: want %v, got %v", i, tc.wantErr, err)
		}

		if tc.wantErr != nil {
			continue
		}

		if tc.want != string(got) {
			t.Errorf("#%d: want %q, got %q", i, tc.want, got)
		}
	}
}
