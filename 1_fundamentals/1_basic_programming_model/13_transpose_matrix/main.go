package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	mat := matrix[[]int, int]{
		{1, 2, 3},
		{4, 5, 6},
	}
	mat.printTranspose(os.Stdout)
}

type matrix[S ~[]E, E any] []S

func (mat matrix[S, E]) printTranspose(w io.Writer) {
	t := mat.transpose()

	for i := range t {
		var j int
		for ; j < len(t[i])-1; j++ {
			fmt.Fprintf(w, "%v\t", t[i][j])
		}
		fmt.Fprintf(w, "%v\n", t[i][j])
	}
	fmt.Fprintln(w)
}

// transpose takes a matrix with m rows and n columns and returns its transpose. If the matrix has
// no rows or no columns, or its rows are of uneven lengths, nil is returned.
func (mat matrix[S, E]) transpose() matrix[S, E] {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return nil
	}

	// Validate that matrix is m*n.
	for i := 1; i < len(mat); i++ {
		if len(mat[i]) != len(mat[0]) {
			return nil
		}
	}

	// Initialize an output matrix with n rows and m columns.
	result := make(matrix[S, E], len(mat[0]))
	for i := range result {
		result[i] = make(S, len(mat))
	}

	// Transpose rows and columns.
	for i := range mat {
		for j := range mat[i] {
			result[j][i] = mat[i][j]
		}
	}

	return result
}
