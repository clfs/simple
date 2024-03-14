package epd

import (
	"fmt"
	"strconv"
	"strings"
)

// Operation represents an EPD operation.
type Operation interface {
	Assemble() (RawOperation, error)
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

func (op RawOperation) Assemble() (RawOperation, error) {
	if op.Opcode == "" {
		return RawOperation{}, fmt.Errorf("empty opcode")
	}
	return op, nil
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

func (op ACN) Assemble() (RawOperation, error) {
	if op.Nodes < 0 {
		return RawOperation{}, fmt.Errorf("invalid node count: %d", op.Nodes)
	}
	return RawOperation{"acn", fmt.Sprintf("%d", op.Nodes)}, nil
}

// ACS represents the number of seconds used for an analysis.
type ACS struct {
	Seconds int
}

func (op ACS) Assemble() (RawOperation, error) {
	if op.Seconds < 0 {
		return RawOperation{}, fmt.Errorf("invalid number of seconds: %d", op.Seconds)
	}
	return RawOperation{"acs", fmt.Sprintf("%d", op.Seconds)}, nil
}

// AvoidMoves represents the "am" operation.
type AvoidMoves struct {
	Operands []string
}

// BM represents the best available moves.
type BM struct {
	Moves []string
}

func (op BM) Assemble() (RawOperation, error) {
	return RawOperation{"bm", strings.Join(op.Moves, " ")}, nil
}

// Comment represents the "c0" through "c9" operations.
type Comment struct {
	Level   int
	Comment string
}

func (c Comment) Assemble() (RawOperation, error) {
	if c.Level < 0 || c.Level > 9 {
		return RawOperation{}, fmt.Errorf("invalid comment level: %d", c.Level)
	}
	return RawOperation{fmt.Sprintf("c%d", c.Level), c.Comment}, nil
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

func (op FMVN) Assemble() (RawOperation, error) {
	if op.Number < 1 {
		return RawOperation{}, fmt.Errorf("invalid full move number: %d", op.Number)
	}
	return RawOperation{"fmvn", fmt.Sprintf("%d", op.Number)}, nil
}

// HMVC represents the half move clock.
type HMVC struct {
	Clock int
}

func (op HMVC) Assemble() (RawOperation, error) {
	if op.Clock < 0 {
		return RawOperation{}, fmt.Errorf("invalid half move clock: %d", op.Clock)
	}
	return RawOperation{"hmvc", fmt.Sprintf("%d", op.Clock)}, nil
}

// ID represents a position identifier.
type ID struct {
	ID string
}

func (op ID) Assemble() (RawOperation, error) {
	return RawOperation{"id", op.ID}, nil
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
