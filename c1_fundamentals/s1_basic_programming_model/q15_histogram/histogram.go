package q15_histogram

func histogram(a []int, m int) []int {
	if len(a) == 0 || m == 0 {
		return nil
	}

	result := make([]int, m)
	for _, n := range a {
		if n >= m {
			return nil
		}

		result[n]++
	}

	return result
}
