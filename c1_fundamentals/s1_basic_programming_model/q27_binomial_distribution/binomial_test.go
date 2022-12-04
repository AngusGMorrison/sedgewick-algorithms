package q27_binomial_distribution

import (
	"fmt"
	"math"
	"testing"
)

func Test_dynamicBinomialRecursive(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		n, k int
		p    float64
		want float64
	}{
		{n: 0, k: 0, p: 1.0, want: 1.0},
		{n: 1, k: -1, p: 1.0, want: 0.0},
		{n: -1, k: 0, p: 1.0, want: 0.0},
		{n: 1, k: 0, p: 1.0, want: 0.0},
		{n: 5, k: 2, p: 0.25, want: 0.263671875},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("dynamicBinomial(%d, %d, %.3f)", tc.n, tc.k, tc.p)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := dynamicBinomialRecursive(tc.n, tc.k, tc.p); !approxEqual(got, tc.want) {
				t.Errorf("%s: want %.3f, got %.3f", name, tc.want, got)
			}
		})
	}
}

func Test_dynamicBinomialLoop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		n, k int
		p    float64
		want float64
	}{
		{n: 0, k: 0, p: 1.0, want: 1.0},
		{n: 1, k: -1, p: 1.0, want: 0.0},
		{n: -1, k: 0, p: 1.0, want: 0.0},
		{n: 1, k: 0, p: 1.0, want: 0.0},
		{n: 5, k: 2, p: 0.25, want: 0.263671875},
	}

	for _, tc := range testCases {
		tc := tc
		name := fmt.Sprintf("dynamicBinomialLoop(%d, %d, %.3f)", tc.n, tc.k, tc.p)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := dynamicBinomialLoop(tc.n, tc.k, tc.p); !approxEqual(got, tc.want) {
				t.Errorf("%s: want %.3f, got %.3f", name, tc.want, got)
			}
		})
	}
}

func Benchmark_binomial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binomial(10, 5, 0.25)
	}
}

func Benchmark_dynamicBinomialRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dynamicBinomialRecursive(10, 5, 0.25)
	}
}

func Benchmark_dynamicBinomialLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dynamicBinomialLoop(10, 5, 0.25)
	}
}

const float64EqualityThreshold = 1e-7

func approxEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) <= float64EqualityThreshold
}
