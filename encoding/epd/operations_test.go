package epd

import "testing"

func TestOperation_Assemble(t *testing.T) {
	cases := []struct {
		in   Operation
		want RawOperation
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

	for i, c := range cases {
		if got := c.in.Assemble(); got != c.want {
			t.Errorf("#%d: want %v, got %v", i, c.want, got)
		}
	}
}
