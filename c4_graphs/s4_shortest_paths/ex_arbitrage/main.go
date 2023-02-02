package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_bellman_ford"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdg"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	if !scanner.Scan() {
		log.Fatalln(scanner.Err())
	}
	vertices, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln(err)
	}

	currencies := make([]string, vertices)
	g := ex_ewdg.NewWeightedAdjList(vertices)
	for v := 0; v < vertices; v++ {
		if !scanner.Scan() {
			log.Fatalln(scanner.Err())
		}

		currencies[v] = scanner.Text()
		for w := 0; w < vertices; w++ {
			if !scanner.Scan() {
				log.Fatalln(scanner.Err())
			}
			rate, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				log.Fatalln(err)
			}
			// Taking the negative logarithm of the weight turns a multiplication problem (e.g. USD
			// to EUR via GBP = USD rate * GBP rate * EUR rate) into an addition problem that can be
			// solved by finding the shortest path. The greatest product of rates corresponds to the
			// smallest sum of the negative logarithms of each rate. The more negative the sum, the
			// greater the arbitrage opportunity. Positive sums mean there is no arbitrage
			// opportunity.
			g.AddEdge(ex_ewdg.NewDirectedEdge(v, w, -math.Log(rate)))
		}
	}

	bf := ex_bellman_ford.NewBellmanFordSP(g, 0) // source is arbitrary, since the graph is complete
	if bf.HasNegativeCycle() {
		stake := 1000.0
		for _, edge := range bf.NegativeCycle() {
			fmt.Printf("%10.5f %s ", stake, currencies[edge.From()])
			stake *= math.Exp(-edge.Weight())
			fmt.Printf("= %10.5f %s\n", stake, currencies[edge.To()])
		}
	} else {
		fmt.Println("No arbitrage opportunity.")
	}
}
