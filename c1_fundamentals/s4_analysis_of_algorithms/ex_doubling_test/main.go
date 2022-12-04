package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/analysis/threesum"
)

const intMax = 10_000_000

func main() {
	for i := 250; ; i += i {
		time := timeTrial(i)
		fmt.Printf("%7d %s\n", i, time)
	}
}

func timeTrial(n int) time.Duration {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = -intMax + rand.Intn(2*intMax)
	}
	start := time.Now()
	_ = threesum.ThreeSum(data)
	return time.Since(start)
}
