package q19_fibonacci

import (
	"fmt"
	"testing"
)

func Test_naive(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("naive(%d)", tc.n)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := naive(tc.n); got != tc.want {
				t.Errorf("%s: want %d, got %d", name, tc.want, got)
			}
		})
	}
}

func Test_dynamicRecursive(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("dynamic(%d)", tc.n)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := dynamicRecursive(tc.n); got != tc.want {
				t.Errorf("%s: want %d, got %d", name, tc.want, got)
			}
		})
	}
}

func Test_dynamicLoop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("dynamic(%d)", tc.n)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := dynamicLoop(tc.n); got != tc.want {
				t.Errorf("%s: want %d, got %d", name, tc.want, got)
			}
		})
	}
}

func Benchmark_naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		naive(20)
	}
}

func Benchmark_dynamicRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dynamicRecursive(20)
	}
}

func Benchmark_dynamicLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dynamicLoop(20)
	}
}
