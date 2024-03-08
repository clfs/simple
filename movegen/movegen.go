package movegen

import (
	"github.com/clfs/simple/core"
	"github.com/clfs/simple/movegen/internal/reference"
)

// LegalMoves returns all legal moves in a position.
func LegalMoves(p core.Position) []core.Move {
	return reference.LegalMoves(p)
}

// Perft walks a move generation tree and returns the number of leaf nodes at
// the given depth.
//
// It panics if depth is negative.
func Perft(p core.Position, depth int) int {
	if depth < 0 {
		panic("negative perft depth")
	}

	if depth == 0 {
		return 1
	}

	moves := LegalMoves(p)

	var nodes int
	for _, m := range moves {
		child := p
		child.Make(m)
		nodes += Perft(child, depth-1)
	}
	return nodes
}

// Divide returns a map from moves to the number of nodes at the decremented
// depth.
func Divide(p core.Position, depth int) map[core.Move]int {
	if depth == 0 {
		return nil
	}

	moves := LegalMoves(p)

	res := make(map[core.Move]int)
	for _, m := range moves {
		child := p
		child.Make(m)
		res[m] = Perft(child, depth-1)
	}
	return res
}
