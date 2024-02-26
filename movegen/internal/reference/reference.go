// Package reference is a reference move generator implementation.
package reference

import "github.com/clfs/simple/core"

// LegalMoves returns all legal moves in a position.
func LegalMoves(p core.Position) []core.Move {
	if p.HalfMoveClock >= 75 {
		return nil
	}

	var moves []core.Move
	moves = append(moves, pawnPushes(p)...)
	moves = append(moves, knightMoves(p)...)

	return moves
}

func pawnPushes(p core.Position) []core.Move {
	var moves []core.Move

	var fromBB core.Bitboard

	if p.SideToMove == core.White {
		fromBB = p.Board[core.WhitePawn]
	} else {
		fromBB = p.Board[core.BlackPawn]
	}

	for from := core.A2; from <= core.H7; from++ {
		if !fromBB.Get(from) {
			continue // empty square
		}

		var to core.Square

		// single push
		if p.SideToMove == core.White {
			to = from.Above()
		} else {
			to = from.Below()
		}

		if p.Board.IsEmpty(to) {
			moves = append(moves, core.Move{From: from, To: to})
		} else {
			continue // double push not possible
		}

		// double push
		if p.SideToMove == core.White && from.Rank() == core.Rank2 {
			to = from.Above().Above()
		} else if p.SideToMove == core.Black && from.Rank() == core.Rank7 {
			to = from.Below().Below()
		}

		if p.Board.IsEmpty(to) {
			moves = append(moves, core.Move{From: from, To: to})
		}
	}

	return moves
}

func knightMoves(p core.Position) []core.Move {
	var moves []core.Move

	var fromBB core.Bitboard

	if p.SideToMove == core.White {
		fromBB = p.Board[core.WhiteKnight]
	} else {
		fromBB = p.Board[core.BlackKnight]
	}

	for from := core.A1; from <= core.H8; from++ {
		if !fromBB.Get(from) {
			continue // empty square
		}

		var tos []core.Square

		f, r := from.File(), from.Rank()

		if f >= core.FileB {
			if r >= core.Rank3 {
				tos = append(tos, from.Below().Below().Left())
			}
			if r <= core.Rank6 {
				tos = append(tos, from.Above().Above().Left())
			}
		}
		if f >= core.FileC {
			if r >= core.Rank2 {
				tos = append(tos, from.Below().Left().Left())
			}
			if r <= core.Rank7 {
				tos = append(tos, from.Above().Left().Left())
			}
		}
		if f <= core.FileF {
			if r >= core.Rank2 {
				tos = append(tos, from.Below().Right().Right())
			}
			if r <= core.Rank7 {
				tos = append(tos, from.Above().Right().Right())
			}
		}
		if f <= core.FileG {
			if r >= core.Rank3 {
				tos = append(tos, from.Below().Below().Right())
			}
			if r <= core.Rank6 {
				tos = append(tos, from.Above().Above().Right())
			}
		}

		for _, to := range tos {
			capturedPiece, occupied := p.Board.Get(to)
			if !occupied || capturedPiece.Color() != p.SideToMove {
				moves = append(moves, core.Move{From: from, To: to})
			}
		}
	}

	return moves
}
