// Package epd implements encoding and decoding of Extended Position Description
// (EPD) as defined in "Standard: Portable Game Notation Specification and
// Implementation Guide", revision 1994.03.12.
//
// All moves are encoded in Standard Algebraic Notation (SAN).
package epd

import (
	"fmt"
	"strconv"
	"strings"
)

// Operation represents an EPD operation.
type Operation interface {
	Assemble() RawOperation
}

// Disassemble attempts to parse raw back into Operations.
// Unrecognized RawOperations are passed through unchanged to the output.
// The allDecoded value is true if ops contains no RawOperations.
func Disassemble(raw []RawOperation) (ops []Operation, allDecoded bool) {
	return
}

// RawOperation represents a raw EPD operation.
type RawOperation struct {
	Opcode string
	Args   string
}

func (op RawOperation) Assemble() RawOperation {
	return op
}

func (op RawOperation) Disassemble() Operation {
	switch op.Opcode {
	default:
		return op
	case "acn":
		n, err := strconv.Atoi(op.Args)
		if err != nil {
			return op
		}
		return ACN{Nodes: n}
	}
}

// ACN represents the number of nodes examined in an analysis.
type ACN struct {
	Nodes int
}

func (op ACN) Assemble() RawOperation {
	return RawOperation{
		Opcode: "acn",
		Args:   fmt.Sprintf("%d", op.Nodes),
	}
}

// ACS represents the number of seconds used for an analysis.
type ACS struct {
	Seconds int
}

func (op ACS) Assemble() RawOperation {
	return RawOperation{
		Opcode: "acs",
		Args:   fmt.Sprintf("%d", op.Seconds),
	}
}

// AvoidMoves represents the "am" operation.
type AvoidMoves struct {
	Operands []string
}

// BM represents the best available moves.
type BM struct {
	Moves []string
}

func (op BM) Assemble() RawOperation {
	return RawOperation{
		Opcode: "bm",
		Args:   strings.Join(op.Moves, " "),
	}
}

// Comment represents the "c0" through "c9" operations.
type Comment struct {
	Level   int
	Comment string
}

func (c Comment) Assemble() RawOperation {
	// TODO: Consider how to handle levels outside of 0-9.
	return RawOperation{
		Opcode: fmt.Sprintf("c%d", c.Level),
		Args:   c.Comment,
	}
}

// CentipawnEvaluation represents the "ce" operation.
type CentipawnEvaluation struct {
	Operand int
}

// DirectMateFullMoveCount represents the "dm" operation.
type DirectMateFullMoveCount struct {
	Operand int
}

// DrawAccept represents the "draw_accept" operation.
type DrawAccept struct{}

// DrawClaim represents the "draw_claim" operation.
type DrawClaim struct{}

// DrawOffer represents the "draw_offer" operation.
type DrawOffer struct{}

// DrawReject represents the "draw_reject" operation.
type DrawReject struct{}

// ECO represents the "eco" operation.
type ECO struct {
	Operand string
}

// FMVN represents the full move number.
type FMVN struct {
	Number int
}

func (op FMVN) Assemble() RawOperation {
	return RawOperation{
		Opcode: "fmvn",
		Args:   fmt.Sprintf("%d", op.Number),
	}
}

// HMVC represents the half move clock.
type HMVC struct {
	Clock int
}

func (op HMVC) Assemble() RawOperation {
	return RawOperation{
		Opcode: "hmvc",
		Args:   fmt.Sprintf("%d", op.Clock),
	}
}

// ID represents a position identifier.
type ID struct {
	ID string
}

func (op ID) Assemble() RawOperation {
	return RawOperation{
		Opcode: "id",
		Args:   op.ID,
	}
}

// NIC represents the "nic" operation.
type NIC struct {
	Operand string
}

// NOOP represents the "noop" operation.
type NOOP struct{}

// PredictedMove represents the "pm" operation.
type PredictedMove struct {
	Operand string
}

// PredictedVariation represents the "pv" operation.
type PredictedVariation struct {
	Operands []string
}

// RepetitionCount represents the "rc" operation.
type RepetitionCount struct {
	Operand int
}

// GameResignation represents the "resign" operation.
type GameResignation struct{}

// SuppliedMove represents the "sm" operation.
type SuppliedMove struct {
	Operand string
}

// TCGameSelector represents the "tcgs" operation.
type TCGameSelector struct {
	Operand string
}

// TCReceiverIdentification represents the "tcri" operation.
type TCReceiverIdentification struct {
	Operand string
}

// TCSenderIdentification represents the "tcsi" operation.
type TCSenderIdentification struct {
	Operand string
}

// Variation represents the "v0" through "v9" operations.
type Variation struct {
	Level   int
	Operand string
}
