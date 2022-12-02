package sublinear

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

const blockSize = 5

func mergeSort[S ~[]E, E constraints.Ordered](s S) {
	aux := make(S, blockSize)
	for sz := 1; sz < len(s); sz += sz {
		for lo := 0; lo < len(s)-sz; lo += sz + sz {
			merge(s, aux, lo, lo+sz-1, min(lo+sz+sz-1, len(s)-1))
		}
	}
}

func merge[S ~[]E, E constraints.Ordered](s, aux S, lo, mid, hi int) {
	if lo >= hi {
		return
	}

	// defer fmt.Printf("%v\n", s[lo:hi+1])

	fmt.Printf("lo: %d, mid: %d, hi: %d, hi-lo+1 < blockSize: %t\n", lo, mid, hi, hi-lo+1 < blockSize)
	fmt.Printf("input: %v\n", s[lo:hi+1])

	// If the input range covers one or fewer blocks, perform a standard merge.
	if hi-lo+1 <= blockSize {
		for i, j := lo, 0; i <= hi; i, j = i+1, j+1 { // copy full range to aux
			aux[j] = s[i]
		}

		auxMid, auxHi := mid-lo, hi-lo
		i, j := 0, auxMid+1
		for k := lo; k <= hi; k++ {
			if i > auxMid { // LHS exhausted
				s[k] = aux[j]
				j++
			} else if j > auxHi { // RHS exhausted
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
		return
	}

	// If the input range covers more than one block, perform a blockwise merge.
	selectionSortBlocks(s, lo, hi, blockSize)
	fmt.Printf("selection sorted: %v\n", s[lo:hi+1])
	// While there is at least one whole block, merge that block with the following block.
	for i := lo; i <= hi-blockSize+1; i += blockSize {
		// Copy current block to aux.
		for j := 0; j < blockSize; j++ {
			aux[j] = s[i+j]
		}

		// Merge aux and the next block in s.
		auxIdx := 0
		nextBlockIdx := i + blockSize
		nextBlockHi := min(nextBlockIdx+blockSize, hi)
		for sortedIdx := i; sortedIdx <= nextBlockHi; sortedIdx++ {
			if auxIdx >= blockSize { // LHS exhausted
				s[sortedIdx] = s[nextBlockIdx]
				nextBlockIdx++
			} else if nextBlockIdx > nextBlockHi { // RHS exhausted
				s[sortedIdx] = aux[auxIdx]
				auxIdx++
			} else if less(aux[auxIdx], s[nextBlockIdx]) {
				s[sortedIdx] = aux[auxIdx]
				auxIdx++
			} else {
				s[sortedIdx] = s[nextBlockIdx]
				nextBlockIdx++
			}
		}
	}
}

func selectionSortBlocks[S ~[]E, E constraints.Ordered](s S, lo, hi, blockSz int) {
	for i := lo; i <= hi; {
		// Find block with smallest first element.
		min := i
		for j := i + blockSz; j <= hi; j += blockSz {
			if s[j] < s[min] {
				min = j
			}
		}

		// Swap block into position.
		if min+blockSz-1 > hi { // min block is the at the end of the slice and has < sz elements
			// Since the min block has fewer elements than the block it is being swapped with,
			// elements from the tail of the longer block are left behind after swapping. These
			// excess blocks must be bubbled up to the end of the array to rejoin their original
			// block.

			// Swap all the elements from the partial block with elements from the block at index i.
			nSwappable := hi - min + 1
			var j int
			for ; j < nSwappable; j++ {
				s[i+j], s[min+j] = s[min+j], s[i+j]
			}

			// Swap the excess elements from the original block at index i to the end of the slice.
			// This is easiest to do in reverse order, starting from the end of the excess elements,
			// which is swapped until it reaches the end of the slice. The next of the excess
			// element then moves one space fewer, and so on.
			excess := blockSz - nSwappable
			for k := excess - 1; k >= 0; k-- { // for each excess element
				for m := i + j + k; m <= hi-excess+k; m++ { // bubble it to its final position
					s[m], s[m+1] = s[m+1], s[m]
				}
			}

			// Increment i by the number of elements in the min block to preserve block boundaries.
			i += nSwappable
		} else {
			// Swap two whole blocks.
			for j := 0; j < blockSz; j++ {
				s[i+j], s[min+j] = s[min+j], s[i+j]
			}
			i += blockSz
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
