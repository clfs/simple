package movegen

import (
	"fmt"

	"github.com/clfs/simple/core"
	"github.com/clfs/simple/encoding/pcn"
)

func ExamplePerft() {
	p := core.NewPosition()

	// Count the number of leaf nodes at each depth in the move tree.
	for i := range 4 {
		fmt.Printf("Perft(%d) = %d\n", i, Perft(p, i))
	}
	// Output:
	// Perft(0) = 1
	// Perft(1) = 20
	// Perft(2) = 400
	// Perft(3) = 8902
}

func ExampleDivide() {
	var (
		p = core.NewPosition()
		m = pcn.MustDecode("e2e4")
	)

	// How many leaf nodes at depth 3 start with e2e4?
	fmt.Println("e2e4:", Divide(p, 3)[m])
	// After e2e4, how many leaf nodes are at depth 2?
	p.Make(m)
	fmt.Println("e2e4:", Perft(p, 2))
	// Output:
	// e2e4: 600
	// e2e4: 600
}
