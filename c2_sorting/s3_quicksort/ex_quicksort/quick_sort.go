package ex_quicksort

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

func QuickSort[S ~[]E, E constraints.Ordered](s S) {
	quickSort(s, 0, len(s)-1)
}

func quickSort[S ~[]E, E constraints.Ordered](s S, lo, hi int) {
	if lo >= hi {
		return
	}

	pivot := partition(s, lo, hi)
	quickSort(s, lo, pivot-1)
	quickSort(s, pivot+1, hi)
}

func partition[S ~[]E, E constraints.Ordered](s S, lo, hi int) int {
	pivot := lo + rand.Intn(hi-lo+1)
	s[pivot], s[hi] = s[hi], s[pivot] // move pivot to the end

	for i := lo; i < hi; i++ {
		if s[i] < s[hi] { // if the current item is less than the pivot
			s[i], s[lo] = s[lo], s[i] // pile elements smaller than the pivot on the left-hand side of the array
			lo++                      // advance the left-hand counter so that these smaller elements can no longer be swapped
		}
	}

	s[lo], s[hi] = s[hi], s[lo]
	return lo
}
