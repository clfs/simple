package fen

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/clfs/simple/core"
)

// numRegexp matches any non-negative integer.
var numRegexp = regexp.MustCompile(`^(0|[1-9]\d*)$`)

var decodePiece = map[rune]core.Piece{
	'P': core.WhitePawn,
	'N': core.WhiteKnight,
	'B': core.WhiteBishop,
	'R': core.WhiteRook,
	'Q': core.WhiteQueen,
	'K': core.WhiteKing,
	'p': core.BlackPawn,
	'n': core.BlackKnight,
	'b': core.BlackBishop,
	'r': core.BlackRook,
	'q': core.BlackQueen,
	'k': core.BlackKing,
}

func decodeSquare(s string) (core.Square, bool) {
	if len(s) != 2 {
		return 0, false
	}

	f := core.File(s[0] - 'a')
	r := core.Rank(s[1] - '1')

	if f < core.FileA || f > core.FileH || r < core.Rank1 || r > core.Rank8 {
		return 0, false
	}
	return core.NewSquare(f, r), true
}

// Decode decodes a FEN string and returns the position it represents.
func Decode(s string) (core.Position, error) {
	var p core.Position

	fields := strings.Split(s, " ")
	if n := len(fields); n != 6 {
		return core.Position{}, fmt.Errorf("invalid number of fields: %d", n)
	}

	// TODO(cfiguereosupran): See if there's a way to iterate over the board
	// squares in order, rather than jump around according to the FEN input.

	// Board.
	offset := int(core.A8) // top left corner
	rows := strings.Split(fields[0], "/")
	if len(rows) != 8 {
		return core.Position{}, fmt.Errorf("invalid number of board rows: %d", len(rows))
	}
	for i, row := range rows {
		var numPrev bool
		for _, r := range row {
			switch r {
			case '1', '2', '3', '4', '5', '6', '7', '8':
				if numPrev {
					return core.Position{}, fmt.Errorf("two consecutive numbers in board row: %s", row)
				}
				offset += int(r - '0') // advance rightwards by n
				numPrev = true
			default:
				piece, ok := decodePiece[r]
				if !ok {
					return core.Position{}, fmt.Errorf("invalid board piece: %c", r)
				}
				p.Board.SetOnEmpty(piece, core.Square(offset))
				offset++ // advance rightwards by 1
				numPrev = false
			}
		}

		// Were all eight squares accounted for?
		if offset != 8*(8-i) {
			return core.Position{}, fmt.Errorf("invalid board row length: %s", row)
		}

		offset -= 16 // advance down by 2
	}

	// Side to move.
	switch fields[1] {
	case "w":
		p.SideToMove = core.White
	case "b":
		p.SideToMove = core.Black
	default:
		return core.Position{}, fmt.Errorf("invalid side to move: %s", fields[1])
	}

	// Castling rights.
	switch fields[2] {
	case "-":
	case "K":
		p.WhiteOO = true
	case "Q":
		p.WhiteOOO = true
	case "k":
		p.BlackOO = true
	case "q":
		p.BlackOOO = true
	case "KQ":
		p.WhiteOO, p.WhiteOOO = true, true
	case "Kk":
		p.WhiteOO, p.BlackOO = true, true
	case "Kq":
		p.WhiteOO, p.BlackOOO = true, true
	case "Qk":
		p.WhiteOOO, p.BlackOO = true, true
	case "Qq":
		p.WhiteOOO, p.BlackOOO = true, true
	case "kq":
		p.BlackOO, p.BlackOOO = true, true
	case "KQk":
		p.WhiteOO, p.WhiteOOO, p.BlackOO = true, true, true
	case "KQq":
		p.WhiteOO, p.WhiteOOO, p.BlackOOO = true, true, true
	case "Kkq":
		p.WhiteOO, p.BlackOO, p.BlackOOO = true, true, true
	case "Qkq":
		p.WhiteOOO, p.BlackOO, p.BlackOOO = true, true, true
	case "KQkq":
		p.WhiteOO, p.WhiteOOO, p.BlackOO, p.BlackOOO = true, true, true, true
	default:
		return core.Position{}, fmt.Errorf("invalid castling rights: %s", fields[2])
	}

	// En passant square.
	if fields[3] != "-" {
		sq, ok := decodeSquare(fields[3])
		if !ok {
			return core.Position{}, fmt.Errorf("invalid e.p. square: %s", fields[3])
		}

		switch sq.Rank() {
		case core.Rank3:
			if p.SideToMove == core.White {
				return core.Position{}, fmt.Errorf("invalid e.p. square for white: %s", fields[3])
			}
		case core.Rank6:
			if p.SideToMove == core.Black {
				return core.Position{}, fmt.Errorf("invalid e.p. square for black: %s", fields[3])
			}
		default:
			return core.Position{}, fmt.Errorf("invalid rank for e.p. square: %s", fields[3])
		}

		p.EnPassant = sq
	}

	// Half move clock.
	if !numRegexp.MatchString(fields[4]) {
		return core.Position{}, fmt.Errorf("invalid half move clock: %s", fields[4])
	}
	hmc, err := strconv.Atoi(fields[4])
	if err != nil || hmc < 0 {
		return core.Position{}, fmt.Errorf("invalid half move clock: %s", fields[4])
	}
	p.HalfMoveClock = hmc

	// Full move number.
	if !numRegexp.MatchString(fields[5]) {
		return core.Position{}, fmt.Errorf("invalid full move number: %s", fields[5])
	}
	fmn, err := strconv.Atoi(fields[5])
	if err != nil || fmn <= 0 {
		return core.Position{}, fmt.Errorf("invalid full move number: %s", fields[5])
	}
	p.FullMoveNumber = fmn

	return p, nil
}
