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

func TestUCI_MarshalText(t *testing.T) {
	cases := []struct {
		in      UCI
		want    string
		wantErr error
	}{
		{in: UCI{}, want: "uci"},
	}

	for i, tc := range cases {
		got, err := tc.in.MarshalText()
		if err != tc.wantErr {
			t.Errorf("#%d: wrong error: want %v, got %v", i, tc.wantErr, err)
		}

		if tc.wantErr != nil {
			continue
		}

		if tc.want != string(got) {
			t.Errorf("#%d: want %q, got %q", i, tc.want, got)
		}
	}
}

func TestIsReady_MarshalText(t *testing.T) {
	cases := []struct {
		in      IsReady
		want    string
		wantErr error
	}{
		{in: IsReady{}, want: "isready"},
	}

	for i, tc := range cases {
		got, err := tc.in.MarshalText()
		if err != tc.wantErr {
			t.Errorf("#%d: wrong error: want %v, got %v", i, tc.wantErr, err)
		}

		if tc.wantErr != nil {
			continue
		}

		if tc.want != string(got) {
			t.Errorf("#%d: want %q, got %q", i, tc.want, got)
		}
	}
}
