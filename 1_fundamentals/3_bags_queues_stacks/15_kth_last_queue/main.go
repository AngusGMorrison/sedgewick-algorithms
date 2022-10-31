package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/bqs/queue"
)

func main() {
	k := flag.Uint("k", 1, "Print the kth-last element.")
	flag.Parse()
	if *k < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	q := enqueueStrings(os.Stdin)
	elem := kthLast[string](q, *k)
	fmt.Println(elem)
}

func enqueueStrings(r io.Reader) *queue.SliceQueue[string] {
	q := queue.NewSliceQueue[string]()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		q.Enqueue(scanner.Text())
	}
	return q
}

func kthLast[E any](q *queue.SliceQueue[E], k uint) E {
	var kthLast E
	var i int
	target := q.Len() - int(k)
	q.Each(func(elem E) {
		if i == target {
			kthLast = elem
		}
		i++
	})

	return kthLast
}
