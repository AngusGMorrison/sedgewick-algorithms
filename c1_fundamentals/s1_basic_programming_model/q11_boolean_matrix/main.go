package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	printMatrix(os.Stdout, [][]bool{
		{true, true, true},
		{false, true, false},
		{true, true, false},
	})
}

// printMatrix prints an n*n boolean matrix with row and column numbers, representing true values
// with "*" and false a space. It verifies that the there is at least one row and that the first row
// is not empty, but does not validate that the input matrix is square.
func printMatrix(w io.Writer, m [][]bool) {
	if len(m) == 0 || len(m[0]) == 0 {
		fmt.Println()
		return
	}

	var builder strings.Builder

	// Column numbers
	for i := 1; i <= len(m[0]); i++ {
		fmt.Fprintf(&builder, "\t%d", i)
	}
	builder.WriteByte('\n')

	// Rows
	for i := 0; i < len(m); i++ {
		builder.WriteString(strconv.Itoa(i + 1))

		for j := 0; j < len(m[i]); j++ {
			if m[i][j] {
				builder.WriteString("\t*")
				continue
			}

			builder.WriteString("\t ")
		}
		builder.WriteByte('\n')
	}

	fmt.Fprintln(w, builder.String())
}
