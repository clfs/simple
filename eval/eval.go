// Package eval describes the signature of evaluation functions.
package eval

import (
	"github.com/clfs/simple/core"
)

var pieceTypeWeights = map[core.PieceType]int{
	core.Pawn:   100,
	core.Knight: 300,
	core.Bishop: 300,
	core.Rook:   500,
	core.Queen:  900,
}

// Eval returns the relative value of a position.
//
// Positive values indicate an advantage for the side to move, and negative
// values indicate an advantage for the opponent.
func Eval(p core.Position) int {
	var res int
	for s := core.A1; s <= core.H8; s++ {
		piece, ok := p.Board.Get(s)
		if !ok {
			continue
		}
		if piece.Color() == p.SideToMove {
			res += pieceTypeWeights[piece.Type()]
		} else {
			res -= pieceTypeWeights[piece.Type()]
		}
	}
	return res
}
