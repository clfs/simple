// Package uci implements encoding and decoding of Universal Chess Interface
// messages.
package uci

import (
	"bytes"
	"errors"
)

// ErrUnmarshalWrongPrefix is returned when unmarshaling a UCI message that does
// not start with the expected message prefix.
var ErrUnmarshalWrongPrefix = errors.New("failed to unmarshal message: wrong message prefix")

// ErrUnmarshalInvalidArgs is returned when unmarshaling a UCI message that
// has invalid arguments.
var ErrUnmarshalInvalidArgs = errors.New("failed to unmarshal message: invalid arguments")

// UCI represents the "uci" command.
//
// It tells the engine to use UCI mode.
type UCI struct{}

func (u *UCI) UnmarshalText(text []byte) error {
	fields := bytes.Fields(text)

	if len(fields) == 0 {
		return ErrUnmarshalWrongPrefix
	}

	if !bytes.Equal(fields[0], []byte("uci")) {
		return ErrUnmarshalWrongPrefix
	}

	if len(fields) > 1 {
		return ErrUnmarshalInvalidArgs
	}

	return nil
}

func (u *UCI) MarshalText() ([]byte, error) {
	return []byte("uci"), nil
}
