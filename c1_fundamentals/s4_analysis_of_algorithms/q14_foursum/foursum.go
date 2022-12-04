package q14_foursum

import "golang.org/x/exp/slices"

func FourSum(a []int) int {
	slices.Sort(a)

	var count int
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			for k := j + 1; k < len(a); k++ {
				if _, ok := slices.BinarySearch(a[k+1:], -a[i]-a[j]-a[k]); ok {
					count++
				}
			}
		}
	}

	return count
}
