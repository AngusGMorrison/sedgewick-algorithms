package q10_least_binary_search

import (
	"golang.org/x/exp/constraints"
)

func BinarySearch[E constraints.Ordered](a []E, key E) int {
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := int((uint(lo) + uint(hi)) >> 1)
		val := a[mid]
		if key < val {
			hi = mid - 1
		} else if key > val {
			lo = mid + 1
		} else {
			if mid == 0 {
				return mid
			}
			if a[mid-1] == key {
				hi = mid - 1
				continue
			}

			return mid
		}
	}

	return -1
}
