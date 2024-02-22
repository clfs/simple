package core

import "testing"

func TestRank_String(t *testing.T) {
	cases := []struct {
		rank Rank
		want string
	}{
		{Rank1, "Rank1"},
		{Rank2, "Rank2"},
		{Rank3, "Rank3"},
		{Rank4, "Rank4"},
		{Rank5, "Rank5"},
		{Rank6, "Rank6"},
		{Rank7, "Rank7"},
		{Rank8, "Rank8"},
		{Rank(42), "Rank(42)"},
	}
	for i, c := range cases {
		got := c.rank.String()
		if got != c.want {
			t.Errorf("%d: got %q, want %q", i, got, c.want)
		}
	}
}
