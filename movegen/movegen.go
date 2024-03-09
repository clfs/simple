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

	var n int
	for _, m := range LegalMoves(p) {
		child := p
		child.Make(m)
		n += Perft(child, depth-1)
	}
	return n
}

// Divide returns a map from legal moves to their Perft counts at the
// decremented depth.
func Divide(p core.Position, depth int) map[core.Move]int {
	m := make(map[core.Move]int)
	for _, move := range LegalMoves(p) {
		child := p
		child.Make(move)
		m[move] = Perft(child, depth-1)
	}
	return m
}
