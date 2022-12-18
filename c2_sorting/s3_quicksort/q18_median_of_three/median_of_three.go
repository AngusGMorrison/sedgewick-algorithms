package q18_median_of_three

import "math/rand"

func QuickSort(s []int) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	quickSort(s, 0, len(s)-1)
}

func quickSort(s []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivot := partition(s, lo, hi)
	quickSort(s, lo, pivot-1)
	quickSort(s, pivot+1, hi)
}

func partition(s []int, lo, hi int) int {
	// Sort the first, middle and last elements. The middle element (the median) will become our
	// pivot.
	pivot := lo + (hi-lo)/2
	if s[pivot] < s[lo] {
		s[pivot], s[lo] = s[lo], s[pivot]
	}
	if s[hi] < s[lo] {
		s[hi], s[lo] = s[lo], s[hi]
	}
	if s[hi] < s[pivot] {
		s[hi], s[pivot] = s[pivot], s[hi]
	}

	// The first, middle and last elements are now sorted relative to each other. Move the median of
	// the three to the second-last position from the end (since the item currently in hi is greater
	// than or equal to it).
	s[hi-1], s[pivot] = s[pivot], s[hi-1]
	pivot = hi - 1

	for i := lo; i < pivot; i++ {
		if s[i] < s[pivot] {
			s[i], s[lo] = s[lo], s[i]
			lo++
		}
	}
	s[pivot], s[lo] = s[lo], s[pivot]
	return lo
}
