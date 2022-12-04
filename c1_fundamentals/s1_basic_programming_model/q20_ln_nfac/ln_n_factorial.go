package q20_ln_nfac

import "math"

// ln(n!) == ln(1) + ln(2) + ... + ln(n-1) + ln(n)
func lnFactorial(n int) float64 {
	if n == 0 {
		return 0
	}

	return math.Log(float64(n)) + lnFactorial(n-1)
}
