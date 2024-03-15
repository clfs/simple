package epd

import "testing"

func FuzzRoundTrip(f *testing.F) {
	f.Add(`rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - foo bar; c0 "my comment"; bm c4 Nf3;`)
	f.Fuzz(func(t *testing.T, s string) {
		p, ops, err := Decode(s)
		if err != nil {
			return
		}

		s2 := Encode(p, ops)
		if s != s2 {
			t.Errorf("changed after round trip: %q -> %q", s, s2)
		}
	})
}
