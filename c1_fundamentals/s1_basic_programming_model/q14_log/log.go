package q14_log

// lg returns the largest int not larger than the base-2 logarithm of n. If lg is undefined for n,
// -1 is returned.
func lg(n int) int {
	if n <= 0 {
		return -1
	}

	var result int
	for n > 1 { // 1 == 2^0, so n == 1 is our base case
		n >>= 1
		result++
	}

	return result
}
