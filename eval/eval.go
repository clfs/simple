// Package eval implements positional evaluation.
package eval

import (
	"github.com/clfs/simple/core"
	"github.com/clfs/simple/movegen"
)

var pieceTypeWeights = map[core.PieceType]int{
	core.Pawn:   100,
	core.Knight: 300,
	core.Bishop: 300,
	core.Rook:   500,
	core.Queen:  900,
}

// Eval returns the relative value of the position. Positive values indicate an
// advantage for the side to move, while negative values indicate an advantage
// for the opponent.
func Eval(p core.Position) int {
	friendlyValue := value(p)
	p.SideToMove = p.SideToMove.Other()
	opponentValue := value(p)
	return friendlyValue - opponentValue
}

func value(p core.Position) int {
	var res int

	// Reward piece presence.
	for s := core.A1; s <= core.H8; s++ {
		piece, ok := p.Board.Get(s)
		if ok && piece.Color() == p.SideToMove {
			res += pieceTypeWeights[piece.Type()]
		}
	}

	// Penalize weak pawn structures.
	dp := doubledPawns(p)
	bp := blockedPawns(p)
	ip := isolatedPawns(p)
	res -= 50 * (dp.Count() + bp.Count() + ip.Count())

	// Reward mobility.
	res += 10 * len(movegen.LegalMoves(p))

	return res
}

// doubledPawns returns a bitboard of all doubled pawns for the side to move.
func doubledPawns(_ core.Position) core.Bitboard {
	return 0
}

// blockedPawns returns a bitboard of all blocked pawns for the side to move.
func blockedPawns(_ core.Position) core.Bitboard {
	return 0
}

// isolatedPawns returns a bitboard of all isolated pawns for the side to move.
func isolatedPawns(_ core.Position) core.Bitboard {
	return 0
}
