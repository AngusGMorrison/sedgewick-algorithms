package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var errUsage = errors.New("usage: gcd int int")

func run() error {
	if len(os.Args) < 3 {
		return errUsage
	}

	p, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return err
	}

	q, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return err
	}

	result := gcd(p, q, 0)
	fmt.Println(result)

	return nil
}

func gcd(p, q, depth int) int {
	fmt.Printf("%sgcd(%d, %d)\n", strings.Repeat("\t", depth), p, q)

	if q == 0 {
		return p
	}

	return gcd(q, p%q, depth+1)
}
