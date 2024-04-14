// Package uci implements encoding and decoding of Universal Chess Interface
// messages.
package uci

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/fen"
	"github.com/clfs/simple/encoding/pcn"
)

// A Message contains a UCI message.
type Message interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

var (
	// ErrEmptyMessage is returned when unmarshaling an empty message.
	ErrEmptyMessage = errors.New("empty message")

	// ErrUnknownMessage is returned when unmarshaling an unknown message.
	ErrUnknownMessage = errors.New("unknown message")

	// ErrWrongMessageType is returned when unmarshaling a message into the
	// wrong type.
	ErrWrongMessageType = errors.New("wrong message type")

	// ErrInvalidArgs is returned when marshaling or unmarshaling a message with
	// invalid arguments.
	ErrInvalidArgs = errors.New("invalid arguments")
)

// Parse parses a UCI message.
func Parse(b []byte) (Message, error) {
	fields := bytes.Fields(b)

	if len(fields) == 0 {
		return nil, ErrEmptyMessage
	}

	var m Message

	prefix := string(fields[0])

	switch prefix {
	case "uci":
		m = new(UCI)
	case "isready":
		m = new(IsReady)
	case "ucinewgame":
		m = new(UCINewGame)
	case "position":
		m = new(Position)
	case "go":
		m = new(Go)
	default:
		return nil, fmt.Errorf("%q: %w", prefix, ErrUnknownMessage)
	}

	if err := m.UnmarshalText(b); err != nil {
		return nil, fmt.Errorf("%q: %w", prefix, err)
	}

	return m, nil
}

// UCI represents the "uci" command.
//
// It tells the engine to use UCI mode.
type UCI struct{}

func (msg *UCI) UnmarshalText(text []byte) error {
	fields := bytes.Fields(text)

	if len(fields) == 0 {
		return ErrEmptyMessage
	}

	if !bytes.Equal(fields[0], []byte("uci")) {
		return ErrWrongMessageType
	}

	if len(fields) > 1 {
		return ErrInvalidArgs
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
		return ErrEmptyMessage
	}

	if !bytes.Equal(fields[0], []byte("isready")) {
		return ErrWrongMessageType
	}

	if len(fields) > 1 {
		return ErrInvalidArgs
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
		return ErrEmptyMessage
	}

	if !bytes.Equal(fields[0], []byte("ucinewgame")) {
		return ErrWrongMessageType
	}

	if len(fields) > 1 {
		return ErrInvalidArgs
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
		return ErrEmptyMessage
	}

	if !bytes.Equal(fields[0], []byte("position")) {
		return ErrWrongMessageType
	}

	if len(fields) < 2 {
		return fmt.Errorf("no position provided: %w", ErrInvalidArgs)
	}

	usesStartpos := bytes.Equal(fields[1], []byte("startpos"))

	if usesStartpos {
		msg.Start = core.NewPosition()
	} else {
		if len(fields) < 7 {
			return fmt.Errorf("too few arguments: %w", ErrInvalidArgs)
		}
		p, err := fen.Decode(string(bytes.Join(fields[1:7], []byte(" "))))
		if err != nil {
			return fmt.Errorf("invalid FEN: %w: %w", err, ErrInvalidArgs)
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
			return fmt.Errorf("invalid move: %w: %w", err, ErrInvalidArgs)
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

// Go represents the "go" command.
//
// It tells the engine to search the current position.
type Go struct {
	SearchMoves  []core.Move
	Ponder       bool
	WTime, BTime time.Duration
	WInc, BInc   time.Duration
	MovesToGo    int
	Depth        int
	Nodes        int
	Mate         int
	MoveTime     time.Duration
	Infinite     bool
}

func (msg *Go) UnmarshalText(text []byte) error {
	fields := strings.Fields(string(text))

	if len(fields) == 0 {
		return ErrEmptyMessage
	}

	if fields[0] != "go" {
		return ErrWrongMessageType
	}

	return errors.New("not implemented")
}

func (msg *Go) MarshalText() ([]byte, error) {
	return nil, errors.New("not implemented")
}
