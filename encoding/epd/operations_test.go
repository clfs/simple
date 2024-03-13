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
			in:   RawOperation{Opcode: "foo", Args: "bar"},
			want: RawOperation{Opcode: "foo", Args: "bar"},
		},
		{
			in:      RawOperation{Args: "bar"},
			wantErr: "empty opcode",
		},
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

		if tc.wantErr != "" {
			if diff := cmp.Diff(tc.wantErr, err.Error()); diff != "" {
				t.Errorf("#%d: wrong error (-want +got):\n%s", i, diff)
			}
			continue
		}

		if err != nil {
			t.Errorf("#%d: error: got %v", i, err)
		}

		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("#%d: wrong RawOperation (-want +got):\n%s", i, diff)
		}
	}
}
