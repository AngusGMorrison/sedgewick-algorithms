package ex_binary_search

// BinarySearch returns the index of key in a if present, or -1 if not present. a must be sorted in
// ascending order.
func BinarySearch(key int, a []int) int {
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2 // avoids integer overflow, making it preferable to (lo+hi)/2
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
