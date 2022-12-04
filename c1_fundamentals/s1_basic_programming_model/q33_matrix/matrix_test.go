package q33_matrix

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func Test_Dot(t *testing.T) {
	t.Parallel()

	t.Run("input vectors are unequal lengths", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			x, y []float64
		}{
			{
				x: []float64{1},
				y: nil,
			},
			{
				x: nil,
				y: []float64{1},
			},
		}

		for _, tc := range testCases {
			tc := tc
			name := fmt.Sprintf("Dot(%v, %v)", tc.x, tc.y)

			t.Run(name, func(t *testing.T) {
				t.Parallel()

				if got := Dot(tc.x, tc.y); got != math.Inf(-1) {
					t.Errorf("%s: want %.7f, got %.7f", name, math.Inf(-1), got)
				}
			})
		}
	})

	t.Run("input vectors are equal lengths", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			x, y []float64
			want float64
		}{
			{
				x:    nil,
				y:    nil,
				want: 0,
			},
			{
				x:    []float64{1},
				y:    []float64{2},
				want: 2,
			},
			{
				x:    []float64{1, 2},
				y:    []float64{2, 4},
				want: 10,
			},
		}

		for _, tc := range testCases {
			tc := tc
			name := fmt.Sprintf("Dot(%v, %v)", tc.x, tc.y)

			t.Run(name, func(t *testing.T) {
				t.Parallel()

				if got := Dot(tc.x, tc.y); !approxEqual(got, tc.want) {
					t.Errorf("%s: want %.7f, got %.7f", name, tc.want, got)
				}
			})
		}
	})
}

func Test_Mult(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		x, y [][]float64
		want [][]float64
	}{
		{
			name: "x and y are not multiplicatively conformable",
			x:    nil,
			y:    [][]float64{{1}},
			want: nil,
		},
		{
			name: "1x1 x 1x1",
			x:    [][]float64{{2}},
			y:    [][]float64{{2}},
			want: [][]float64{{4}},
		},
		{
			name: "2x1 x 1x2",
			x: [][]float64{
				{2},
				{1},
			},
			y: [][]float64{
				{2, 3},
			},
			want: [][]float64{
				{4, 6},
				{2, 3},
			},
		},
		{
			name: "2x2 x 2x2",
			x: [][]float64{
				{1, 2},
				{3, 4},
			},
			y: [][]float64{
				{2, 4},
				{3, 1},
			},
			want: [][]float64{
				{8, 6},
				{18, 16},
			},
		},
		{
			name: "3x3 x 3x3",
			x: [][]float64{
				{1, 2, 3},
				{3, 2, 1},
				{2, 1, 3},
			},
			y: [][]float64{
				{4, 5, 6},
				{6, 5, 4},
				{5, 4, 6},
			},
			want: [][]float64{
				{31, 27, 32},
				{29, 29, 32},
				{29, 27, 34},
			},
		},
		{
			name: "3x3 x I",
			x: [][]float64{
				{1, 2, 3},
				{3, 2, 1},
				{2, 1, 3},
			},
			y: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: [][]float64{
				{1, 2, 3},
				{3, 2, 1},
				{2, 1, 3},
			},
		},
		{
			name: "I x 3x3",
			x: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			y: [][]float64{
				{1, 2, 3},
				{3, 2, 1},
				{2, 1, 3},
			},
			want: [][]float64{
				{1, 2, 3},
				{3, 2, 1},
				{2, 1, 3},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := Mult(tc.x, tc.y); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nwant:\n\t%v,\ngot:\n\t%v", tc.want, got)
			}
		})
	}
}

func Test_Transpose(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input [][]float64
		want  [][]float64
	}{
		{
			name:  "input matrix is empty",
			input: nil,
			want:  nil,
		},
		{
			name: "input matrix is non-square",
			input: [][]float64{
				{1},
				{1, 2},
			},
			want: nil,
		},
		{
			name: "input matrix has length 1",
			input: [][]float64{
				{1},
			},
			want: [][]float64{
				{1},
			},
		},
		{
			name: "input matrix has length 2",
			input: [][]float64{
				{1, 2},
				{3, 4},
			},
			want: [][]float64{
				{1, 3},
				{2, 4},
			},
		},
		{
			name: "input matrix has length 3",
			input: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]float64{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := Transpose(tc.input); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nwant:\n\t%v,\ngot:\n\t%v", tc.want, got)
			}
		})
	}
}

func Test_multiplicablyConformable(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		x, y [][]float64
		want bool
	}{
		{
			name: "x has length 0",
			x:    nil,
			y: [][]float64{
				{1},
			},
			want: false,
		},
		{
			name: "y has length 0",
			x: [][]float64{
				{1},
			},
			y:    nil,
			want: false,
		},
		{
			name: "x[0] has length 0",
			x: [][]float64{
				nil,
			},
			y: [][]float64{
				{1},
			},
			want: false,
		},
		{
			name: "y[0] has length 0",
			x: [][]float64{
				{1},
			},
			y: [][]float64{
				nil,
			},
			want: false,
		},
		{
			name: "x has fewer columns than y has rows",
			x: [][]float64{
				{1},
			},
			y: [][]float64{
				{1},
				{2},
			},
			want: false,
		},
		{
			name: "x has more columns than y has rows",
			x: [][]float64{
				{1, 2},
			},
			y: [][]float64{
				{1},
			},
			want: false,
		},
		{
			name: "x contains rows of different lengths",
			x: [][]float64{
				{1, 2},
				{1, 2, 3},
			},
			y: [][]float64{
				{1},
				{2},
			},
			want: false,
		},
		{
			name: "y contains rows of different lengths",
			x: [][]float64{
				{1, 2},
				{1, 2, 3},
			},
			y: [][]float64{
				{1},
				{1, 2},
			},
			want: false,
		},
		{
			name: "x and y are multiplicable",
			x: [][]float64{
				{1, 2},
			},
			y: [][]float64{
				{1},
				{2},
			},
			want: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := multiplicativelyConformable(tc.x, tc.y); got != tc.want {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}
}

func Test_MultMatByVec(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		mat  [][]float64
		vec  []float64
		want [][]float64
	}{
		{
			name: "mat is nil",
			mat:  nil,
			vec:  []float64{1},
			want: nil,
		},
		{
			name: "vec is nil",
			mat:  [][]float64{{1}},
			vec:  nil,
			want: nil,
		},
		{
			name: "mat and vec are not multiplicatively conformable",
			mat: [][]float64{
				{1, 2},
				{3, 4},
			},
			vec:  []float64{1},
			want: nil,
		},
		{
			name: "mat and vec are multiplicatively conformable",
			mat: [][]float64{
				{1, 2},
			},
			vec: []float64{3, 4},
			want: [][]float64{
				{11},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := MultMatByVec(tc.mat, tc.vec); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nwant:\n\t%v,\ngot:\n\t%v", tc.want, got)
			}
		})
	}
}

func Test_MultVecByMat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		vec  []float64
		mat  [][]float64
		want [][]float64
	}{
		{
			name: "vec is nil",
			vec:  nil,
			mat:  [][]float64{{1}},
			want: nil,
		},
		{
			name: "mat is nil",
			vec:  []float64{1},
			mat:  nil,
			want: nil,
		},
		{
			name: "vec and mat are not multiplicatively conformable",
			vec:  []float64{1, 2},
			mat: [][]float64{
				{3},
				{4},
			},
			want: nil,
		},
		{
			name: "vec and mat are multiplicatively conformable",
			vec:  []float64{3, 4},
			mat: [][]float64{
				{1, 2, 3},
			},
			want: [][]float64{
				{3, 6, 9},
				{4, 8, 12},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := MultVecByMat(tc.vec, tc.mat); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nwant:\n\t%v,\ngot:\n\t%v", tc.want, got)
			}
		})
	}
}

const float64EqualityThreshold = 1e-7

func approxEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) <= float64EqualityThreshold
}
