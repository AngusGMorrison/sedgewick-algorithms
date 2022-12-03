package sublinear

import (
	"math"

	"golang.org/x/exp/constraints"
)

// blockSize is the length of sorted subsequences (blocks) we must produce via standard merges
// before selection sorting the blocks comprising the input slice using their first element as the
// sort key. This ensures that each block is internally ordered, which is a requirement for merging
// with other blocks. Must be a power of 2.
const blockSize = 4

func mergeSort(s []int) {
	// fmt.Printf("input: %v\n", s)

	// In order to merge blocks of elements, those blocks must be in sorted order. We use insertion
	// sort to sort each of the blocks comprising s, since insertion sort may outperform merge sort
	// for n <= 15.
	for i := 0; i < len(s); i += blockSize {
		insertionSort(s, i, min(i+blockSize-1, len(s)-1))
	}

	// fmt.Printf("after insertion sort: %v\n", s)

	// Once blocks are sorted, we can merge them together. First, we sort each *whole block* into
	// position using its first element as the sort key. I.e. block2 < block1 iff block2[0] <
	// block1[0]. Then, we merge block1 -> block2, block2 -> block3, etc. This requires more than
	// one pass, since, for example, elements that should have their final position in block2, but
	// are currently in block4, will only reach block3 when block3 is merged with block4 on the
	// first pass. The number of times we must merge sorted subarrays to produce a sorted array is
	// no more than lg len(s), provided the length of the subarrays is a power of 2.
	aux := make([]int, blockSize)
	iterations := blockwiseIterations(len(s))
	for i := 0; i < iterations; i++ {
		blockwiseMerge(s, aux)
	}
}

// blockwiseIterations returns
func blockwiseIterations(len int) int {
	return len/blockSize - 1
}

func insertionSort(s []int, lo, hi int) {
	for i := lo + 1; i <= hi; i++ {
		elem := s[i]
		var j int
		for j = i; j > lo && less(s[j], s[j-1]); j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
		s[j] = elem
	}
}

// stdMerge is used to sort subarrays with len <= blockSize, which aux is guaranteed to be able
// to hold.
func stdMerge(s, aux []int, lo, mid, hi int) {
	if lo >= hi {
		return
	}

	// Copy to aux.
	for i, j := lo, 0; i <= hi; i, j = i+1, j+1 {
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
}

func blockwiseMerge(s, aux []int) {
	if len(s) <= blockSize { // already sorted
		return
	}

	// defer fmt.Printf("%v\n", s)

	// fmt.Printf("lo: %d, mid: %d, hi: %d, hi-lo+1 < blockSize: %t\n", lo, mid, hi, hi-lo+1 < blockSize)
	// fmt.Printf("input: %v\n", s)

	// If the input range covers more than one block, perform a blockwise merge.
	excessKey := math.MaxInt
	excess := len(s) % blockSize
	if excess != 0 {
		excessKey = s[len(s)-excess]
	}
	selectionSortBlocks(s)
	// fmt.Printf("selection sorted: %v\n", s)

	// Find the new location of excess key.
	excessIdx := -1
	for i, key := range s {
		if key == excessKey {
			excessIdx = i
			break
		}
	}

	// While there is at least one whole block, merge that block with the following block.
	for i := 0; i < len(s)-blockSize; {
		curBlockSize := blockSize
		if i == 0 && i == excessIdx { // values are required to be unique othewise earlier duplicates are in previous blocks and aren't eligible for later merges containing duplicates, e.g. []int{62, 89, 28, 74, 11, 45, 37, 6, 95, 66, 28}
			curBlockSize = excess
		}
		// Copy current block to aux.
		for j := 0; j < curBlockSize; j++ {
			aux[j] = s[i+j]
		}
		// fmt.Printf("aux: %v\n", aux)

		// Merge aux and the next block in s.
		auxIdx := 0
		nextBlockIdx := i + curBlockSize
		nextBlockSize := blockSize
		if i == 0 && nextBlockIdx == excessIdx {
			nextBlockSize = excess
		}
		nextBlockHi := min(nextBlockIdx+nextBlockSize-1, len(s)-1)
		for sortedIdx := i; sortedIdx <= nextBlockHi; sortedIdx++ {
			if auxIdx >= curBlockSize { // LHS exhausted
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
		// fmt.Printf("after i=%d: %v\n", i, s)
		i += curBlockSize
	}
}

func selectionSortBlocks[S ~[]E, E constraints.Ordered](s S) {
	for i := 0; i < len(s); {
		// Find block with smallest first element.
		min := i
		for j := i + blockSize; j < len(s); j += blockSize {
			if s[j] < s[min] {
				min = j
			}
		}

		// Swap block into position.
		if min+blockSize-1 >= len(s) { // min block is the at the end of the slice and has < sz elements
			// Since the min block has fewer elements than the block it is being swapped with,
			// elements from the tail of the longer block are left behind after swapping. These
			// excess blocks must be bubbled up to the end of the array to rejoin their original
			// block.

			// Swap all the elements from the partial block with elements from the block at index i.
			nSwappable := len(s) - min
			var j int
			for ; j < nSwappable; j++ {
				s[i+j], s[min+j] = s[min+j], s[i+j]
			}

			// Swap the excess elements from the original block at index i to the end of the slice.
			// This is easiest to do in reverse order, starting from the end of the excess elements,
			// which is swapped until it reaches the end of the slice. The next of the excess
			// element then moves one space fewer, and so on.
			excess := blockSize - nSwappable
			for k := excess - 1; k >= 0; k-- { // for each excess element
				for m := i + j + k; m < len(s)-excess+k; m++ { // bubble it to its final position
					s[m], s[m+1] = s[m+1], s[m]
				}
			}

			// Increment i by the number of elements in the min block to preserve block boundaries.
			i += nSwappable
		} else {
			// Swap two whole blocks.
			for j := 0; j < blockSize; j++ {
				s[i+j], s[min+j] = s[min+j], s[i+j]
			}
			i += blockSize
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

// input: [81 87 47 59 81 18 25 40 56 0 94 11 62 89 28 74]
// after insertion sort: [47 59 81 87 18 25 40 81 0 11 56 94 28 62 74 89]
// selection sorted: [0 11 56 94 18 25 40 81 28 62 74 89 47 59 81 87]
// aux: [0 11 56 94]
// after i=0: [0 11 18 25 40 56 81 94 28 62 74 89 47 59 81 87]
// aux: [40 56 81 94]
// after i=4: [0 11 18 25 28 40 56 62 74 81 89 94 47 59 81 87]
// aux: [74 81 89 94]
// after i=8: [0 11 18 25 28 40 56 62 47 59 74 81 81 87 89 94]
//     /Users/amorrison/Documents/github.com/angusgmorrison/sedgewick_algorithms/2_sorting/2_mergesort/12_sublinear_extra_space/sublinear_test.go:53: want sorted slice, got
//         	[0 11 18 25 28 40 56 62 47 59 74 81 81 87 89 94]
