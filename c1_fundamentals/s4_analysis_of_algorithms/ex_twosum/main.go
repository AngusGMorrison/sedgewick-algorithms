package twosum

import (
	"sort"

	"golang.org/x/exp/slices"
)

func TwoSum(a []int) int {
	var count int
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 0 {
				count++
			}
		}
	}

	return count
}

func TwoSumFast(a []int) int {
	var count int
	sort.Sort(sort.IntSlice(a))
	for i := 0; i < len(a); i++ {
		if got, ok := slices.BinarySearch(a, -a[i]); ok && got > i {
			count++
		}
	}

	return count
}
