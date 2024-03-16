package epd

import (
	"fmt"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/pcn"
)

func ExampleDecode() {
	s := `4k3/8/P7/8/8/8/8/4K3 w - - bm a7; c0 "My example";`

	p, ops, err := Decode(s)
	if err != nil {
		// Handle error.
	}

	// Decode sets default values for the half move clock and full move number
	// if the EPD string does not specify them.
	fmt.Println("Half move clock:", p.HalfMoveClock)
	fmt.Println("Full move number:", p.FullMoveNumber)

	for _, op := range ops {
		fmt.Printf("%#v\n", op)
	}

	// Output:
	// Half move clock: 0
	// Full move number: 1
	// epd.Op{Opcode:"bm", Operand:"a7"}
	// epd.Op{Opcode:"c0", Operand:"\"My example\""}
}

func ExampleEncode() {
	p := core.NewPosition()

	p.Make(pcn.MustDecode("e2e4"))
	p.Make(pcn.MustDecode("e7e5"))

	// By default, Encode discards the position's half move clock and full move
	// number. To include them, pass the corresponding operations.
	ops := []Op{
		{OpcodeHalfMoveClock, fmt.Sprint(p.HalfMoveClock)},
		{OpcodeFullMoveNumber, fmt.Sprint(p.FullMoveNumber)},
	}

	fmt.Println(Encode(p, ops))
	// Output:
	// rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 hmvc 0; fmvn 2;
}
