package q16_closest_pair

import "golang.org/x/exp/slices"

func ClosestPair(a []float64) []float64 {
	if len(a) < 2 {
		return nil
	}

	slices.Sort(a)

	first, second := a[0], a[1]
	diff := second - first
	for i := 2; i < len(a); i++ {
		if a[i]-a[i-1] < diff {
			first = a[i-1]
			second = a[i]
			diff = second - first
		}
	}

	return []float64{first, second}
}
