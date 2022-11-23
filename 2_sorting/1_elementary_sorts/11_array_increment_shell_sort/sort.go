package sort

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Shell[S ~[]E, E constraints.Ordered](s S) {
	increment := []int{1, 4, 13, 40, 121, 364, 1093}
	for idx := initialIncrementIdx(len(s)); idx >= 0; idx-- {
		h := increment[idx]
		for i := h; i < len(s); i++ {
			for j := i; j >= h; j -= h {
				if less(s[j], s[j-h]) {
					swap(s, j, j-h)
				}
			}
		}
	}
}

// initialIncrementIdx calculates the index of the first value of h based on the length of the input
// array. Since the kth element in the geometric sequence calculating h is given by, (3^k-1)/2,
// finding the index is equivalent to solving for k.
func initialIncrementIdx(len int) int {
	return int(math.Log10(float64(len/3)) / math.Log10(3))
}

func less[E constraints.Ordered](a, b E) bool {
	return a < b
}

func swap[S ~[]E, E any](s S, i, j int) {
	s[i], s[j] = s[j], s[i]
}
