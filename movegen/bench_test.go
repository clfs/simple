package movegen

import (
	"fmt"
	"testing"

	"github.com/clfs/simple/core"
)

func BenchmarkPerft(b *testing.B) {
	p := core.NewPosition()
	b.ResetTimer()

	for depth := range 4 {
		b.Run(fmt.Sprint(depth), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Perft(p, depth)
			}
		})
	}
}
