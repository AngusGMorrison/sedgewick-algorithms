package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/analysis/26_collinearity/collinearity"
)

const intMax = 1000

func main() {
	prev := timeTrial(125)
	for i := 250; ; i += i {
		t := timeTrial(i)
		fmt.Printf("%6d %7.1f %5.1f\n", i, t, t/prev)
		prev = t
	}
}

func timeTrial(n int) float64 {
	data := make([]collinearity.Point, n)
	for i := 0; i < n; i++ {
		data[i] = collinearity.Point{
			X: float64(-intMax + rand.Intn(2*intMax)),
			Y: float64(-intMax + rand.Intn(2*intMax)),
		}
	}
	start := time.Now()
	_ = collinearity.CountCollinearN2(data)
	return float64(time.Since(start).Seconds())
}
