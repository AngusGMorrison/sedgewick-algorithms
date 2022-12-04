package q14_log

import (
	"fmt"
	"testing"
)

func Test_lg(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		n, want int
	}{
		{-5, -1},
		{0, -1},
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{8, 3},
		{16, 4},
		{32, 5},
		{64, 6},
		{128, 7},
		{256, 8},
		{257, 8},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("lg(%d)", tc.n)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := lg(tc.n); got != tc.want {
				t.Errorf("%s: want %d, got %d", name, tc.n, got)
			}
		})
	}
}
