package improvements

import (
	"golang.org/x/exp/constraints"
)

const insertionSortThreshold = 15

func MergeSort[S ~[]E, E constraints.Ordered](s S) {
	mergeSort(s, make(S, len(s)), 0, len(s)-1, false)
}

// mergeSort combines two of the three suggested improvements:
//  1. Switch to insertion sort for small arrays
//  2. Avoid copying the array to aux and back on each merge by alternating the array that merge
//     outputs to.
//
// Does not implement the third suggestion, to skip the merge call if the two subarrays are already
// in sorted order, since this is incompatible with #2, which relies on the merge call to move data
// to the array the parent call expects to find it in.
func mergeSort[S ~[]E, E constraints.Ordered](s, aux S, lo, hi int, toAux bool) {
	if lo > hi { // preferred to hi-lo == 0, since hi may be -1
		return
	}
	// Base case: insertion sort subarrays at or below the length threshold and copy the sorted
	// subarray to aux (copying all insertion-sorted subarrays is equivalent to copying all n
	// elements in the original slice once). This copy allows us to begin alternating s and aux as
	// merge destinations.
	if hi-lo <= insertionSortThreshold {
		insertionSort(s, lo, hi)
		if toAux { // avoid useless copy when original array is shorter than threshold
			copy(aux[lo:hi+1], s[lo:hi+1])
		}
		return
	}

	mid := lo + (hi-lo)/2
	// Alternate the array that each level's recursive calls output to to eliminate copying at each
	// level.
	mergeSort(s, aux, lo, mid, !toAux)
	mergeSort(s, aux, mid+1, hi, !toAux)
	if toAux {
		merge(s, aux, lo, mid, hi) // merge s into aux
	} else {
		merge(aux, s, lo, mid, hi) // merge aux into s
	}
}

func merge[S ~[]E, E constraints.Ordered](in, out S, lo, mid, hi int) {
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if j > hi {
			out[k] = in[i]
			i++
		} else if i > mid {
			out[k] = in[j]
			j++
		} else if less(in[i], in[j]) {
			out[k] = in[i]
			i++
		} else {
			out[k] = in[j]
			j++
		}
	}
}

func insertionSort[S ~[]E, E constraints.Ordered](s S, lo, hi int) {
	if lo >= hi {
		return
	}
	for i := lo + 1; i <= hi; i++ {
		elem := s[i]
		var j int
		for j = i; j > lo && less(elem, s[j-1]); j-- {
			s[j] = s[j-1]
		}
		s[j] = elem
	}
}

func less[E constraints.Ordered](a, b E) bool {
	return a < b
}
