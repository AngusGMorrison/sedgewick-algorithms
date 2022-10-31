package binary

import (
	"strconv"
	"strings"

	"github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/bqs/stack"
)

func toBinaryString(n int) string {
	s := stack.NewSliceStack[string]()
	for n > 0 {
		s.Push(strconv.Itoa(n % 2))
		n /= 2
	}

	var builder strings.Builder
	s.Each(func(s string) {
		builder.WriteString(s)
	})

	return builder.String()
}
