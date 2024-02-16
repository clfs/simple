package fen

import "testing"

func FuzzRoundTrip(f *testing.F) {
	f.Add(Starting)
	for _, fen := range readFENs(f, "testdata/valid.fen") {
		f.Add(fen)
	}

	f.Fuzz(func(t *testing.T, s string) {
		p, err := Decode(s)
		if err != nil {
			t.Skip() // invalid FEN
		}

		s2 := Encode(p)
		if s != s2 {
			t.Fatalf("%q -> %q", s, s2)
		}
	})
}
