package uci

import (
	"errors"
	"fmt"
	"testing"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
	"github.com/clfs/simple/encoding/pcn"
	"github.com/google/go-cmp/cmp"
)

var parseTests = []struct {
	in   string
	want Message
	err  error
}{
	{in: "", err: ErrEmptyMessage},
	{in: " ", err: ErrEmptyMessage},
	{in: "foo", err: ErrUnknownMessage},
	{in: "uci", want: &UCI{}},
	{in: "uci foo", err: ErrInvalidArgs},
	{in: "isready", want: &IsReady{}},
	{in: "isready foo", err: ErrInvalidArgs},
	{in: "ucinewgame", want: &UCINewGame{}},
	{in: "ucinewgame foo", err: ErrInvalidArgs},
	{in: "position", err: ErrInvalidArgs},
	{in: "position foo", err: ErrInvalidArgs},
	{
		in: "position startpos",
		want: &Position{
			Start: core.NewPosition(),
		},
	},
	{
		in: "position 3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1",
		want: &Position{
			Start: fen.MustDecode("3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1"),
		},
	},
	{
		in: "position 3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1 h5c5 a5b6",
		want: &Position{
			Start: fen.MustDecode("3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1"),
			Moves: []core.Move{
				pcn.MustDecode("h5c5"),
				pcn.MustDecode("a5b6"),
			},
		},
	},
	{
		in: "position startpos e2e4 e7e5",
		want: &Position{
			Start: core.NewPosition(),
			Moves: []core.Move{
				pcn.MustDecode("e2e4"),
				pcn.MustDecode("e7e5"),
			},
		},
	},
}

func TestParse(t *testing.T) {
	for i, tt := range parseTests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := Parse([]byte(tt.in))
			if !errors.Is(err, tt.err) {
				t.Errorf("wrong error: want %v, got %v", tt.err, err)
			}

			if tt.err != nil {
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want, got)")
			}
		})
	}
}

var marshalTests = []struct {
	in  Message
	out string
	err error
}{
	{in: &UCI{}, out: "uci"},
	{in: &IsReady{}, out: "isready"},
	{in: &UCINewGame{}, out: "ucinewgame"},
}

func TestMarshalText(t *testing.T) {
	for i, tc := range marshalTests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := tc.in.MarshalText()
			if err != tc.err {
				t.Errorf("wrong error: want %v, got %v", tc.err, err)
			}

			if tc.err != nil {
				return
			}

			if tc.out != string(got) {
				t.Errorf("want %q, got %q", tc.out, got)
			}
		})
	}
}
