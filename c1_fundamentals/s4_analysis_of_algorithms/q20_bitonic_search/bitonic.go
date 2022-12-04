package q20_bitonic_search

func BitonicSearch(a []int, key int) int {
	if len(a) == 0 {
		return -1
	}
	if len(a) == 1 {
		if a[0] == key {
			return 0
		}
		return -1
	}

	tp := transitionPoint(a)

	// Search the increasing part of the array.
	lo := 0
	hi := tp
	for lo <= hi {
		mid := lo + (hi-lo)/2
		midVal := a[mid]
		if midVal < key {
			lo = mid + 1
		} else if midVal > key {
			hi = mid - 1
		} else {
			return mid
		}
	}

	// Search the decreasing part of the array.
	lo = tp
	hi = len(a) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		midVal := a[mid]
		if midVal < key {
			hi = mid - 1
		} else if midVal > key {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

// transitionPoint returns the index of the maximum value of a bitonic array in O(logn) time.
func transitionPoint(bitonic []int) int {
	if len(bitonic) < 2 { // can't form an increasing sequence followed by a decreasing sequence with fewer than 2 numbers
		return -1
	}

	lo := 0
	hi := len(bitonic) - 1
	boundary := -1
	for boundary < 0 {
		mid := lo + (hi-lo)/2
		midVal := bitonic[mid]
		if mid == 0 && midVal > bitonic[mid+1] { // mid is array start
			boundary = mid
		} else if mid == len(bitonic)-1 && midVal > bitonic[mid-1] { // mid is array end
			boundary = mid
		} else if midVal > bitonic[mid-1] && midVal > bitonic[mid+1] { // mid is array maximum
			boundary = mid
		} else if midVal > bitonic[mid-1] && midVal < bitonic[mid+1] { // mid is in increasing sequence
			lo = mid + 1
		} else { // mid is in decreasing sequence
			hi = mid - 1
		}
	}

	return boundary
}
