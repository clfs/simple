//go:generate go run tablegen.go -output tables.go

// Package movegen generates moves.
package movegen

import "github.com/clfs/simple/core"

// Moves returns available moves in the position.
//
// With respect to the FIDE Laws of Chess, edition January 2023:
//
//   - Moves does not account for threefold reptition (9.2).
//   - Moves does not account for fivefold reptition (9.6.1).
//   - Moves returns no moves if the seventy-five move rule applies (9.6.2).
//   - Moves may return one or more moves even if the fifty-move rule applies (9.3).
func Moves(p core.Position) []core.Move {
	// Seventy-five move rule.
	if p.HalfMoveClock >= 75 {
		return nil
	}

	return nil
}
