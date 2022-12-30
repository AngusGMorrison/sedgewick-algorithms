package q6_recursive_select

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// Select returns the k+1th-smallest element.
func Select[S ~[]E, E constraints.Ordered](s S, k int) E {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return slct(s, k, 0, len(s)-1)
}

func slct[S ~[]E, E constraints.Ordered](s S, k, lo, hi int) E {
	if lo >= hi {
		return s[k]
	}

	pivot := partition(s, lo, hi)
	if pivot == k {
		return s[k]
	} else if pivot < k {
		return slct(s, k, pivot+1, hi)
	} else {
		return slct(s, k, lo, pivot-1)
	}
}

func partition[S ~[]E, E constraints.Ordered](s S, lo, hi int) int {
	for i := lo + 1; i < hi; i++ {
		if s[i] < s[hi] {
			s[i], s[lo] = s[lo], s[i]
			lo++
		}
	}

	s[lo], s[hi] = s[hi], s[lo]
	return lo
}
