package epd

import "testing"

func TestUnknown_EncodeOp(t *testing.T) {
	cases := []struct {
		in   Unknown
		want string
	}{
		{Unknown{"foo", "bar"}, "foo bar;"},
	}

	for i, c := range cases {
		got, err := c.in.EncodeOp()
		if err != nil {
			t.Errorf("#%d: error: %v", i, err)
		}
		if got != c.want {
			t.Errorf("#%d: want %q, got %q", i, c.want, got)
		}
	}
}
