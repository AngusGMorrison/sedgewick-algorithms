package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdag_shortest_paths"
	"github.com/angusgmorrison/sedgewick_algorithms/c4_graphs/s4_shortest_paths/ex_ewdg"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatalln(scanner.Err())
	}
	vertices, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln(err)
	}

	// Instantiate a graph containing a start and and vertex for each vertex in the input, plus 2
	// vertices to serve as a source and a sink.
	g := ex_ewdg.NewWeightedAdjList(2*vertices + 2)
	source := 2 * vertices
	sink := source + 1

	for v := 0; v < vertices; v++ {
		if !scanner.Scan() {
			log.Fatalln(scanner.Err())
		}

		fields := strings.Fields(scanner.Text())
		duration, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			log.Fatalln(err)
		}

		vEnd := v + vertices
		// Create a weighted edge between a start and end vertex corresponding to one job.
		g.AddEdge(ex_ewdg.NewDirectedEdge(v, vEnd, duration))
		// Create a zero-weight edge from the source to the start of the job.
		g.AddEdge(ex_ewdg.NewDirectedEdge(source, v, 0))
		// Create a zero-weight edge from the end of the job to the sink.
		g.AddEdge(ex_ewdg.NewDirectedEdge(vEnd, sink, 0))
		// Create zero-weight edges between the end of the job and jobs that depend it.
		for j := 1; j < len(fields); j++ {
			dep, err := strconv.Atoi(fields[j])
			if err != nil {
				log.Fatalln(err)
			}
			g.AddEdge(ex_ewdg.NewDirectedEdge(vEnd, dep, 0))
		}
	}

	// Determine the longest path from the source to the sink. The length of any path from the
	// source s to any vertex v is a lower bound on the start/finish time represented by v, because
	// we could not do better than scheduling those jobs one after another. The length of the
	// longest path from source to sink is both the lower and the upper bound on the finish times of
	// all jobs, since we can complete them no faster (using a shorter route to the sink would mean
	// that some jobs would still be running when we arrived) nor any slower, since every job starts
	// immediately after the finish of all the jobs it depends on.
	longestPath, err := NewAcyclicLP(g, source)
	if err != nil {
		log.Fatalln(err)
	}

	for v := 0; v < vertices; v++ {
		fmt.Printf("%4d: %5.1f\n", v, longestPath.WeightTo(v))
	}
	fmt.Printf("Total time: %5.1f\n", longestPath.WeightTo(sink))
}

type AcylicLP struct {
	*ex_ewdag_shortest_paths.AcyclicSP
}

func NewAcyclicLP(g ex_ewdg.EdgeWeightedDigraph, source int) (*AcylicLP, error) {
	negatedSP, err := ex_ewdag_shortest_paths.NewAcyclicSP(invertWeights(g), source)
	if err != nil {
		return nil, fmt.Errorf("NewAcyclicSP: %w", err)
	}

	return &AcylicLP{negatedSP}, nil
}

func invertWeights(g ex_ewdg.EdgeWeightedDigraph) ex_ewdg.EdgeWeightedDigraph {
	negation := ex_ewdg.NewWeightedAdjList(g.V())
	for _, edge := range g.Edges() {
		negation.AddEdge(ex_ewdg.NewDirectedEdge(edge.From(), edge.To(), -edge.Weight()))
	}
	return negation
}
