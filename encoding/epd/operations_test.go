package epd

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOperation_Assemble(t *testing.T) {
	cases := []struct {
		in      Operation
		want    RawOperation
		wantErr string
	}{
		{
			in:   ACN{Nodes: 42},
			want: RawOperation{"acn", "42"},
		},
		{
			in:   ACS{Seconds: 42},
			want: RawOperation{"acs", "42"},
		},
	}

	for i, tc := range cases {
		got, err := tc.in.Assemble()

		// Early continue if an error is expected.
		if tc.wantErr != "" {
			if err == nil {
				t.Errorf("#%d: wrong error: want %q, got <nil>", i, tc.wantErr)
			} else if err.Error() != tc.wantErr {
				t.Errorf("#%d: wrong error: want %q, got %q", i, tc.wantErr, err)
			}
			continue
		}

		if err != nil {
			t.Errorf("#%d: wrong error: want <nil>, got %v", i, err)
		}

		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("#%d: wrong RawOperation (-want +got):\n%s", i, diff)
		}
	}
}
