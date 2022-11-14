package fibonacci

func FibonacciSearch(a []int, key int) int {
	if len(a) == 0 {
		return -1
	}

	// Find the first Fibonacci number greater than or equal to len(a).
	var fibLo int
	fibMid := 1
	fibHi := 1
	for fibHi < len(a)-1 {
		fibMid, fibHi = fibHi, fibHi+fibMid
	}

	// After each loop iteration, decrement the Fibonacci sequence.
	for offset := 0; fibHi > 0; fibHi, fibMid = fibMid, fibLo {
		fibLo = fibHi - fibMid             // fibLo shrinks exponentially with each loop iteration
		idx := min(offset+fibLo, len(a)-1) // ensure we don't run out of bounds
		if a[idx] == key {
			return idx
		}

		// If the key is greater than the current element, increase the offset. If it is less than
		// the current element, we don't need to do anything in the body of the loop, because fibHi
		// and fibMid will be decremented at the end of the loop, and these determine the upper
		// bound of our search range.
		if a[idx] < key {
			offset += fibLo
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
