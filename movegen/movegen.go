//go:generate go run tablegen.go -output tables.go

// Package movegen generates moves.
package movegen

import "github.com/clfs/simple/core"

// Moves returns legal moves in the position.
//
// It does not account for threefold repetition, fivefold repetition, the
// fifty-move rule, or the seventy-five move rule.
func Moves(p core.Position) []core.Move {
	return nil
}

// IsLegal returns true if the move is legal in the position.
//
// It does not account for threefold repetition, fivefold repetition, the
// fifty-move rule, or the seventy-five move rule.
func IsLegal(p core.Position, m core.Move) bool {
	return false
}
