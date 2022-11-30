package bottomup

import "golang.org/x/exp/constraints"

func Sort[S ~[]E, E constraints.Ordered](s S) {
	aux := make(S, len(s))
	// Do lg n passes of pairwise merges.
	for sz := 1; sz < len(s); sz += sz { // double the size of the subarray to sort on each iteration
		// lo marks the beginning of an array pair. We ensure there is always enough room for one
		// half of a pair before entering the loop. The final subarray is of size sz only when the
		// array size is an even multiple of sz.
		for lo := 0; lo < len(s)-sz; lo += sz + sz {
			merge(s, aux, lo, lo+sz-1, min(lo+sz+sz-1, len(s)-1))
		}
	}
}

func merge[S ~[]E, E constraints.Ordered](s, aux S, lo, mid, hi int) {
	i, j := lo, mid+1
	// Copy range to aux.
	for k := lo; k <= hi; k++ {
		aux[k] = s[k]
	}
	// Merge back to s.
	for k := lo; k <= hi; k++ {
		if i > mid { // LHS exhausted
			s[k] = aux[j]
			j++
		} else if j > hi { // RHS exhausted
			s[k] = aux[i]
			i++
		} else if less(aux[i], aux[j]) { // LHS smaller
			s[k] = aux[i]
			i++
		} else { // RHS smaller
			s[k] = aux[j]
			j++
		}
	}
}

func min[E constraints.Ordered](a, b E) E {
	if a < b {
		return a
	}

	return b
}

func less[E constraints.Ordered](a, b E) bool {
	return a < b
}
