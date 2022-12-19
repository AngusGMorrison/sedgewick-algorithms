package ex_three_way_partition

import "math/rand"

func QuickSort(s []int) {
	quickSort(s, 0, len(s)-1)
}

func quickSort(s []int, lo, hi int) {
	if lo >= hi {
		return
	}

	// Choose a pivot and move it to the start.
	pivotIdx := lo + rand.Intn(hi-lo+1)
	s[lo], s[pivotIdx] = s[pivotIdx], s[lo]
	pivot := s[lo]

	// Partition the array into three sections: items less than the pivot at the left-hand end,
	// items equal to the pivot in the middle, and items greater than the pivot at the right-hand
	// end.
	// |---------------|--------------|--------------|--------------|
	// lo (< pivot)    lt (== pivot)  i (unchecked)  gt (> pivot)   hi
	lt := lo      // elements less than the pivot are swapped with the less-than pointer. lt always points to the pivot
	i := lo + 1   // points to the current unchecked item
	gt := hi      // elements greater than the pivot are swapped with the greater-than pointer
	for i <= gt { // caution: the element at the current value of gt has not yet been checked; we must check it before terminating the loop
		if s[i] < pivot {
			s[i], s[lt] = s[lt], s[i]
			lt++
			i++
		} else if s[i] > pivot {
			s[i], s[gt] = s[gt], s[i]
			gt--
			// i stays the same, since we haven't previously checked the value that we just swapped into s[i] from s[gt]
		} else {
			i++ // s[i] == pivot, so keep it in place
		}
	}

	// The middle section containing pivot elements is already in sorted order, so we can skip it in
	// the recursive calls to sort. This is why three-way partitioning is so efficient for slices
	// containing large numbers of duplicates.
	quickSort(s, lo, lt-1) // lt points to the first element equal to pivot
	quickSort(s, gt+1, hi) // gt points to the last element equal to pivot
}
