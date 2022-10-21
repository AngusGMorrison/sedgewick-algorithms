package main

import (
	"reflect"
	"testing"
)

func Test_matrix_transpose(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		mat  matrix[[]int, int]
		want matrix[[]int, int]
	}{
		{
			name: "nil matrix",
			mat:  nil,
			want: nil,
		},
		{
			name: "nil row",
			mat:  nil,
			want: nil,
		},
		{
			name: "unequal row lengths",
			mat: matrix[[]int, int]{
				{1, 2, 3},
				{1, 2},
			},
			want: nil,
		},
		{
			name: "square",
			mat: matrix[[]int, int]{
				{1, 2},
				{3, 4},
			},
			want: matrix[[]int, int]{
				{1, 3},
				{2, 4},
			},
		},
		{
			name: "non-square",
			mat: matrix[[]int, int]{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: matrix[[]int, int]{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.mat.transpose(); !reflect.DeepEqual(tc.want, got) {
				t.Errorf("%s:\nwant:\n%+v\n\ngot%+v\n\n", tc.name, tc.want, got)
			}
		})
	}
}
