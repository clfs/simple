// Package epd implements encoding and decoding of Extended Position Description
// (EPD) as defined in "Standard: Portable Game Notation Specification and
// Implementation Guide", revision 1994.03.12.
package epd

// An Op is an EPD operation.
type Op struct {
	Opcode string
	Args   string
}
