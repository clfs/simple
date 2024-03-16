// Package epd implements encoding and decoding of Extended Position Description
// (EPD) as defined in "Standard: Portable Game Notation Specification and
// Implementation Guide", revision 1994.03.12.
//
// This package makes the simplifying assumption that all EPD operations contain
// at most one operand.
package epd

// An Op is an EPD operation.
type Op struct {
	Opcode  string
	Operand string
}

// EPD opcode constants.
const (
	OpcodeFullMoveNumber = "fmvn"
	OpcodeHalfMoveClock  = "hmvc"
)

// Starting is the EPD string for the starting position.
const Starting = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -"
