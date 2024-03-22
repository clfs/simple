// Package search implements move search.
package search

import (
	"context"
	"errors"
	"math"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/eval"
	"github.com/clfs/simple/movegen"
)

func negamax(p core.Position, depth int) int {
	if depth <= 0 {
		return eval.Eval(p)
	}

	score := math.MinInt
	for _, m := range movegen.LegalMoves(p) {
		child := p
		child.Make(m)
		score = max(score, -negamax(child, depth-1))
	}
	return score
}

// ErrNoLegalMoves is returned by Search when there are no legal moves in a
// position.
var ErrNoLegalMoves = errors.New("no legal moves")

// Search searches for the best move in a position.
func Search(ctx context.Context, p core.Position, best chan<- core.Move) error {
	for depth := 3; ; depth++ {
		var (
			bestScore = math.MinInt
			bestMove  core.Move
		)

		moves := movegen.LegalMoves(p)

		if len(moves) == 0 {
			return ErrNoLegalMoves
		}

		for _, m := range moves {
			child := p
			child.Make(m)

			if s := -negamax(child, depth-1); s > bestScore {
				bestScore, bestMove = s, m
			}
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case best <- bestMove:
		}
	}
}
