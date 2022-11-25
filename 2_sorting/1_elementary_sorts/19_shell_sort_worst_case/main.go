package main

// This script demonstrates that the worst case number of compares for shell short for an array of
// 100 elements containing the numbers 1 through 100 occurs for certain randomly ordered arrays, by
// a narrow margin over reversed arrays.
/*
	Array           Compares
    Sorted          342
    Reversed        601
    Random          604 [40 48 84 97 89 41 53 99 70 76 64 91 63 15 61 52 87 60 83 67 7 71 69 72 75 42 86 82 45 95 66 93 73 47 65 55 38 79 28 94 22 24 44 37 77 23 59 34 96 21 51 100 80 36 58 10 54 39 9 31 12 3 16 46 90 33 85 20 35 88 2 56 11 78 49 98 13 30 29 17 8 32 6 18 62 57 25 1 50 27 14 26 19 74 81 43 5 68 4 92]
*/

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"
)

const randomTrials = 1000

func main() {
	sorted := make([]int, 100)
	for i := range sorted {
		sorted[i] = i + 1
	}
	sortedCompares := shellSortWithCompares(sorted)
	if !sort.IntsAreSorted(sorted) {
		log.Fatalf("want sorted array got\n\t%v\n", sorted)
	}

	reversed := make([]int, 100)
	for i := range sorted {
		reversed[i] = 100 - i
	}
	reversedCompares := shellSortWithCompares(reversed)

	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	maxRandomCompares := -1
	var maxCompareSlice []int
	for i := 0; i < randomTrials; i++ {
		random := make([]int, 100)
		for i := range random {
			random[i] = i + 1
		}
		rng.Shuffle(len(random), func(i, j int) {
			random[i], random[j] = random[j], random[i]
		})
		unsortedRandom := make([]int, 100)
		copy(unsortedRandom, random)
		compares := shellSortWithCompares(random)
		if compares > maxRandomCompares {
			maxRandomCompares = compares
			maxCompareSlice = unsortedRandom
		}
	}

	fmt.Printf("%s\t\t%d\n", "Sorted", sortedCompares)
	fmt.Printf("%s\t%d\n", "Reversed", reversedCompares)
	fmt.Printf("%s\t\t%d %v\n", "Random", maxRandomCompares, maxCompareSlice)
}

// shellSortWithCompares shell sorts a and returns the number of compares required to do so.
func shellSortWithCompares(a []int) int {
	h := 1
	for h < len(a)/3 {
		h = h*3 + 1
	}

	var compares int
	for h >= 1 {
		for i := h; i < len(a); i++ {
			compares++
			j := i
			elem := a[j]
			for ; j >= h && a[j] < a[j-h]; j -= h {
				compares++
				a[j] = a[j-h]
			}
			a[j] = elem
		}
		h /= 3
	}

	return compares
}
