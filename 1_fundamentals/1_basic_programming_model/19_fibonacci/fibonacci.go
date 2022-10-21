package fibonacci

func naive(n int) int {
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return naive(n-1) + naive(n-2)
}

func dynamicRecursive(n int) int {
	if n < 0 {
		return -1
	}

	cache := make([]int, n+1)

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		if cache[n] != 0 {
			return cache[n]
		}

		cache[n] = fib(n-1) + fib(n-2)
		return cache[n]
	}

	return fib(n)
}

func dynamicLoop(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}

	cache := make([]int, n+1)

	// Base case.
	cache[0] = 0
	cache[1] = 1

	// Sequence.
	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]
}
