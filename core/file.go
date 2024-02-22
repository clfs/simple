package core

import "fmt"

// A File is a column on the chess board.
type File uint64

// File constants.
const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

func (f File) String() string {
	switch f {
	case FileA:
		return "FileA"
	case FileB:
		return "FileB"
	case FileC:
		return "FileC"
	case FileD:
		return "FileD"
	case FileE:
		return "FileE"
	case FileF:
		return "FileF"
	case FileG:
		return "FileG"
	case FileH:
		return "FileH"
	default:
		return fmt.Sprintf("File(%d)", f)
	}
}

func (f File) Valid() bool {
	return f >= FileA && f <= FileH
}
