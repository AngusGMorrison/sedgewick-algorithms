package q33_matrix

import (
	"math"
)

// Dot returns the dot product of vectors x and y.
func Dot(x, y []float64) float64 {
	if len(x) != len(y) {
		return math.Inf(-1)
	}

	var dotProduct float64
	for i, f := range x {
		dotProduct += f * y[i]
	}

	return dotProduct
}

// Mult returns the result of multiplying matrices x and y.
func Mult(x, y [][]float64) [][]float64 {
	if !multiplicativelyConformable(x, y) {
		return nil
	}

	// If x is i×j and y is j×k, then result is i×k.
	result := make([][]float64, len(x))
	for i := range result {
		result[i] = make([]float64, len(y[0]))
	}

	// These loops can be arbitrarily permuted, with the fastest permutation depending on CPU
	// caching for the particular input matrices (typically ijk).
	//
	// i is the current row in x, which is also the current row of the result.
	for i := 0; i < len(x); i++ {
		// j is the current column in x and the current row in y.
		for j := 0; j < len(x[i]); j++ {
			// k is the current column in y, which is also the current column of the result.
			for k := 0; k < len(y[j]); k++ {
				// For each value of j (the row of x and the column of y), we add to the partial
				// result in each position [i][k] of the output matrix. When j == len(x[i])
				// (equivalent to j == len(y)), all the corresponding pairs from row x[i] and column
				// y[i][j] have been processed.
				result[i][k] += x[i][j] * y[j][k]
			}
		}
	}

	return result
}

// Transpose returns the transpose of the square matrix x.
func Transpose(x [][]float64) [][]float64 {
	if isEmpty(x) || !isSquare(x) {
		return nil
	}

	result := newMatrix(len(x), len(x))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x); j++ {
			result[i][j] = x[j][i]
		}
	}

	return result
}

// MultMatByVec returns the matrix created by multiplying the matrix mat by the vector vec.
func MultMatByVec(mat [][]float64, vec []float64) [][]float64 {
	vecM := vectorToMatrix(vec)
	return Mult(mat, vecM)
}

// MultMatByVec returns the matrix created by multiplying the vector vec by the matrix mat.
func MultVecByMat(vec []float64, mat [][]float64) [][]float64 {
	vecM := vectorToMatrix(vec)
	return Mult(vecM, mat)
}

func vectorToMatrix(v []float64) [][]float64 {
	if len(v) == 0 {
		return nil
	}

	m := newMatrix(len(v), 1)
	for i, f := range v {
		m[i][0] = f
	}

	return m
}

// newMatrix returns an initialized matrix with the specified number of rows and columns.
func newMatrix(rows int, cols int) [][]float64 {
	m := make([][]float64, rows)
	for i := range m {
		m[i] = make([]float64, cols)
	}

	return m
}

func multiplicativelyConformable(x, y [][]float64) bool {
	if isEmpty(x) || isEmpty(y) {
		return false
	}

	// x must have as many columns as y has rows.
	if len(x[0]) != len(y) {
		return false
	}

	// All rows in x must be the same length.
	if !rowLengthsEqual(x) || !rowLengthsEqual(y) {
		return false
	}

	return true
}

func isEmpty(x [][]float64) bool {
	if len(x) == 0 || len(x[0]) == 0 {
		return true
	}

	return false
}

func isSquare(x [][]float64) bool {
	if isEmpty(x) {
		return true
	}

	if len(x) == len(x[0]) && rowLengthsEqual(x) {
		return true
	}

	return false
}

func rowLengthsEqual(x [][]float64) bool {
	for i := 1; i < len(x); i++ {
		if len(x[i]) != len(x[0]) {
			return false
		}
	}

	return true
}
