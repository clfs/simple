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

		// Encode's output must always be decodable.
		p2, err := Decode(s2)
		if err != nil {
			t.Fatalf("%q -> %q -> dec failed: %v", s, s2, err)
		}

		// After the second round-trip, the FEN must be stable.
		s3 := Encode(p2)
		if s2 != s3 {
			t.Fatalf("%q -> %q -> %q", s, s2, s3)
		}
	})
}
