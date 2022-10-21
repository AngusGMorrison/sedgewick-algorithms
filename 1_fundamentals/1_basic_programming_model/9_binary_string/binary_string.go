package bstring

import (
	"math"
	"strings"
)

func toLittleEndianString(n int) string {
	if n == 0 {
		return "0"
	}

	var builder strings.Builder
	for n > 0 {
		if 1&n == 0 {
			builder.WriteByte('0')
		} else {
			builder.WriteByte('1')
		}

		n >>= 1
	}

	return builder.String()
}

func toBigEndianString(n int) string {
	if n == 0 {
		return "0"
	}

	nBits := int(math.Log2(float64(n))) + 1
	binary := make([]byte, nBits)
	for bit := nBits - 1; bit >= 0; bit-- {
		if 1&n == 0 {
			binary[bit] = '0'
		} else {
			binary[bit] = '1'
		}

		n >>= 1
	}

	return string(binary)
}
