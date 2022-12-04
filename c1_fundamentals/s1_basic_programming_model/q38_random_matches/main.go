package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	min = 1e6
	max = 1e7
)

func main() {
	runs := flag.Int("runs", 0, "The number of trials to run.")
	flag.Parse()

	sizes := []int{1e3, 1e4, 1e5, 1e6}
	t := newTrial(*runs, sizes)
	t.run()
	t.print()
}

type trial struct {
	runs    int
	rng     *rand.Rand
	sizes   []int
	results []float64
}

func newTrial(runs int, sizes []int) *trial {
	src := rand.NewSource(time.Now().UnixNano())
	return &trial{
		runs:    runs,
		rng:     rand.New(src),
		sizes:   sizes,
		results: make([]float64, len(sizes)),
	}
}

func (t *trial) run() {
	for i, s := range t.sizes {
		t.results[i] = t.runAllForSize(s)
	}
}

// runAllForSize runs t.runs trials on arrays of the given size and returns the average number of
// intersections across all runs.
func (t *trial) runAllForSize(size int) float64 {
	var isx int
	for i := 0; i < t.runs; i++ {
		isx += t.runOnceForSize(size)
	}

	return float64(isx) / float64(t.runs)
}

// runOnceForSize generates two arrays of size n populated with random numbers in the range [min,
// max). It returns the number of values that appear in both arrays.
func (t *trial) runOnceForSize(n int) int {
	a1 := t.randomArray(n)
	a2 := t.randomArray(n)

	return intersections(a1, a2)
}

// randomArray generates an array of the given size populated with random numbers in the range
// [min,max).
func (t *trial) randomArray(size int) []int {
	a := make([]int, size)
	for i := range a {
		a[i] = min + t.rng.Intn(max-min)
	}

	return a
}

func (t *trial) print() {
	fmt.Printf("Size\tResult\n")
	for i := range t.results {
		fmt.Printf("%.0e:\t%.3f\n", float64(t.sizes[i]), t.results[i])
	}
}

// intersections returns the number of elements in a1 that also appear in a2.
func intersections(a1, a2 []int) int {
	var count int
	for _, k := range a1 {
		if binarySearch(k, a2) > -1 {
			count++
		}
	}

	return count
}

// binarySearch returns the index of key in a, or -1 if key is not present.
func binarySearch(key int, a []int) int {
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
