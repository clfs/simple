// Package epd implements encoding and decoding of Extended Position Description
// (EPD) as defined in "Standard: Portable Game Notation Specification and
// Implementation Guide", revision 1994.03.12.
//
// For simplicity, this package assumes all semicolons in an EPD string are
// operation terminators.
package epd

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// An Op3 is an EPD operation.
type Op3 struct {
	Opcode   string
	Operands string
}

var (
	ErrNoOpcode        = errors.New("operation has no opcode")
	ErrInvalidOpcode   = errors.New("operation has invalid opcode")
	ErrInvalidOperands = errors.New("operation has invalid operands")
)

// MarshalText implements encoding.TextMarshaler for Op.
func (op *Op3) MarshalText() ([]byte, error) {
	if op.Opcode == "" {
		return nil, ErrNoOpcode
	}
	if strings.Contains(op.Opcode, ";") {
		return nil, ErrInvalidOpcode
	}
	if strings.Contains(op.Operands, ";") {
		return nil, ErrInvalidOperands
	}

	if op.Operands == "" {
		return fmt.Appendf(nil, "%s;", op.Opcode), nil
	}
	return fmt.Appendf(nil, "%s %s;", op.Opcode, op.Operands), nil
}

// UnmarshalText implements encoding.TextUnmarshaler for Op.
func (op *Op3) UnmarshalText(text []byte) error {
	b := bytes.TrimSpace(text)

	opcode, operands, _ := bytes.Cut(b, []byte(" "))

	if len(opcode) == 0 {
		return ErrNoOpcode
	}
	if bytes.ContainsRune(opcode, ';') {
		return ErrInvalidOpcode
	}
	if bytes.ContainsRune(operands, ';') {
		return ErrInvalidOperands
	}

	op.Opcode = string(opcode)
	op.Operands = string(operands)

	return nil
}

// EPD opcode constants.
const (
	OpcodeFullMoveNumber = "fmvn"
	OpcodeHalfMoveClock  = "hmvc"
)
