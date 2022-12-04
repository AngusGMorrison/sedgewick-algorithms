package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := flag.Uint("n", 0, "The number of intervals to read in.")
	flag.Parse()

	ivls, err := readIntervals(*n, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ivls.intersectingPairs())
}

type intervals []interval

func (i intervals) intersectingPairs() []intervals {
	var isxn []intervals
	for j := 0; j < len(i); j++ {
		for k := j + 1; k < len(i); k++ {
			if i[j].intersects(i[k]) {
				isxn = append(isxn, []interval{i[j], i[k]})
			}
		}
	}

	return isxn
}

type interval struct {
	start, end float64
}

func (i interval) intersects(j interval) bool {
	return i.start <= j.end && i.end >= j.start
}

func readIntervals(n uint, r io.Reader) (intervals, error) {
	scanner := bufio.NewScanner(r)
	ivls := make(intervals, 0, n)
	for i := 0; i < int(n) && scanner.Scan(); i++ {
		ivl, err := parseInterval(scanner.Text())
		if err != nil {
			return nil, err
		}
		ivls = append(ivls, ivl)
	}

	return ivls, nil
}

func parseInterval(raw string) (interval, error) {
	fields := strings.Split(raw, ",")
	if len(fields) != 2 {
		return interval{}, fmt.Errorf("interval must consist of two comma-separated floats: got %q", raw)
	}
	f1, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return interval{}, err
	}
	f2, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return interval{}, err
	}
	if f1 >= f2 {
		return interval{}, fmt.Errorf("interval start must be < interval end: got %q", raw)
	}

	return interval{start: f1, end: f2}, nil
}
