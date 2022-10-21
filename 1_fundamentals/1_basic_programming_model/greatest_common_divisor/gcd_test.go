package gcd

import (
	"fmt"
	"testing"
)

func Test_GCD(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		p, q, want int
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
		{2, 1, 1},
		{1, 2, 1},
		{5, 10, 5},
		{10, 5, 5},
		{60, 12, 12},
		{60, 12, 12},
		{13, 17, 1},
		{17, 13, 1},
		{270, 33, 3},
		{33, 270, 3},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("GCD(%d, %d)", tc.p, tc.q)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := GCD(tc.p, tc.q); got != tc.want {
				t.Errorf("%s: want %d, got %d", name, tc.want, got)
			}
		})
	}
}
