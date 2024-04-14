package uci

import (
	"errors"
	"fmt"
	"reflect"
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

var unmarshalTests = []struct {
	in  string
	typ reflect.Type
	out Message
	err error
}{
	{in: "uci", typ: reflect.TypeOf(UCI{}), out: &UCI{}},
	{in: "foo", typ: reflect.TypeOf(UCI{}), err: ErrUnmarshalWrongPrefix},
	{in: "uci foo", typ: reflect.TypeOf(UCI{}), err: ErrUnmarshalInvalidArgs},
	{in: " ", typ: reflect.TypeOf(UCI{}), err: ErrUnmarshalEmptyMessage},

	{in: "isready", typ: reflect.TypeOf(IsReady{}), out: &IsReady{}},
	{in: "foo", typ: reflect.TypeOf(IsReady{}), err: ErrUnmarshalWrongPrefix},
	{in: "isready foo", typ: reflect.TypeOf(IsReady{}), err: ErrUnmarshalInvalidArgs},
	{in: " ", typ: reflect.TypeOf(IsReady{}), err: ErrUnmarshalEmptyMessage},

	{in: "ucinewgame", typ: reflect.TypeOf(UCINewGame{}), out: &UCINewGame{}},
	{in: "foo", typ: reflect.TypeOf(UCINewGame{}), err: ErrUnmarshalWrongPrefix},
	{in: "ucinewgame foo", typ: reflect.TypeOf(UCINewGame{}), err: ErrUnmarshalInvalidArgs},
	{in: " ", typ: reflect.TypeOf(UCINewGame{}), err: ErrUnmarshalEmptyMessage},

	{
		in:  "position startpos",
		typ: reflect.TypeOf(Position{}),
		out: &Position{
			Start: core.NewPosition(),
		},
	},
	{
		in:  "position foo",
		typ: reflect.TypeOf(Position{}),
		err: ErrUnmarshalInvalidArgs,
	},
	{
		in:  " ",
		typ: reflect.TypeOf(Position{}),
		err: ErrUnmarshalEmptyMessage,
	},
	{
		in:  "foo",
		typ: reflect.TypeOf(Position{}),
		err: ErrUnmarshalWrongPrefix,
	},
	{
		in:  "position",
		typ: reflect.TypeOf(Position{}),
		err: ErrUnmarshalInvalidArgs,
	},
	{
		in:  "position 3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1",
		typ: reflect.TypeOf(Position{}),
		out: &Position{
			Start: fen.MustDecode("3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1"),
		},
	},
	{
		in:  "position 3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1 h5c5 a5b6",
		typ: reflect.TypeOf(Position{}),
		out: &Position{
			Start: fen.MustDecode("3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1"),
			Moves: []core.Move{
				pcn.MustDecode("h5c5"),
				pcn.MustDecode("a5b6"),
			},
		},
	},
	{
		in:  "position startpos e2e4 e7e5",
		typ: reflect.TypeOf(Position{}),
		out: &Position{
			Start: core.NewPosition(),
			Moves: []core.Move{
				pcn.MustDecode("e2e4"),
				pcn.MustDecode("e7e5"),
			},
		},
	},
}

func TestUnmarshalText(t *testing.T) {
	for i, tc := range unmarshalTests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			in := []byte(tc.in)

			// Create a Message that contains the specified type.
			msg := reflect.New(tc.typ).Interface().(Message)

			err := msg.UnmarshalText(in)
			if err != tc.err {
				t.Errorf("wrong error: want %v, got %v", tc.err, err)
			}

			if tc.err != nil {
				return
			}

			if diff := cmp.Diff(tc.out, msg); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
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
