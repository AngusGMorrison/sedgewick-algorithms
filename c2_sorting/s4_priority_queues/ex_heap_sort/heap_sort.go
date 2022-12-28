package ex_heap_sort

import "golang.org/x/exp/constraints"

// HeapSort sorts s[0] through s[len(s)-1] using two loops. The first constructs a max heap. The
// second exchanges the largest element s[0] with the final unordered element and repairs the heap
// until the heap has been fully replaced by the sorted array.
func HeapSort[S ~[]E, E constraints.Ordered](s S) {
	if len(s) <= 1 {
		return
	}
	// Starting from the lowest node with children, construct a max heap by progressively turning
	// each branch of the tree into its own sub-heap until the root is reached.
	for k := (len(s) - 1) / 2; k >= 0; k-- {
		sink(s, k, len(s)-1)
	}

	// Populate the array in reverse order, swapping largest item from the heap (the root) with the
	// item at the end of the array, then sinking the new root into the correct position.
	for k := len(s) - 1; k > 0; {
		s[0], s[k] = s[k], s[0]
		k--
		sink(s, 0, k)
	}
}

func sink[S ~[]E, E constraints.Ordered](s S, lo, hi int) {
	for (lo<<1)+1 <= hi { // while k has children
		j := (lo << 1) + 1 // select the greatest child
		if j < hi && s[j] < s[j+1] {
			j++
		}
		if !(s[lo] < s[j]) { // swap k with its greatest child if k is smaller
			break
		}
		s[lo], s[j] = s[j], s[lo]
		lo = j // proceed to the next level
	}
}
