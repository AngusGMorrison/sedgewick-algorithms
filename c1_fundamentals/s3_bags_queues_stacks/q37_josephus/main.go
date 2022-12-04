package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/queue"
)

func main() {
	n := flag.Int("n", 1, "The number of people in the circle.")
	m := flag.Int("m", 1, "Eliminate every mth person.")
	flag.Parse()

	if *n <= 0 || *m <= 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	circle := newQueue(*n)
	counter := 1
	eliminations := make([]int, 0, circle.Len())
	for circle.Len() > 0 {
		person, ok := circle.Dequeue()
		if !ok {
			log.Fatalln("dequeued from empty queue")
		}
		if counter == *m {
			eliminations = append(eliminations, person)
			counter = 1
			continue
		}

		circle.Enqueue(person)
		counter++
	}

	fmt.Println(eliminations)
}

func newQueue(n int) *queue.SliceQueue[int] {
	q := queue.NewSliceQueue[int]()
	for i := 0; i < n; i++ {
		q.Enqueue(i)
	}

	return q
}
