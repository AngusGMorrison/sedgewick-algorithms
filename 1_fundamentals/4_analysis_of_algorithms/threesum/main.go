package threesum

import (
	"sort"

	"golang.org/x/exp/slices"
)

func ThreeSum(a []int) int {
	var count int
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			for k := j + 1; k < len(a); k++ {
				if a[i]+a[j]+a[k] == 0 {
					count++
				}
			}
		}
	}

	return count
}

func ThreeSumFast(a []int) int {
	sort.Sort(sort.IntSlice(a))
	var count int
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if got, ok := slices.BinarySearch(a, -a[i]-a[j]); ok && got > j {
				count++
			}
		}
	}

	return count
}
