package epd

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func FuzzRoundTrip(f *testing.F) {
	for _, tc := range decodeTestCases {
		f.Add(tc.in)
	}
	f.Fuzz(func(t *testing.T, s string) {
		p, ops, err := Decode(s)
		if err != nil {
			return
		}

		p2, ops2, err := Decode(Encode(p, ops))
		if err != nil {
			t.Errorf("round trip failed: %v", err)
		}
		if diff := cmp.Diff(p, p2); diff != "" {
			t.Errorf("position changed (-old +new): %v", diff)
		}
		if diff := cmp.Diff(ops, ops2); diff != "" {
			t.Errorf("operations changed (-old +new): %v", diff)
		}
	})
}
