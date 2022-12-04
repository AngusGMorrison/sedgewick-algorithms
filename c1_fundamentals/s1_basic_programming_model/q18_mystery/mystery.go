package q18_mystery

func multiply(a, b int) int {
	if b == 0 {
		return 0
	}

	if b%2 == 0 {
		return multiply(a+a, b/2)
	}

	return multiply(a+a, b/2) + a
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}

	if b%2 == 0 {
		return pow(a*a, b/2)
	}

	return pow(a*a, b/2) * a
}
