package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	for i := 10; ; i *= 10 {
		fmt.Println("Slice size: ", i)
		s := make([]int, i)
		for i := range s {
			s[i] = rng.Int()
		}
		ShellSort(s)
		fmt.Println()
	}
}

func ShellSort[S ~[]E, E constraints.Ordered](s S) {
	h := 1
	for h < len(s)/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		var compares int
		for i := h; i < len(s); i++ {
			for j := i; j >= h; j -= h {
				compares++
				if less(s[j], s[j-h]) {
					swap(s, j, j-h)
				}
			}
		}
		fmt.Printf("h=%d; ratio=%.3f\n", h, float64(compares)/float64(len(s)))
		h /= 3
	}
}

func less[E constraints.Ordered](a, b E) bool {
	return a < b
}

func swap[S ~[]E, E any](s S, i, j int) {
	{
		s[i], s[j] = s[j], s[i]
	}
}
