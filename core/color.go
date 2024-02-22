package core

// A Color is white or black.
type Color bool

// Color constants.
const (
	White Color = false
	Black Color = true
)

func (c Color) String() string {
	if c {
		return "Black"
	}
	return "White"
}

// Other returns the other color.
func (c Color) Other() Color {
	return !c
}

// Uint64 returns 0 for white and 1 for black.
func (c Color) Uint64() uint64 {
	if c {
		return 1
	}
	return 0
}
