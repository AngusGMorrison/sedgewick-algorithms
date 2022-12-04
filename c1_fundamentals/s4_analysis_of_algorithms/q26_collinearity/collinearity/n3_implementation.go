package collinearity

// countCollinearN3 is a naive implementation of an algorithm to count the number of collinear
// triples among points. It runs in O(n^3) time and O(1) space.
func countCollinearN3(points []Point) int {
	var count int
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			gradientIJ := gradient(points[i], points[j])
			for k := j + 1; k < len(points); k++ {
				if gradient(points[j], points[k]) == gradientIJ {
					count++
				}
			}
		}
	}

	return count
}
