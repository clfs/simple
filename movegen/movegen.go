package movegen

import (
	"github.com/clfs/simple/core"
	"github.com/clfs/simple/movegen/internal/reference"
)

// MovesFunc is the signature of a function that generates moves.
//
// Implementations are unable to account for threefold and fivefold repetition.
type MovesFunc func(p core.Position) []core.Move

// LegalMoves returns all legal moves in a position.
func LegalMoves(p core.Position) []core.Move {
	return reference.LegalMoves(p)
}
