package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/stack"
	"golang.org/x/exp/slices"
)

func main() {
	if err := evaluatePostfix(os.Stdin, os.Stdout); err != nil {
		log.Fatalln(err)
	}
}

var errEmptyVals = errors.New("attempted to pop value from empty stack")

func evaluatePostfix(r io.Reader, w io.Writer) error {
	vals := stack.NewSliceStack[int]()
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		t := scanner.Text()
		if isOperator(t) {
			val, ok := vals.Pop()
			if !ok {
				return errEmptyVals
			}

			if isUnaryOperator(t) {
				switch t {
				case "sqrt":
					val = int(math.Sqrt(float64(val)))
				default:
					return fmt.Errorf("unhandled unary operator %q", t)
				}
			} else {
				lhs, ok := vals.Pop()
				if !ok {
					return errEmptyVals
				}
				switch t {
				case "+":
					val += lhs
				case "-":
					val = lhs - val
				case "*":
					val *= lhs
				case "/":
					val = lhs / val
				default:
					return fmt.Errorf("unhandled binary operator %q", t)
				}
			}
			vals.Push(val)
		} else {
			val, err := strconv.Atoi(t)
			if err != nil {
				return err
			}
			vals.Push(val)
		}
	}

	result, ok := vals.Pop()
	if !ok {
		return errEmptyVals
	}

	_, _ = fmt.Fprintln(w, result)
	return nil
}

var (
	binaryOperators = []string{"+", "-", "*", "/"}
	unaryOperators  = []string{"sqrt"}
	operators       = append(binaryOperators, unaryOperators...)
)

func isOperator(s string) bool {
	return slices.Contains(operators, s)
}

func isUnaryOperator(s string) bool {
	return slices.Contains(unaryOperators, s)
}
