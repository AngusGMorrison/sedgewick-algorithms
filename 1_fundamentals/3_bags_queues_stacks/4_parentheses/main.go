package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/bqs/stack"
)

func main() {
	fmt.Println(balanced(os.Stdin))
}

func balanced(r io.Reader) bool {
	parens := stack.NewSliceStack[string]()
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		switch p := scanner.Text(); p {
		case "(", "[", "{":
			parens.Push(p)
		case ")", "]", "}":
			openParen, ok := parens.Pop()
			if !ok {
				return false
			}

			switch p {
			case ")":
				if openParen != "(" {
					return false
				}
			case "]":
				if openParen != "[" {
					return false
				}
			case "}":
				if openParen != "{" {
					return false
				}
			}
		default:
			return false
		}
	}

	return true
}
