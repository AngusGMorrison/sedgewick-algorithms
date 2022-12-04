package q17_farthest_pair

import "math"

func FarthestPair(a []float64) []float64 {
	if len(a) < 2 {
		return nil
	}

	min := math.Inf(1)
	max := math.Inf(-1)
	for _, n := range a {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return []float64{min, max}
}
