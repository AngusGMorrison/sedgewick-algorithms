package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
)

func main() {
	n := flag.Int("n", 2, "The number of points to plot. Minimum 2.")
	flag.Parse()

	if *n < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	points := generatePoints(*n)
	dist := distanceBetweenClosestPair(points)
	fmt.Println(dist)
}

// generatePoints generates n random points in the unit square.
func generatePoints(n int) []gg.Point {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	points := make([]gg.Point, n)
	for i := 0; i < n; i++ {
		points[i] = gg.Point{
			X: rng.Float64(),
			Y: rng.Float64(),
		}
	}

	return points
}

func distanceBetweenClosestPair(points []gg.Point) float64 {
	min := math.Inf(1)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if dist := points[i].Distance(points[j]); dist < min {
				min = dist
			}
		}
	}

	return min
}
