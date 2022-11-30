package topdown

import "golang.org/x/exp/constraints"

func Sort[S ~[]E, E constraints.Ordered](s S) {
	sort(s, make(S, len(s)), 0, len(s)-1)
}

func sort[S ~[]E, E constraints.Ordered](s, aux S, lo, hi int) {
	if hi <= lo { // subarray 0..1
		return
	}

	mid := lo + (hi-lo)/2
	sort(s, aux, lo, mid)
	sort(s, aux, mid+1, hi)
	merge(s, aux, lo, mid, hi)
}

func merge[S ~[]E, E constraints.Ordered](s, aux S, lo, mid, hi int) {
	i, j := lo, mid+1
	// Copy s to aux.
	for k := lo; k <= hi; k++ {
		aux[k] = s[k]
	}
	// Merge back to aux.
	for k := lo; k <= hi; k++ {
		// Must check for subarray exhaustion before checking ordering to avoid running out of
		// bounds.
		if i > mid { // LHS exhausted
			s[k] = aux[j]
			j++
		} else if j > hi { // RHS exhausted
			s[k] = aux[i]
			i++
		} else if less(aux[i], aux[j]) { // LHS value smaller
			s[k] = aux[i]
			i++
		} else { // RHS value smaller
			s[k] = aux[j]
			j++
		}
	}
}

func less[E constraints.Ordered](a, b E) bool {
	return a < b
}
