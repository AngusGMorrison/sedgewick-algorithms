package q16_natural_mergesort

import (
	"golang.org/x/exp/constraints"
)

func MergeSort[S ~[]E, E constraints.Ordered](s S) {
	if len(s) == 0 {
		return
	}

	aux := make(S, len(s))
	for {
		for lo := 0; lo < len(s); {
			var mid int
			for mid = lo; mid < len(s)-1 && s[mid] <= s[mid+1]; mid++ {
			}
			if lo == 0 && mid == len(s)-1 { // array is sorted
				return
			}

			var hi int
			for hi = mid + 1; hi < len(s)-1 && s[hi] <= s[hi+1]; hi++ {
			}
			if hi == len(s) {
				hi = mid
			}

			merge(s, aux, lo, mid, hi)

			lo = hi + 1
		}
	}
}

func merge[S ~[]E, E constraints.Ordered](s, aux S, lo, mid, hi int) {
	if lo >= hi {
		return
	}

	for k := lo; k <= hi; k++ { // copy to aux
		aux[k] = s[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			s[k] = aux[j]
			j++
		} else if j > hi {
			s[k] = aux[i]
			i++
		} else if less(aux[i], aux[j]) {
			s[k] = aux[i]
			i++
		} else {
			s[k] = aux[j]
			j++
		}
	}
}

func less[E constraints.Ordered](a, b E) bool {
	return a < b
}
