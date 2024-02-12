package fen

import (
	"fmt"
	"strings"

	"github.com/clfs/simple/core"
)

// encodeSquare encodes a square in lower case.
func encodeSquare(s core.Square) string {
	return fmt.Sprintf("%s", strings.ToLower(s.String()))
}

// Encode encodes a position as a FEN string.
func Encode(p core.Position) string {
	var b strings.Builder

	// Board.
	for r := core.Rank8; r >= core.Rank1; r-- {
		skip := 0
		for f := core.FileA; f <= core.FileH; f++ {
			sq := core.NewSquare(f, r)
			piece, ok := p.Board.Get(sq)
			if !ok {
				skip++
				continue
			}
			if skip > 0 {
				fmt.Fprintf(&b, "%d", skip)
				skip = 0
			}
			switch piece {
			case core.WhitePawn:
				fmt.Fprintf(&b, "P")
			case core.WhiteKnight:
				fmt.Fprintf(&b, "N")
			case core.WhiteBishop:
				fmt.Fprintf(&b, "B")
			case core.WhiteRook:
				fmt.Fprintf(&b, "R")
			case core.WhiteQueen:
				fmt.Fprintf(&b, "Q")
			case core.WhiteKing:
				fmt.Fprintf(&b, "K")
			case core.BlackPawn:
				fmt.Fprintf(&b, "p")
			case core.BlackKnight:
				fmt.Fprintf(&b, "n")
			case core.BlackBishop:
				fmt.Fprintf(&b, "b")
			case core.BlackRook:
				fmt.Fprintf(&b, "r")
			case core.BlackQueen:
				fmt.Fprintf(&b, "q")
			case core.BlackKing:
				fmt.Fprintf(&b, "k")
			}
		}
		if skip > 0 {
			fmt.Fprintf(&b, "%d", skip)
		}
		if r != core.Rank1 {
			fmt.Fprintf(&b, "/")
		}
	}

	fmt.Fprintf(&b, " ")

	// Side to move.
	if p.SideToMove == core.White {
		fmt.Fprintf(&b, "w")
	} else {
		fmt.Fprintf(&b, "b")
	}

	fmt.Fprintf(&b, " ")

	// Castling rights.
	if !p.WhiteOO && !p.WhiteOOO && !p.BlackOO && !p.BlackOOO {
		fmt.Fprintf(&b, "-")
	} else {
		if p.WhiteOO {
			fmt.Fprintf(&b, "K")
		}
		if p.WhiteOOO {
			fmt.Fprintf(&b, "Q")
		}
		if p.BlackOO {
			fmt.Fprintf(&b, "k")
		}
		if p.BlackOOO {
			fmt.Fprintf(&b, "q")
		}
	}

	fmt.Fprintf(&b, " ")

	// En passant square.
	if p.EnPassant == 0 {
		fmt.Fprintf(&b, "-")
	} else {
		fmt.Fprintf(&b, "%s", encodeSquare(p.EnPassant))
	}

	fmt.Fprintf(&b, " ")

	// Half move clock.
	fmt.Fprintf(&b, "%d", p.HalfMoveClock)

	fmt.Fprintf(&b, " ")

	// Full move counter.
	fmt.Fprintf(&b, "%d", p.FullMoveCounter)

	return b.String()
}
