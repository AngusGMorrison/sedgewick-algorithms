package ex_gcd

// GCD returns the greatest common divisor of p and q.
func GCD(p int, q int) int {
	if q == 0 {
		return p
	}

	r := p % q
	return GCD(q, r)
}
