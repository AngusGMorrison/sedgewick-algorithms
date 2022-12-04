package q15_faster_sum

import "golang.org/x/exp/slices"

// a must be sorted
func TwoSumLinear(a []int) int {
	var count int
	lo := 0
	hi := len(a) - 1
	for lo < hi && a[lo] <= 0 && a[hi] >= 0 {
		if a[lo] < -a[hi] {
			lo++
		} else if a[lo] > -a[hi] {
			hi--
		} else {
			// Count the negative and positive elements that may form a valid pair, and multiply
			// them get the number of pairs that may be generated from these elements.
			var lowCount, hiCount int
			for val := a[lo]; lo < hi && a[lo] == val; lo++ {
				lowCount++
			}
			for val := a[hi]; lo <= hi && a[hi] == val; hi-- { // note lo <= hi, since lo may have advanced to hi, but hi has not yet been processed
				hiCount++
			}

			count += lowCount * hiCount
		}
	}

	return count
}

func ThreeSumQuadratic(a []int) int {
	slices.Sort(a) // sorting doesn't affect the asymptotic running time

	var count int
	for i := 0; i < len(a)-2 && a[i] <= 0; i++ {
		lo := i + 1
		hi := len(a) - 1
		for twoSum := a[i] + a[lo]; lo < hi && twoSum <= 0; twoSum = a[i] + a[lo] {
			if twoSum < -a[hi] {
				lo++
			} else if twoSum > -a[hi] {
				hi--
			} else {
				var lowCount, hiCount int
				for val := twoSum; lo < hi && a[i]+a[lo] == val; lo++ {
					lowCount++
				}
				for val := a[hi]; lo <= hi && a[hi] == val; hi-- {
					hiCount++
				}

				count += lowCount * hiCount
			}
		}
	}

	return count
}
