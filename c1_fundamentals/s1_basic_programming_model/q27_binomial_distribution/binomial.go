package q27_binomial_distribution

import "math"

/*
Binomial(100, 50) ≅ 100-choose-50 ≅ 2^100 (very approximately). Alternatively, each recursive call
has two branches, and the "left" branch is 100 levels deep. This puts us in the vicinity of 2^100
calls.
*/

// recursive slow
func binomial(n, k int, p float64) float64 {
	if n == 0 && k == 0 {
		return 1
	}
	if n < 0 || k < 0 {
		return 0
	}

	return (1-p)*binomial(n-1, k, p) + p*binomial(n-1, k-1, p)
}

// recursive memoized
func dynamicBinomialRecursive(n, k int, p float64) float64 {
	if n < 0 || k < 0 {
		return 0
	}

	cache := make([][]float64, n+1) // n+1 because the zero slot in the cache is treated as if it were occupied by the base-case return value
	for i := range cache {
		cache[i] = make([]float64, k+1)
	}

	var binomial func(n, k int) float64
	binomial = func(n, k int) float64 {
		if n == 0 && k == 0 {
			return 1
		}
		if n < 0 || k < 0 {
			return 0
		}

		if cache[n][k] == 0 {
			cache[n][k] = ((1-p)*binomial(n-1, k) + p*binomial(n-1, k-1))
		}

		return cache[n][k]
	}

	return binomial(n, k)
}

// loop memoized
func dynamicBinomialLoop(n, k int, p float64) float64 {
	if n < 0 || k < 0 {
		return 0
	}

	cache := make([][]float64, n+1)
	for i := range cache {
		cache[i] = make([]float64, k+1)
	}

	// Base cases.
	cache[0][0] = 1
	for i := 1; i <= n; i++ {
		// The base case for each value of n > 0 (row) where k == 0 (column):
		// binomial(0, 0, p) == 1 and binomial(0, -1, p) == 0
		// => binomial(1, 0, p) == (1-p)*binomial(0, 0, p) + p*binomial(0, -1, p) == (1-p)*1 + 0 == (1-p)
		// => binomial(2, 0, p) == (1-p)*(binomial(1, 0, p)) + p*binomail(1, -1, p) == (1-p)(1-p) + 0 = (1-p)^2
		// ...
		cache[i][0] = math.Pow(1-p, float64(i))
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			// Pascal's triangle.
			cache[i][j] = p*cache[i-1][j-1] + (1-p)*cache[i-1][j]
		}
	}

	return cache[n][k]
}
