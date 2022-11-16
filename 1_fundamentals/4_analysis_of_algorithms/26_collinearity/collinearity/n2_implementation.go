package collinearity

import (
	"math"
)

// CountCollinearN2 counts the number of collinear triples among the input points in N^2 time and
// space. This implementation is vulnerable to floating point rounding errors, which may lead
// identical lines to appear distinct incorrectly. Go's math/big package does not provide floats or
// rationals suitable for use as map keys, which this implementation depends upon, so resolving this
// issue would likely require a custom Rational type.
func CountCollinearN2(points []Point) int {
	pointSets := newCollinearPointSetsFromPoints(points)
	return countTriples(pointSets)
}

// collinearPointSets maps a line to the set of points that lie on the line.
type collinearPointSets map[line]map[Point]bool

// newCollinearPointSetsFromPoints generates the lines represented by each pair of points in O(n^2)
// time. Pairs of points that are collinear will generate the same line. We then map unique lines to
// the set of the collinear points that fall on each line.
func newCollinearPointSetsFromPoints(points []Point) collinearPointSets {
	pointSets := make(collinearPointSets)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			l := newLine(points[i], points[j])
			if _, ok := pointSets[l]; !ok {
				pointSets[l] = make(map[Point]bool)
			}
			pointSets[l][points[i]] = true
			pointSets[l][points[j]] = true
		}
	}

	return pointSets
}

// countTriples calculates the number of triples that can be formed from each set of collinear
// points. In the case where there are no collinear points, we have n^2 sets of 1 item = n^2. In
// the case where all points are collinear, we have 1 set of n items = n. In all cases, we
// calculate n-choose-3 (the number of triples formed by collinear points) in O(1) time (amortized).
// This gives an order of growth of n^2.
func countTriples(pointSets collinearPointSets) int {
	bcg := newBinomialCoefficientGenerator()
	var triples int
	for _, pointSet := range pointSets {
		if len(pointSet) < 3 {
			continue // not enough points to form a triple
		}
		triples += bcg.nCr(len(pointSet), 3)
	}
	return triples
}

// Point represents a point in the plane.
type Point struct {
	X, Y float64
}

// line represents a unique line in the plane. Any number of points may lie on a line.
type line struct {
	// gradient is required to distinguish between lines passing through the origin.
	gradient float64
	// We require both x- and y-intercepts to create distinct identities for vertical lines at
	// different x-coordinates, and vice versa.
	yIntercept float64
	xIntercept float64
}

func newLine(p1, p2 Point) line {
	m := gradient(p1, p2)
	return line{
		gradient:   m,
		yIntercept: yIntercept(p1, m),
		xIntercept: xIntercept(p1, m),
	}
}

// gradient returns the gradient of the line formed by two points.
func gradient(p1, p2 Point) float64 {
	if p2.X-p1.X == 0 {
		return math.Inf(1)
	}
	return (p2.Y - p1.Y) / (p2.X - p1.X)
}

// yIntercept returns the y-intercept of the line with the given gradient that passes through Point
// p. Since we require yIntercepts to be comparable, but NaN != NaN, the gradient of vertical lines
// is math.MaxFloat64 by convention.
func yIntercept(p Point, gradient float64) float64 {
	if math.IsInf(gradient, 0) {
		return math.MaxFloat64
	}
	return p.Y - gradient*p.X
}

// xIntercept returns the x-intercept of the line with the given gradient that passes through Point
// p. Since we require xIntercepts to be comparable, but NaN != NaN, the gradient of horizontal
// lines is math.MaxFloat64 by convention.
func xIntercept(p Point, gradient float64) float64 {
	if gradient == 0 {
		return math.MaxFloat64
	}
	return p.X - (p.Y / gradient)
}

// binomialCoefficientGenerator allows us to calculate binomial coefficients in constant time
// (amortized), by memoizing the factorials it calculates.
type binomialCoefficientGenerator struct {
	factorials map[int]int
}

func newBinomialCoefficientGenerator() *binomialCoefficientGenerator {
	return &binomialCoefficientGenerator{
		factorials: map[int]int{
			0: 1,
			1: 1,
			2: 2,
		},
	}
}

// nCr calculates the binomial coefficient n-choose-r.
func (bcg *binomialCoefficientGenerator) nCr(n, r int) int {
	if n <= 0 || r < 0 || r > n {
		return 0
	}
	return bcg.factorial(n) / (bcg.factorial(r) * (bcg.factorial(n - r)))
}

// factorial calculates n! in O(1) amortized time.
func (bcg *binomialCoefficientGenerator) factorial(n int) int {
	fac, ok := bcg.factorials[n]
	if !ok {
		fac = n * bcg.factorial(n-1)
		bcg.factorials[n] = fac
	}
	return fac
}
