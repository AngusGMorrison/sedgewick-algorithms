package main

import (
	"fmt"
	"math"
)

func main() {
	a()
	b()
	c()
}

func a() {
	t := 9.0
	for math.Abs(t-9.0/t) > 0.001 {
		t = (9.0/t + t) / 2.0
	}
	fmt.Printf("%.5f\n", t)
}

// This is the sum of natural numbers from 1-999.
func b() {
	var sum int
	for i := 1; i < 1000; i++ {
		for j := 0; j < i; j++ {
			sum++
		}
	}
	fmt.Println(sum)
}

// Starting from 1000, add 1000 each time i doubles. There are 9 doublings before i exceeds 1000.
func c() {
	var sum int
	for i := 1; i < 1000; i *= 2 {
		for j := 0; j < 1000; j++ {
			sum++
		}
	}
	fmt.Println(sum)
}
