package ex_rabin_karp

import (
	"math"
)

const (
	q = 683303        // some large prime to ensure all bits of key are used in hash
	r = math.MaxUint8 // max value of unsigned byte
)

func IndexOf(pat, text string) int {
	n := len(pat)
	if n > len(text) {
		return -1
	}

	// Precalculate the hash of the pattern and the first n bytes of text.
	patHash := hash(pat, n)
	textHash := hash(text[:n], n)
	if patHash == textHash {
		return 0
	}

	maxIdx := len(text) - n
	highDigitFactor := uint64(math.Pow(r, float64(n-1))) % q // factor required to remove the leading digit from textHash
	for i := 1; i <= maxIdx; i++ {
		leading := uint64(text[i-1]) * highDigitFactor % q // digit to remove from textHash
		next := uint64(text[i-1+n])                        // digit to add to textHash
		// Remove the leading digit.
		textHash = (textHash - leading + q) % q // extra q is added to ensure that the intermediate value remains positive so the remainder operation works correctly
		// Add the next digit.
		textHash = (textHash*r + next) % q
		if textHash == patHash {
			return i
		}
	}

	return -1
}

// Hash the first n bytes of s.
func hash(s string, n int) uint64 {
	var h uint64
	for i := 0; i < n; i++ {
		h = (r*h + uint64(s[i])) % q
	}
	return h
}

// Prime that serves as hash base, from the Go standard library.
const baseQ = 16777619

// IndexOfAlternative uses the fact that all unsigned integer arithmetic in Go is modular
// arithmetic. Instead of representing a string as a base-r number mod q, we represent it as a
// base-q number mod math.MaxUint32 using integer overflow. This improves performance by avoiding
// the division operation.
func IndexOfAlternative(pat, text string) int {
	n := len(pat)
	if n > len(text) {
		return -1
	}

	patHash := hashAlternative(pat)
	textHash := hashAlternative(text[:n])
	if patHash == textHash {
		return 0
	}

	factor := uint32(math.Pow(baseQ, float64(n-1)))
	for i := n; i < len(text); i++ {
		textHash -= uint32(text[i-n]) * factor
		textHash *= baseQ
		textHash += uint32(text[i])
		if textHash == patHash {
			return i - n + 1
		}
	}

	return -1
}

func hashAlternative(s string) uint32 {
	var h uint32
	for i := range s {
		h = h*baseQ + uint32(s[i])
	}
	return h
}
