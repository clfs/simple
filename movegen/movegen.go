package movegen

import (
	"github.com/clfs/simple/core"
	"github.com/clfs/simple/movegen/internal/reference"
)

// LegalMoves returns all legal moves in a position.
func LegalMoves(p core.Position) []core.Move {
	return reference.LegalMoves(p)
}

// Perft returns the number of leaf nodes at the selected depth in a position's
// move tree.
//
// Perft returns zero if depth is negative.
func Perft(p core.Position, depth int) int {
	if depth < 0 {
		return 0
	}

	if depth == 0 {
		return 1
	}

	var n int
	for _, m := range LegalMoves(p) {
		child := p
		child.Make(m)
		n += Perft(child, depth-1)
	}
	return n
}

// Divide returns a map from legal moves to Perft node counts at a decremented
// depth.
//
// Divide returns nil if depth is not positive.
func Divide(p core.Position, depth int) map[core.Move]int {
	if depth < 1 {
		return nil
	}

	res := make(map[core.Move]int)
	for _, move := range LegalMoves(p) {
		child := p
		child.Make(move)
		res[move] = Perft(child, depth-1)
	}
	return res
}
