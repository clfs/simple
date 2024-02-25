package pcn

import (
	"fmt"
	"testing"

	"github.com/clfs/simple/core"
)

func ExampleEncode() {
	moves := []core.Move{
		// Double pawn push.
		{From: core.B2, To: core.B4},
		// Queen promotion.
		{From: core.F2, To: core.F1, Promotion: core.Queen},
		// White short castle.
		{From: core.E1, To: core.G1},
	}
	for _, m := range moves {
		fmt.Println(Encode(m))
	}
	// Output:
	// b2b4
	// f2f1q
	// e1g1
}

func TestDecode(t *testing.T) {
	cases := []struct {
		in      string
		want    core.Move
		wantErr string
	}{
		{in: "b2b4", want: core.Move{From: core.B2, To: core.B4}},
		{in: "f2f1q", want: core.Move{From: core.F2, To: core.F1, Promotion: core.Queen}},
		{in: "e1g1", want: core.Move{From: core.E1, To: core.G1}},

		{in: "b2b4x", wantErr: "invalid promotion: x"},
		{in: "b2b", wantErr: "invalid length: 3"},
		{in: "b2b4qq", wantErr: "invalid length: 6"},
		{in: "k1b2", wantErr: "invalid source square: k1"},
		{in: "b2k1", wantErr: "invalid target square: k1"},
	}

	for _, tc := range cases {
		got, err := Decode(tc.in)
		if err != nil {
			if err.Error() != tc.wantErr {
				t.Errorf("Decode(%q) error = %q, want %q", tc.in, err, tc.wantErr)
			}
			continue
		}
		if got != tc.want {
			t.Errorf("Decode(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}
