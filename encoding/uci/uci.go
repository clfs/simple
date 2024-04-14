// Package uci implements encoding and decoding of Universal Chess Interface
// messages.
package uci

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
	"github.com/clfs/simple/encoding/pcn"
)

// A Message contains a UCI message.
type Message interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

// ErrEmptyMessage is returned when an empty message is read.
var ErrEmptyMessage = errors.New("empty message")

// ErrUnknownMessage is returned when an unknown message is read.
var ErrUnknownMessage = errors.New("unknown message")

// ErrInvalidArgs is returned when marshalling or unmarshalling a message with
// invalid arguments.
var ErrInvalidArgs = errors.New("invalid message args")

// Parse parses a UCI message.
func Parse(b []byte) (Message, error) {
	fields := bytes.Fields(b)

	if len(fields) == 0 {
		return nil, ErrEmptyMessage
	}

	var m Message

	switch prefix := string(fields[0]); prefix {
	case "uci":
		m = new(UCI)
	default:
		return nil, fmt.Errorf("%w: %q", ErrUnknownMessage, prefix)
	}

	if err := m.UnmarshalText(b); err != nil {
		return nil, err
	}

	return m, nil
}

// ErrUnmarshalWrongPrefix is returned when unmarshaling a UCI message that does
// not start with the expected message prefix.
var ErrUnmarshalWrongPrefix = errors.New("failed to unmarshal message: wrong message prefix")

// ErrUnmarshalInvalidArgs is returned when unmarshaling a UCI message that
// has invalid arguments.
var ErrUnmarshalInvalidArgs = errors.New("failed to unmarshal message: invalid arguments")

// ErrUnmarshalEmptyMessage is returned when unmarshaling an empty UCI message.
var ErrUnmarshalEmptyMessage = errors.New("failed to unmarshal message: empty message")

// UCI represents the "uci" command.
//
// It tells the engine to use UCI mode.
type UCI struct{}

func (msg *UCI) UnmarshalText(text []byte) error {
	fields := bytes.Fields(text)

	if len(fields) == 0 {
		return ErrUnmarshalEmptyMessage
	}

	if !bytes.Equal(fields[0], []byte("uci")) {
		return ErrUnmarshalWrongPrefix
	}

	if len(fields) > 1 {
		return ErrUnmarshalInvalidArgs
	}

	return nil
}

func (msg *UCI) MarshalText() ([]byte, error) {
	return []byte("uci"), nil
}

// IsReady represents the "isready" command.
//
// It asks the engine if it's ready.
type IsReady struct{}

func (msg *IsReady) UnmarshalText(text []byte) error {
	fields := bytes.Fields(text)

	if len(fields) == 0 {
		return ErrUnmarshalEmptyMessage
	}

	if !bytes.Equal(fields[0], []byte("isready")) {
		return ErrUnmarshalWrongPrefix
	}

	if len(fields) > 1 {
		return ErrUnmarshalInvalidArgs
	}

	return nil
}

func (msg *IsReady) MarshalText() ([]byte, error) {
	return []byte("isready"), nil
}

// UCINewGame represents the "ucinewgame" command.
//
// It tells the engine that a new game is starting.
type UCINewGame struct{}

func (msg *UCINewGame) UnmarshalText(text []byte) error {
	fields := bytes.Fields(text)

	if len(fields) == 0 {
		return ErrUnmarshalEmptyMessage
	}

	if !bytes.Equal(fields[0], []byte("ucinewgame")) {
		return ErrUnmarshalWrongPrefix
	}

	if len(fields) > 1 {
		return ErrUnmarshalInvalidArgs
	}

	return nil
}

func (msg *UCINewGame) MarshalText() ([]byte, error) {
	return []byte("ucinewgame"), nil
}

// Position represents the "position" command.
//
// It tells the engine to set up a position.
type Position struct {
	Start core.Position
	Moves []core.Move
}

func (msg *Position) UnmarshalText(text []byte) error {
	fields := bytes.Fields(text)

	if len(fields) == 0 {
		return ErrUnmarshalEmptyMessage
	}

	if !bytes.Equal(fields[0], []byte("position")) {
		return ErrUnmarshalWrongPrefix
	}

	if len(fields) < 2 {
		return ErrUnmarshalInvalidArgs
	}

	usesStartpos := bytes.Equal(fields[1], []byte("startpos"))

	if usesStartpos {
		msg.Start = core.NewPosition()
	} else {
		if len(fields) < 7 {
			return ErrUnmarshalInvalidArgs
		}
		p, err := fen.Decode(string(bytes.Join(fields[1:7], []byte(" "))))
		if err != nil {
			return ErrUnmarshalInvalidArgs
		}
		msg.Start = p
	}

	if len(fields) == 2 {
		return nil
	}

	var moveFields [][]byte
	if usesStartpos {
		moveFields = fields[2:]
	} else {
		moveFields = fields[7:]
	}

	var moves []core.Move
	for _, f := range moveFields {
		m, err := pcn.Decode(string(f))
		if err != nil {
			return ErrUnmarshalInvalidArgs
		}

		moves = append(moves, m)
	}

	msg.Moves = moves

	return nil
}

func (msg *Position) MarshalText() ([]byte, error) {
	var b bytes.Buffer

	b.WriteString("position ")

	b.WriteString(fen.Encode(msg.Start))

	for _, m := range msg.Moves {
		b.WriteByte(' ')
		b.WriteString(pcn.Encode(m))
	}

	return b.Bytes(), nil
}
