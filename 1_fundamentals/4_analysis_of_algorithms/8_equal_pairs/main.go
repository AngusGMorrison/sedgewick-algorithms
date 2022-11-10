package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/exp/slices"
)

func main() {
	var ints []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		ints = append(ints, i)
	}

	pairs := countEqualPairs(ints)
	fmt.Println(pairs)
}

func countEqualPairs(a []int) int {
	slices.Sort(a)

	var count int
	for i, elem := range a {
		if _, ok := slices.BinarySearch(a[i+1:], elem); ok {
			count++
		}
	}

	return count
}
