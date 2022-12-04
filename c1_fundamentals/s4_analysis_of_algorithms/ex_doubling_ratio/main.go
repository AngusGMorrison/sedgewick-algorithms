package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/analysis/threesum"
)

const intMax = 10_000_000

func main() {
	prev := timeTrial(125)
	for i := 250; ; i += i {
		t := timeTrial(i)
		fmt.Printf("%6d %7.1f %5.1f\n", i, t, t/prev)
		prev = t
	}
}

func timeTrial(n int) float64 {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = -intMax + rand.Intn(2*intMax)
	}
	start := time.Now()
	_ = threesum.ThreeSum(data)
	return float64(time.Since(start).Seconds())
}
