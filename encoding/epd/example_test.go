package epd

import (
	"fmt"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/pcn"
)

func ExampleDecode() {

}

func ExampleEncode() {
	ops := []Op{
		{Opcode: "foo", Args: "bar"},
		{Opcode: "c0", Args: `"my comment"`},
	}

	// Indian Defense: 1. d4 Nf6
	p := core.NewPosition()
	p.Make(pcn.MustDecode("d2d4"))
	p.Make(pcn.MustDecode("g8f6"))

	fmt.Println(Encode(p, ops))

	// Output:
	// rn1qkb1r/ppp1pppp/8/3p4/3Pn3/8/PPP2PPP/RNBQKBNR w KQkq - foo bar; c0 "my comment";
}
