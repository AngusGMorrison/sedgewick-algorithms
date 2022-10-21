package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	if len(os.Args) != 4 {
		return _errUsage
	}

	var ints [3]int
	for i, arg := range os.Args[1:4] {
		intArg, err := strconv.Atoi(arg)
		if err != nil {
			return argumentError{position: i + 1, arg: arg}
		}

		ints[i] = intArg
	}

	if ints[0] == ints[1] && ints[1] == ints[2] {
		fmt.Println("equal")

		return nil
	}

	fmt.Println("not equal")

	return nil
}

var _errUsage = errors.New("usage: equal int int int")

type argumentError struct {
	position int
	arg      string
}

func (e argumentError) Error() string {
	return fmt.Sprintf("arg %d, %q, is not an int", e.position, e.arg)
}
