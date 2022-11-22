package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/union_find/unionfind"
)

func main() {
	n := flag.Int("n", 9, "[0, n) is the range of the union-find.")
	flag.Parse()
	fmt.Println(count(*n))
}

func count(n int) int {
	wqu := unionfind.NewWeightedQuickUnion(n)
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	var connections int
	for wqu.Count() > 1 {
		p, q := rng.Intn(n), rng.Intn(n)
		if wqu.Connected(p, q) {
			continue
		}

		wqu.Union(p, q)
		connections++
	}

	return connections
}
