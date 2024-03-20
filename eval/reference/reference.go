// Package reference implements a reference evaluation function that only
// considers material value.
package reference

import "github.com/clfs/simple/core"

var pieceTypeWeights = map[core.PieceType]int{
	core.Pawn:   100,
	core.Knight: 300,
	core.Bishop: 300,
	core.Rook:   500,
	core.Queen:  900,
}

// Eval returns the total material advantage for the side to move.
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
