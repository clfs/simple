// Package epd implements encoding and decoding of Extended Position Description
// (EPD) as defined in "Standard: Portable Game Notation Specification and
// Implementation Guide", revision 1994.03.12.
//
// All moves are encoded in Standard Algebraic Notation (SAN).
package epd

// Op represents an EPD operation.
type Op struct {
	Opcode   string
	Operands []string
}

// EPD opcode constants.
const (
	OpcodeFullMoveNumber = "fmvn"
	OpcodeHalfMoveClock  = "hmvc"
)

// Unknown represents an unknown EPD operation.
type Unknown struct {
	Opcode string
	Input  string
}

// AnalysisCountNodes represents the "acn" operation.
type AnalysisCountNodes struct {
	Operand int
}

// AnalysisCountSeconds represents the "acs" operation.
type AnalysisCountSeconds struct {
	Operand int
}

// AvoidMoves represents the "am" operation.
type AvoidMoves struct {
	Operands []string
}

// BestMoves represents the "bm" operation.
type BestMoves struct {
	Operands []string
}

// Comment represents the "c0" through "c9" operations.
type Comment struct {
	Level   int
	Operand string
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

// FullMoveNumber represents the "fmvn" operation.
type FullMoveNumber struct {
	Operand int
}

// HalfMoveClock represents the "hmvc" operation.
type HalfMoveClock struct {
	Operand int
}

// PositionIdentification represents the "id" operation.
type PositionIdentification struct {
	Operand string
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
