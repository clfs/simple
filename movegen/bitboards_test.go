package movegen

import (
	"testing"

	"github.com/clfs/simple/core"
)

func TestBitboardMoveCounts(t *testing.T) {
	cases := []struct {
		bbs  [64]core.Bitboard
		want int
	}{
		{whitePawnPushes, 6*8 + 8},         // 56
		{blackPawnPushes, 6*8 + 8},         // 56
		{whitePawnAttacks, 6 * 7 * 2},      // 84
		{blackPawnAttacks, 6 * 7 * 2},      // 84
		{knightAttacks, 8 * 7 * 6},         // 336
		{bishopAttacks, 0},                 // ?
		{rookAttacks, 14 * 64},             // 896
		{queenAttacks, 0},                  // ?
		{kingAttacks, 7*7*8 + 6*4*5 + 4*3}, // 544
	}
	for i, tc := range cases {
		var got int
		for _, bb := range tc.bbs {
			got += bb.Count()
		}
		if got != tc.want {
			t.Errorf("#%d: got %d, want %d", i, got, tc.want)
		}
	}
}
