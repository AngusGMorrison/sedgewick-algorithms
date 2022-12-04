// Djikstra's two-Stack algorithm for expression evalulation.
package main

import (
	"bufio"
	"math"
	"os"
	"strconv"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/stack"
)

func main() {
	ops := stack.NewSliceStack[string]()
	vals := stack.NewSliceStack[float64]()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		switch t := scanner.Text(); t {
		case "(":
		case "+", "-", "*", "/", "sqrt":
			ops.Push(t)
		case ")":
			v, _ := vals.Pop()
			switch op, _ := ops.Pop(); op {
			case "+":
				left, _ := vals.Pop()
				v = left + v
			case "-":
				left, _ := vals.Pop()
				v = left - v
			case "*":
				left, _ := vals.Pop()
				v = left * v
			case "/":
				left, _ := vals.Pop()
				v = left / v
			case "sqrt":
				left, _ := vals.Pop()
				v = math.Sqrt(left)
			}
			vals.Push(v)
		default:
			v, _ := strconv.ParseFloat(t, 64)
			vals.Push(v)
		}
	}
}
