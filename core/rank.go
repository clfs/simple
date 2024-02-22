package core

import "fmt"

// A Rank is a row on the chess board.
type Rank uint64

// Rank constants.
const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

func (r Rank) String() string {
	switch r {
	case Rank1:
		return "Rank1"
	case Rank2:
		return "Rank2"
	case Rank3:
		return "Rank3"
	case Rank4:
		return "Rank4"
	case Rank5:
		return "Rank5"
	case Rank6:
		return "Rank6"
	case Rank7:
		return "Rank7"
	case Rank8:
		return "Rank8"
	default:
		return fmt.Sprintf("Rank(%d)", r)
	}
}

func (r Rank) Valid() bool {
	return r >= Rank1 && r <= Rank8
}
