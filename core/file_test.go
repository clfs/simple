package core

import "testing"

func TestFile_String(t *testing.T) {
	cases := []struct {
		file File
		want string
	}{
		{FileA, "FileA"},
		{FileB, "FileB"},
		{FileC, "FileC"},
		{FileD, "FileD"},
		{FileE, "FileE"},
		{FileF, "FileF"},
		{FileG, "FileG"},
		{FileH, "FileH"},
		{File(42), "File(42)"},
	}
	for i, c := range cases {
		got := c.file.String()
		if got != c.want {
			t.Errorf("%d: got %q, want %q", i, got, c.want)
		}
	}
}
