package q34_hot_or_cold

import (
	"fmt"
	"math"
	"testing"
)

func Test_solver_Solve(t *testing.T) {
	t.Parallel()

	testCases := []int{8, 16, 32, 64, 128, 256, 512, 1024, 2048}
	var messages []string

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d", tc), func(t *testing.T) {
			var guesses []int
			for i := 1; i <= tc; i++ {
				t.Run(fmt.Sprintf("target=%d", i), func(t *testing.T) {
					s := solver{
						n:      tc,
						target: i,
					}
					if got := s.Solve(); got != i {
						t.Errorf("want %d, got %d in %d guesses", i, got, s.guesses)
					}
					guesses = append(guesses, s.guesses)
				})
			}
			messages = append(messages, fmt.Sprintf("Solved n=%d with max guesses: %d", tc, sliceMax(guesses)))
		})
	}

	for _, msg := range messages {
		t.Log(msg)
	}
}

func sliceMax(a []int) int {
	max := math.MinInt64
	for _, elem := range a {
		if elem > max {
			max = elem
		}
	}

	return max
}
