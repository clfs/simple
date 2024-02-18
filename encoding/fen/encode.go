package fen

import (
	"fmt"
	"strings"

	"github.com/clfs/simple/core"
)

// encodeSquare encodes a square in lower case.
func encodeSquare(s core.Square) string {
	return strings.ToLower(s.String())
}

var encodePiece = map[core.Piece]rune{
	core.WhitePawn:   'P',
	core.WhiteKnight: 'N',
	core.WhiteBishop: 'B',
	core.WhiteRook:   'R',
	core.WhiteQueen:  'Q',
	core.WhiteKing:   'K',
	core.BlackPawn:   'p',
	core.BlackKnight: 'n',
	core.BlackBishop: 'b',
	core.BlackRook:   'r',
	core.BlackQueen:  'q',
	core.BlackKing:   'k',
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
			b.WriteRune(encodePiece[piece])
		}
		if skip > 0 {
			fmt.Fprintf(&b, "%d", skip)
		}
		if r != core.Rank1 {
			b.WriteRune('/')
		}
	}

	b.WriteRune(' ')

	// Side to move.
	if p.SideToMove == core.White {
		b.WriteRune('w')
	} else {
		b.WriteRune('b')
	}

	b.WriteRune(' ')

	// Castling rights.
	if !p.WhiteOO && !p.WhiteOOO && !p.BlackOO && !p.BlackOOO {
		b.WriteRune('-')
	} else {
		if p.WhiteOO {
			b.WriteRune('K')
		}
		if p.WhiteOOO {
			b.WriteRune('Q')
		}
		if p.BlackOO {
			b.WriteRune('k')
		}
		if p.BlackOOO {
			b.WriteRune('q')
		}
	}

	b.WriteRune(' ')

	// En passant square.
	if p.EnPassant == 0 {
		b.WriteRune('-')
	} else {
		fmt.Fprintf(&b, "%s", encodeSquare(p.EnPassant))
	}

	b.WriteRune(' ')

	// Half move clock.
	fmt.Fprintf(&b, "%d", p.HalfMoveClock)

	b.WriteRune(' ')

	// Full move counter.
	fmt.Fprintf(&b, "%d", p.FullMoveNumber)

	return b.String()
}
