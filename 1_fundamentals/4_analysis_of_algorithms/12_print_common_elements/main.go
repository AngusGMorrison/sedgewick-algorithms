package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	a := []int{2, 4, 6, 8, 10}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printCommon(a, b)
}

func printCommon(a, b []int) {
	slices.Sort(a)
	slices.Sort(b)

	var ia, ib int
	for ia < len(a) && ib < len(b) {
		if a[ia] < b[ib] {
			ia++
		} else if a[ia] > b[ib] {
			ib++
		} else {
			fmt.Println(a[ia])
			ia++
			ib++
		}
	}
}
