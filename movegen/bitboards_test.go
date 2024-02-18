package movegen

import (
	"testing"

	"github.com/clfs/simple/core"
)

func TestBitboardMoveCounts(t *testing.T) {
	cases := []struct {
		in   [64]core.Bitboard
		want int
	}{
		// 8 single pushes from 6 possible ranks, plus 8 double pushes.
		{whitePawnPushes, 56},
		{blackPawnPushes, 56},

		// 7 left attacks and 7 right attacks from 6 possible ranks.
		{whitePawnAttacks, 84},
		{blackPawnAttacks, 84},

		// https://math.stackexchange.com/a/2204782
		// n(n-1)(n-2), where n=8
		{knightAttacks, 336},

		// https://math.stackexchange.com/a/3054557
		// 2n(n-1)(2n-1)/3, where n=8
		{bishopAttacks, 560},

		// 7 vertical attacks and 7 horizontal attacks from each square.
		{rookAttacks, 896},

		// Add the bishop and rook attacks.
		{queenAttacks, 1456},

		// 36 squares have 8 attacks, 24 edges have 5, and 4 corners have 3.
		{kingAttacks, 420},
	}
	for i, tc := range cases {
		var got int
		for _, bb := range tc.in {
			got += bb.Count()
		}
		if got != tc.want {
			t.Errorf("#%d: got %d, want %d", i, got, tc.want)
		}
	}
}
