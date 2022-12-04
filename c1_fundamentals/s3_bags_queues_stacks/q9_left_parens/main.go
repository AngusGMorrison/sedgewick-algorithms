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
	if err := leftParenthesize(os.Stdin, os.Stdout); err != nil {
		log.Fatalln(err)
	}
}

var (
	errEmptyExprStack = errors.New("attempted to pop expression from empty stack")
	errEmptyOpsStack  = errors.New("attempted to pop operator from empty stack")
)

func leftParenthesize(r io.Reader, w io.Writer) error {
	exprs := stack.NewSliceStack[string]()
	ops := stack.NewSliceStack[string]()

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		switch t := scanner.Text(); t {
		case ")":
			op, ok := ops.Pop()
			if !ok {
				return errEmptyOpsStack
			}
			rhs, ok := exprs.Pop()
			if !ok {
				return errEmptyExprStack
			}

			var parenthesizedExpr string
			switch op {
			case "sqrt":
				parenthesizedExpr = parenthesizeUnary(rhs, op)
			default:
				lhs, ok := exprs.Pop()
				if !ok {
					return errEmptyExprStack
				}

				parenthesizedExpr = parenthesizeBinary(lhs, rhs, op)
			}
			exprs.Push(parenthesizedExpr)
		case "+", "-", "*", "/", "sqrt":
			ops.Push(t)
		default:
			exprs.Push(t)
		}
	}

	result, ok := exprs.Pop()
	if !ok {
		return errEmptyExprStack
	}

	_, _ = fmt.Fprintln(w, result)

	return nil
}

func parenthesizeUnary(arg, op string) string {
	return fmt.Sprintf("( %s %s )", op, arg)
}

func parenthesizeBinary(left, right, op string) string {
	return fmt.Sprintf("( %s %s %s )", left, op, right)
}
