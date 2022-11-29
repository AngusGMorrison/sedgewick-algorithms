package merge

import "golang.org/x/exp/constraints"

func Sort[S ~[]E, E constraints.Ordered](s S) {
	sort(s, make(S, len(s)), 0, len(s)-1)
}

func sort[S ~[]E, E constraints.Ordered](s, aux S, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	sort(s, aux, lo, mid)
	sort(s, aux, mid+1, hi)
	merge(s, aux, lo, mid, hi)
}

// By copying the second half of s to aux in reverse order, we can merge aux back into s by
// maintaining pointers to opposite ends of aux and moving them closer together for each element
// merged. We cannot exhaust either "half" of aux, because the pointers will eventually arrive at
// the same element.
func merge[S ~[]E, E constraints.Ordered](s, aux S, lo, mid, hi int) {
	i, j, k := lo, mid+1, hi

	for m := lo; m < j; m++ {
		aux[m] = s[m]
	}
	for m := j; m <= hi; m++ {
		aux[m] = s[hi-(m-j)]
	}
	for m := lo; m <= hi; m++ {
		if less(aux[i], aux[k]) {
			s[m] = aux[i]
			i++
		} else {
			s[m] = aux[k]
			k--
		}
	}
}

func less[E constraints.Ordered](a, b E) bool {
	return a < b
}
