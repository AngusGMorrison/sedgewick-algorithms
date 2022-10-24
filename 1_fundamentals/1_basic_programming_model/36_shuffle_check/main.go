package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	m := flag.Int("m", 0, "The size of the array to be shuffled.")
	n := flag.Int("n", 0, "The number of times to shuffle the array.")
	flag.Parse()

	shuffleTest(*m, *n)
}

func shuffleTest(m, n int) {
	t := newTable(m)
	for i := 0; i < n; i++ {
		a := newArr(m)
		shuffle(a)
		for j, k := range a {
			t[k][j]++
		}
	}

	t.print()
}

type table [][]int

func (t table) print() {
	var builder strings.Builder
	// Write column headings.
	for i := range t[0] {
		fmt.Fprintf(&builder, "\t%d", i)
	}
	builder.WriteByte('\n')

	// Write rows.
	for i, row := range t {
		fmt.Fprintf(&builder, "%d: ", i)
		for _, v := range row {
			fmt.Fprintf(&builder, "\t%d", v)
		}
		builder.WriteByte('\n')
	}

	fmt.Println(builder.String())
}

func newTable(m int) table {
	t := make(table, m)
	for i := range t {
		t[i] = make([]int, m)
	}

	return t
}

func newArr(m int) []int {
	a := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = i
	}

	return a
}

func shuffle(a []int) {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	for i := 0; i < len(a); i++ {
		// Exchange a[i] with a random element in a[i:len(a)].
		r := i + rng.Intn(len(a)-i)
		a[i], a[r] = a[r], a[i]
	}
}
