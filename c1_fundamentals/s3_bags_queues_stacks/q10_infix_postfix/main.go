package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/stack"
)

func main() {
	if err := infixToPostfix(os.Stdin, os.Stdout); err != nil {
		log.Fatalln(err)
	}
}

var (
	errEmptyExprs = errors.New("attempted to pop expression from empty stack")
	errEmptyOps   = errors.New("attempted to pop operator from empty stack")
)

func infixToPostfix(r io.Reader, w io.Writer) error {
	exprs := stack.NewSliceStack[string]()
	ops := stack.NewSliceStack[string]()
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		switch t := scanner.Text(); t {
		case "(":
		case "+", "-", "*", "/", "sqrt":
			ops.Push(t)
		case ")":
			op, ok := ops.Pop()
			if !ok {
				return errEmptyOps
			}
			right, ok := exprs.Pop()
			if !ok {
				return errEmptyExprs
			}

			var pf string
			switch op {
			case "sqrt":
				pf = postfixifyUnary(right, op)
			default:
				left, ok := exprs.Pop()
				if !ok {
					return errEmptyExprs
				}
				pf = postfixifyBinary(left, right, op)
			}

			exprs.Push(pf)
		default:
			exprs.Push(t)
		}
	}

	postfix, ok := exprs.Pop()
	if !ok {
		return errEmptyExprs
	}

	_, _ = fmt.Fprintln(w, postfix)
	return nil
}

func postfixifyUnary(arg, op string) string {
	return fmt.Sprintf("%s %s", arg, op)
}

func postfixifyBinary(left, right, op string) string {
	return fmt.Sprintf("%s %s %s", left, right, op)
}
