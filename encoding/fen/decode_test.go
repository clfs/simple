package fen

import (
	"bufio"
	"os"
	"testing"

	"github.com/clfs/simple/core"
	"github.com/google/go-cmp/cmp"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		moves []core.Move
		in    string
	}{
		{
			in: Starting,
		},
		{
			moves: []core.Move{
				{From: core.E2, To: core.E4},
			},
			in: "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
		},
		{
			moves: []core.Move{
				{From: core.E2, To: core.E4},
				{From: core.E7, To: core.E5},
			},
			in: "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2",
		},
		{
			moves: []core.Move{
				{From: core.E2, To: core.E4},
				{From: core.E7, To: core.E5},
				{From: core.E1, To: core.E2},
			},
			in: "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPPKPPP/RNBQ1BNR b kq - 1 2",
		},
		{
			moves: []core.Move{
				{From: core.E2, To: core.E4},
				{From: core.E7, To: core.E5},
				{From: core.E1, To: core.E2},
				{From: core.E8, To: core.E7},
			},
			in: "rnbq1bnr/ppppkppp/8/4p3/4P3/8/PPPPKPPP/RNBQ1BNR w - - 2 3",
		},
	}

	for i, tc := range cases {
		want := core.NewPosition()
		for _, m := range tc.moves {
			want.Make(m)
		}

		got, err := Decode(tc.in)
		if err != nil {
			t.Errorf("#%d: Decode() error: %v", i, err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("#%d: Decode() mismatch (-want +got):\n%s", i, diff)
		}
	}
}

func readFENs(tb testing.TB, name string) []string {
	tb.Helper()

	f, err := os.Open(name)
	if err != nil {
		tb.Fatal(err)
	}
	defer f.Close()

	var lines []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); err != nil {
		tb.Fatal(err)
	}

	return lines
}

func TestDecode_Valid(t *testing.T) {
	for _, fen := range readFENs(t, "testdata/valid.fen") {
		if _, err := Decode(fen); err != nil {
			t.Errorf("Decode(%q) error: %v", fen, err)
		}
	}
}
