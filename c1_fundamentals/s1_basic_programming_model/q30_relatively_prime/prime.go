package q30_relatively_prime

func buildRelativePrimeMatrix(size uint) [][]bool {
	if size == 0 {
		return nil
	}

	m := make([][]bool, size)
	for i := range m {
		m[i] = make([]bool, size)
	}

	for i := range m {
		for j := range m[i] {
			if gcd(i, j) == 1 {
				m[i][j] = true
			}
		}
	}

	return m
}

// gcd finds the greatest common divisor of p and q using Euclid's algorithm.
func gcd(p, q int) int {
	if q == 0 {
		return p
	}

	return gcd(q, p%q)
}
