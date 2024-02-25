package movegen

import (
	"github.com/clfs/simple/core"
	"github.com/clfs/simple/movegen/internal/reference"
)

// LegalMoves returns all legal moves in a position.
func LegalMoves(p core.Position) []core.Move {
	return reference.LegalMoves(p)
}
