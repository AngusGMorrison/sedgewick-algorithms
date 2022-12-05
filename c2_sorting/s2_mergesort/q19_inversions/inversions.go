package q19_inversions

import "golang.org/x/exp/constraints"

// Inversions sorts a copy of s, incrementing a counter each time two elements are swapped.
func Inversions[S ~[]E, E constraints.Ordered](s S) int {
	// Copy the input array to avoid mutating the original.
	cp := make(S, len(s))
	copy(cp, s)

	return sortAndCount(cp)
}

func sortAndCount[S ~[]E, E constraints.Ordered](s S) int {
	var inversions int
	aux := make(S, len(s))
	for sz := 1; sz < len(s); sz += sz {
		for j := 0; j < len(s)-sz; j += sz + sz {
			inversions += mergeAndCount(s, aux, j, j+sz-1, min(j+sz+sz-1, len(s)-1))
		}
	}
	return inversions
}

func mergeAndCount[S ~[]E, E constraints.Ordered](s, aux S, lo, mid, hi int) int {
	if lo >= hi {
		return 0
	}

	for k := lo; k <= hi; k++ {
		aux[k] = s[k]
	}

	var inversions int
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			s[k] = aux[j]
			j++
		} else if j > hi {
			s[k] = aux[i]
			i++
		} else if less(s[i], s[j]) {
			s[k] = s[i]
			i++
		} else {
			s[k] = s[j]
			j++
			inversions++
		}
	}

	return inversions
}

func less[E constraints.Ordered](a, b E) bool {
	return a < b
}

func min[E constraints.Ordered](a, b E) E {
	if a < b {
		return a
	}
	return b
}
