package core

import "testing"

func TestColor_String(t *testing.T) {
	cases := []struct {
		color Color
		want  string
	}{
		{White, "White"},
		{Black, "Black"},
	}
	for i, c := range cases {
		got := c.color.String()
		if got != c.want {
			t.Errorf("%d: got %q, want %q", i, got, c.want)
		}
	}
}

func TestColor_Other(t *testing.T) {
	cases := []struct {
		color Color
		want  Color
	}{
		{White, Black},
		{Black, White},
	}
	for _, c := range cases {
		got := c.color.Other()
		if got != c.want {
			t.Errorf("%s.Other() = %s, want %s", c.color, got, c.want)
		}
	}
}
