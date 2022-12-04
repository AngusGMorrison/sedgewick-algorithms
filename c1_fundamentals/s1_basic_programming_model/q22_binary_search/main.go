package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []int{10, 11, 12, 16, 18, 23, 29, 33, 48, 54, 57, 68, 77, 84, 98}
	rank(23, a)
}

func rank(key int, a []int) int {
	if len(a) == 0 {
		return -1
	}

	return rankRecursive(key, 0, len(a)-1, 0, a)
}

func rankRecursive(key, lo, hi, depth int, a []int) int {
	fmt.Printf("%slo: %d, hi: %d\n", strings.Repeat("\t", depth), lo, hi)

	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)/2
	if key < a[mid] {
		return rankRecursive(key, lo, mid-1, depth+1, a)
	} else if key > a[mid] {
		return rankRecursive(key, mid+1, hi, depth+1, a)
	} else {
		return mid
	}
}
