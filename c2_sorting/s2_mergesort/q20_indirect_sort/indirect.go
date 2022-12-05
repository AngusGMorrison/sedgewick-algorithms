package q20_indirect_sort

func IndirectSort(s []int) []int {
	aux := make([]int, len(s))
	// idx contains the current positions of elements in s, like pointers. To start with, the 0th
	// element in s is at idx[0]. When s is "sorted", these pointers will move to different
	// positions in idx, but can still be used to look up the original elements in s. E.g. if idx[4]
	// == 0, then s[0] is the 5th-smallest element in s.
	idx := make([]int, len(s))
	for i := 0; i < len(idx); i++ {
		idx[i] = i
	}

	sort(s, idx, aux, 0, len(s)-1)

	return idx
}

func sort(s, idx, aux []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	sort(s, idx, aux, lo, mid)
	sort(s, idx, aux, mid+1, hi)
	merge(s, idx, aux, lo, mid, hi)
}

func merge(s, idx, aux []int, lo, mid, hi int) {
	if lo >= hi {
		return
	}

	for k := lo; k <= hi; k++ {
		aux[k] = idx[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			idx[k] = aux[j]
			j++
		} else if j > hi {
			idx[k] = aux[i]
			i++
			// aux contains the positions of elements in the hypothetical partially-sorted array,
			// so we look up the values of these elements to compare them.
		} else if less(s[aux[i]], s[aux[j]]) {
			idx[k] = aux[i]
			i++
		} else {
			idx[k] = aux[j]
			j++
		}
	}
}

func less(a, b int) bool {
	return a < b
}
