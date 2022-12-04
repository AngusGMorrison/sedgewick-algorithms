package q2_threesum_overflow

import (
	"math"
	"sort"

	"golang.org/x/exp/slices"
)

func ThreeSum(a []int) int {
	var count int
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			// If the sum of the values at indices i and j overflows by anything other than 1, there
			// is no need to check any value of k, because it cannot be large enough to offset the
			// overflow and form a triple totalling zero.
			//
			// In the case that this sum overflows by 1, we know that a[i]+a[j] == math.MaxInt+1.
			// It's then possible to form a valid triple with math.MinInt.
			sumIJOverflows, ijOverflowSize := additionOverflows(a[i], a[j])
			if sumIJOverflows && ijOverflowSize != 1 {
				continue
			}

			for k := j + 1; k < len(a); k++ {
				if sumIJOverflows && a[k] == math.MinInt64 {
					count++
					continue
				}

				sumIJ := a[i] + a[j]
				if overflows, _ := additionOverflows(sumIJ, a[k]); overflows {
					continue
				}
				if sumIJ+a[k] == 0 {
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
			if overflows, overflowSize := additionOverflows(a[i], a[j]); overflows && overflowSize != 1 { // -math.MaxInt-1 == math.MinInt
				continue
			}
			if got, ok := slices.BinarySearch(a, -a[i]-a[j]); ok && got > j {
				count++
			}
		}
	}

	return count
}

// Returns true if adding the operands would cause an integer overflow, along with the size and
// direction of the overflow.
func additionOverflows(a, b int) (bool, int) {
	if b < 0 {
		return a < math.MinInt-b, a - (math.MinInt - b)
	}

	return a > math.MaxInt-b, a - (math.MaxInt - b)
}
