// Package reference is a reference move generator implementation.
package reference

import (
	"slices"

	"github.com/clfs/simple/core"
)

// LegalMoves returns all legal moves in a position.
func LegalMoves(p core.Position) []core.Move {
	if p.HalfMoveClock >= 75 {
		return nil
	}

	moves := slices.Concat(
		pawnPushes(p),
		pawnAttacks(p),
		knightAttacks(p),
		bishopAttacks(p),
		rookAttacks(p),
		queenAttacks(p),
		kingAttacks(p),
	)

	moves = slices.DeleteFunc(moves, func(m core.Move) bool {
		p2 := p
		p2.Make(m)
		return isEnemyKingTargeted(p2)
	})

	return moves
}

// A translation describes movement in a direction.
type translation struct {
	df, dr int
}

// Available translations by piece type.
var (
	knightTranslations = []translation{{-2, 1}, {-2, -1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	bishopTranslations = []translation{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	rookTranslations   = []translation{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	queenTranslations  = []translation{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	kingTranslations   = []translation{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
)

// translate tries to apply a translation to a square.
// If successful, it returns the new square.
func translate(s core.Square, t translation) (core.Square, bool) {
	f := s.File() + core.File(t.df)
	r := s.Rank() + core.Rank(t.dr)

	if !f.Valid() || !r.Valid() {
		return 0, false
	}

	return core.NewSquare(f, r), true
}

// pawnPushes returns available pawn pushes, without considering checks.
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

// pawnAttacks returns available pawn attacks, without considering checks.
func pawnAttacks(p core.Position) []core.Move {
	return nil // TODO
}

// knightAttacks returns available knight attacks, without considering checks.
func knightAttacks(p core.Position) []core.Move {
	return steppingAttacks(p, core.Knight)
}

// steppingAttacks returns available stepping attacks, which include knight
// moves and non-castling king moves.
//
// It does not consider checks.
func steppingAttacks(p core.Position, pt core.PieceType) []core.Move {
	var (
		moves        []core.Move
		translations []translation
	)

	switch pt {
	case core.Knight:
		translations = knightTranslations
	case core.King:
		translations = kingTranslations
	default:
		panic("invalid piece type")
	}

	piece := core.NewPiece(p.SideToMove, pt)
	fromBB := p.Board[piece]

	for from := core.A1; from <= core.H8; from++ {
		// Skip non-starting squares.
		if !fromBB.Get(from) {
			continue
		}

		for _, t := range translations {
			// Apply the translation. If it fails, try the next one.
			to, ok := translate(from, t)
			if !ok {
				continue
			}

			// If the destination's empty or occupied by the enemy, generate
			// the move.
			piece, ok := p.Board.Get(to)
			if !ok || piece.Color() != p.SideToMove {
				moves = append(moves, core.Move{
					From: from,
					To:   to,
				})
			}
		}
	}

	return moves
}

// slidingAttacks returns available sliding attacks, which include bishop, rook,
// and queen moves.
//
// It does not consider checks.
func slidingAttacks(p core.Position, pt core.PieceType) []core.Move {
	var moves []core.Move

	var translations []translation

	switch pt {
	case core.Bishop:
		translations = bishopTranslations
	case core.Rook:
		translations = rookTranslations
	case core.Queen:
		translations = queenTranslations
	default:
		panic("invalid piece type")
	}

	piece := core.NewPiece(p.SideToMove, pt)
	fromBB := p.Board[piece]

	for from := core.A1; from <= core.H8; from++ {
		// Skip non-starting squares.
		if !fromBB.Get(from) {
			continue
		}

		for _, t := range translations {
			// Repeatedly apply the translation until it fails.
			for to, ok := translate(from, t); ok; to, ok = translate(to, t) {
				// Always generate the move, even if there's a piece already on
				// the destination square. Friendly captures are removed later.
				moves = append(moves, core.Move{From: from, To: to})
				// We're blocked, so try the next translation direction.
				if p.Board.IsOccupied(to) {
					break
				}
			}
		}
	}

	// Remove moves that capture friendly pieces.
	moves = slices.DeleteFunc(moves, func(m core.Move) bool {
		piece, ok := p.Board.Get(m.To)
		return ok && piece.Color() == p.SideToMove
	})

	return moves
}

// bishopAttacks returns available bishop attacks, without considering checks.
func bishopAttacks(p core.Position) []core.Move {
	return slidingAttacks(p, core.Bishop)
}

// rookAttacks returns available rook attacks, without considering checks.
func rookAttacks(p core.Position) []core.Move {
	return slidingAttacks(p, core.Rook)
}

// queenAttacks returns available queen attacks, without considering checks.
func queenAttacks(p core.Position) []core.Move {
	return slidingAttacks(p, core.Queen)
}

// kingAttacks returns available king attacks, without considering checks.
func kingAttacks(p core.Position) []core.Move {
	return nil // TODO
}

// isEnemyKingTargeted returns true if the enemy king is targeted by an attack.
//
// Note that a king may target an adjacent king.
func isEnemyKingTargeted(p core.Position) bool {
	s := p.EnemyKing()

	moves := slices.Concat(
		pawnAttacks(p),
		knightAttacks(p),
		bishopAttacks(p),
		rookAttacks(p),
		queenAttacks(p),
		kingAttacks(p),
	)

	for _, m := range moves {
		if m.To == s {
			return true
		}
	}
	return false
}
