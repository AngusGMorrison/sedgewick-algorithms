package bsearch

func search(key int, a []int) int {
	if len(a) == 0 {
		return -1
	}

	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

// countSmallerThan returns the number of elements smaller than key in the sorted slice a.
func countSmallerThan(key int, a []int) int {
	idx := search(key, a)
	if idx == -1 {
		return -1
	}

	nSmallerElems := idx
	for i := idx - 1; i >= 0 && a[i] == a[idx]; i-- {
		nSmallerElems--
	}

	return nSmallerElems
}

func countGreaterThan(key int, a []int) int {
	idx := search(key, a)
	if idx == -1 {
		return -1
	}

	var i int
	for i = idx + 1; i < len(a) && a[i] == a[idx]; i++ {
	}

	return len(a) - i
}

func count(key int, a []int) int {
	nSmaller := countSmallerThan(key, a)
	if nSmaller == -1 {
		return 0
	}

	nGreater := countGreaterThan(key, a)

	return len(a) - (nSmaller + nGreater)
}
