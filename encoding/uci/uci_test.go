package uci

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
