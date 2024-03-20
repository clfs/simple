// Package eval implements positional evaluation.
package eval

import "github.com/clfs/simple/core"

// EvalFunc is the signature of an evaluation function used to determine the
// relative value of a position.
//
// Positive values indicate an advantage for the side to move, and negative
// values indicate an advantage for the opponent.
type EvalFunc func(core.Position) int
