package q2_threesum_overflow

import (
	"fmt"
	"math"
	"testing"
)

func Test_ThreeSum(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input     []int
		wantCount int
	}{
		{
			input:     []int{math.MaxInt, 1, math.MinInt},
			wantCount: 1,
		},
		{
			input:     []int{1, math.MaxInt, math.MinInt},
			wantCount: 1,
		},
		{
			input:     []int{math.MaxInt - 5, 5, -math.MaxInt},
			wantCount: 1,
		},
		{
			input:     []int{math.MaxInt - 5, 5, -math.MaxInt},
			wantCount: 1,
		},
		{
			// If overflows could occur, this would result in a count of 1 (evaluates to
			// (math.MinInt+1)+math.MaxInt, which equals 0).
			input:     []int{math.MaxInt, 2, math.MaxInt},
			wantCount: 0,
		},
		{
			input:     []int{math.MaxInt, math.MaxInt, 2},
			wantCount: 0,
		},
		{
			input:     []int{2, math.MaxInt, math.MaxInt},
			wantCount: 0,
		},
		{
			input:     []int{1, 2, -3},
			wantCount: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("ThreeSum(%v)", tc.input)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := ThreeSum(tc.input); got != tc.wantCount {
				t.Errorf("want count %d, got %d", tc.wantCount, got)
			}
		})
	}
}

func Test_ThreeSumFast(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input     []int
		wantCount int
	}{
		{
			input:     []int{math.MaxInt, 1, math.MinInt},
			wantCount: 1,
		},
		{
			input:     []int{1, math.MaxInt, math.MinInt},
			wantCount: 1,
		},
		{
			input:     []int{math.MaxInt - 5, 5, -math.MaxInt},
			wantCount: 1,
		},
		{
			input:     []int{math.MaxInt - 5, 5, -math.MaxInt},
			wantCount: 1,
		},
		{
			// If overflows could occur, this would result in a count of 1 (evaluates to
			// (math.MinInt+1)+math.MaxInt, which equals 0).
			input:     []int{math.MaxInt, 2, math.MaxInt},
			wantCount: 0,
		},
		{
			input:     []int{math.MaxInt, math.MaxInt, 2},
			wantCount: 0,
		},
		{
			input:     []int{2, math.MaxInt, math.MaxInt},
			wantCount: 0,
		},
		{
			input:     []int{1, 2, -3},
			wantCount: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("ThreeSumFast(%v)", tc.input)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := ThreeSumFast(tc.input); got != tc.wantCount {
				t.Errorf("want count %d, got %d", tc.wantCount, got)
			}
		})
	}
}
