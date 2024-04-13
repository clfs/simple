package uci

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var unmarshalTests = []struct {
	in  string
	ptr any // new(type)
	out any
	err error
}{
	{in: "uci", ptr: new(UCI), out: UCI{}},
	{in: "foo", ptr: new(UCI), err: ErrUnmarshalWrongPrefix},
	{in: "uci foo", ptr: new(UCI), err: ErrUnmarshalInvalidArgs},
	{in: " ", ptr: new(UCI), err: ErrUnmarshalEmptyMessage},

	{in: "isready", ptr: new(IsReady), out: IsReady{}},
	{in: "foo", ptr: new(IsReady), err: ErrUnmarshalWrongPrefix},
	{in: "isready foo", ptr: new(IsReady), err: ErrUnmarshalInvalidArgs},
	{in: " ", ptr: new(IsReady), err: ErrUnmarshalEmptyMessage},
}

func TestUnmarshalText(t *testing.T) {
	for i, tc := range unmarshalTests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			in := []byte(tc.in)

			typ := reflect.TypeOf(tc.ptr).Elem()

			v := reflect.New(typ)

			err := v.Interface().(Message).UnmarshalText(in)
			if err != tc.err {
				t.Errorf("wrong error: want %v, got %v", tc.err, err)
			}

			if tc.err != nil {
				return
			}

			if diff := cmp.Diff(tc.out, v.Elem().Interface()); diff != "" {
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
