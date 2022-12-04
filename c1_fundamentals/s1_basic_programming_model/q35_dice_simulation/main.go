package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	const (
		sides  = 6
		trials = 10_000_000
	)
	exactDist := distribution(sides)
	empiricalDist := simulate(sides, trials)

	fmt.Printf("Sides: %d, Trials: %d\n", sides, trials)
	printDistribution("Exact Distribution", exactDist)
	printDistribution("Empirical Distribution", empiricalDist)
}

// distribution returns a []float64 at each index i of which is the probability of rolling i on two
// dice with the given number of sides.
func distribution(sides int) []float64 {
	dist := make([]float64, 2*sides+1)
	for i := 1; i <= sides; i++ {
		for j := 1; j <= sides; j++ {
			dist[i+j]++
		}
	}

	for k := 2; k <= 2*sides; k++ {
		dist[k] /= math.Pow(float64(sides), 2)
	}

	return dist
}

func simulate(sides, trials int) []float64 {
	src := rand.NewSource(time.Now().UnixNano())
	roller := diceRoller{
		sides: sides,
		rng:   rand.New(src),
	}
	dist := make([]float64, 2*sides+1)
	for i := 0; i < trials; i++ {
		rolled := roller.rollN(2)
		dist[rolled]++
	}

	for i := 2; i <= 2*sides; i++ {
		dist[i] /= float64(trials)
	}

	return dist
}

type diceRoller struct {
	sides int
	rng   *rand.Rand
}

func (dr *diceRoller) rollN(n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum += dr.roll()
	}

	return sum
}

func (dr *diceRoller) roll() int {
	return dr.rng.Intn(dr.sides) + 1
}

func printDistribution(title string, d []float64) {
	var builder strings.Builder
	builder.WriteString(title)
	builder.WriteByte('\n')
	for i := 1; i < len(d); i++ {
		fmt.Fprintf(&builder, "%s\t", strconv.Itoa(i))
	}
	builder.WriteByte('\n')
	for i := 1; i < len(d); i++ {
		fmt.Fprintf(&builder, "%.3f\t", d[i])
	}
	builder.WriteByte('\n')

	fmt.Println(builder.String())
}
