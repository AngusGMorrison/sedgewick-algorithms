package localminimum

// Question 1.4.18 as written in the 2011 printing contains a typo. A local minimum for some index i
// is such that a[i-1] > a[i] and a[i] < a[i+1], not a[i-1] < a[i] < a[i+1].
//
// Elements at either end of the array may be considered a local minimum if they are smaller than
// their one neighbor.
func LocalMinimumRecursive(a []int) int {
	if len(a) == 0 {
		return -1
	}
	if len(a) == 1 {
		return 0
	}
	if len(a) == 2 {
		if a[0] < a[1] {
			return 0
		}
		return 1
	}

	mid := len(a) / 2
	cur := a[mid]
	if a[mid-1] > cur && cur < a[mid+1] {
		return mid
	}

	// If the element to the left of cur is smaller, the left half of the array must contain a local
	// minimum.
	if a[mid-1] < cur {
		return LocalMinimumRecursive(a[:mid])
	}
	// The element to the right of cur is smaller, so the right half of the array must contain a
	// local minimum.
	return LocalMinimumRecursive(a[mid+1:]) + len(a[:mid+1])
}

func LocalMinimumLoop(a []int) int {
	if len(a) == 1 {
		return 0
	}

	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2

		if mid == 0 { // mid is first element
			if a[0] < a[1] {
				return 0
			}
			return 1
		}

		if mid == len(a)-1 { // mid is last element
			if a[len(a)-1] < a[len(a)-2] {
				return len(a) - 1
			}

			return len(a) - 2
		}

		if a[mid-1] > a[mid] && a[mid+1] > a[mid] {
			return mid
		}

		// If the element to the left of mid is smaller, the left-hand side of the array must
		// contain a local minimum. Otherwise, the right-hand side must contain a local minimum.
		if a[mid-1] < a[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}
