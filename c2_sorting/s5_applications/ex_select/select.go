package ex_select

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// Select returns the k+1th-smallest element in linear time.
func Select[S ~[]E, E constraints.Ordered](s S, k int) E {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	lo := 0
	hi := len(s) - 1
	for lo < hi {
		// Partition the array such that everything left of pivot is smaller than pivot (which is
		// the pivoth-smallest element), and everything to the right is larger. Partition
		// progressively smaller arrays on each iteration until the k+1th-smallest element is found.
		// In the average case, the number of elements to partition is halved on each iteration. N +
		// N/2 + ... ~= 2N.
		pivot := partition(s, lo, hi)
		if pivot == k {
			return s[k]
		} else if pivot > k {
			hi = pivot - 1
		} else if pivot < k {
			lo = pivot + 1
		}
	}
	return s[k] // the k+1th-smallest is the element where lo == hi
}

func partition[S ~[]E, E constraints.Ordered](s S, lo, hi int) int {
	var i int
	for i = lo; i < hi; i++ {
		if s[i] < s[hi] { // use s[hi] as the pivot
			s[i], s[lo] = s[lo], s[i]
			lo++
		}
	}
	s[i], s[hi] = s[hi], s[i]
	return i
}
