package collinearity

import "math"

// threeSum returns the number of triples in a that sum to zero using the observation that a + b + c
// = 0 if and only if the points (a, a^3), (b, b^3) and (c, c^3) are collinear.
func threeSum(a []int) int {
	points := make([]Point, len(a))
	for i, elem := range a {
		points[i] = Point{X: float64(elem), Y: math.Pow(float64(elem), 3)}
	}
	return CountCollinearN2(points)
}
