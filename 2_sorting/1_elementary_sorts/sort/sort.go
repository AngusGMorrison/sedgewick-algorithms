package sort

import "golang.org/x/exp/constraints"

func Selection[S ~[]E, E constraints.Ordered](s S) {
	for i := 0; i < len(s); i++ {
		min := i
		for j := i + 1; j < len(s); j++ {
			if less(s[j], s[min]) {
				min = j
			}
		}
		swap(s, i, min)
	}
}

func Insertion[S ~[]E, E constraints.Ordered](s S) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && less(s[j], s[j-1]); j-- {
			swap(s, j, j-1)
		}
	}
}

func Shell[S ~[]E, E constraints.Ordered](s S) {
	h := 1
	for h < len(s)/3 {
		h = h*3 + 1 // 1, 4, 13, ... (3^k-1)/2
	}
	for h >= 1 {
		for i := h; i < len(s); i++ { // note that i increases by 1, but j decreases by h
			for j := i; j >= h && less(s[j], s[j-h]); j -= h { // j >= h prevents us from decrementing past the start of the array
				swap(s, j, j-h)
			}
		}
		h /= 3
	}
}

func less[E constraints.Ordered](a, b E) bool {
	return a < b
}

func swap[S ~[]E, E any](s S, i, j int) {
	s[i], s[j] = s[j], s[i]
}
