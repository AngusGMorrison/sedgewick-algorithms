package localminimum

import (
	"reflect"
	"testing"
)

func Test_DivideAndConquer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input [][]int
		want  []int
	}{
		// {
		// 	name:  "len(input) == 0",
		// 	input: nil,
		// 	want:  nil,
		// },
		{
			name: "len(input) == 1",
			input: [][]int{
				{1},
			},
			want: []int{0, 0},
		},
		{
			name: "minimum is in top-left corner",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: []int{0, 0},
		},
		{
			name: "minimum is in top-right corner",
			input: [][]int{
				{4, 3, 2, 1},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: []int{0, 3},
		},
		{
			name: "minimum is in bottom-right corner",
			input: [][]int{
				{13, 14, 15, 16},
				{9, 10, 11, 12},
				{5, 6, 7, 8},
				{4, 3, 2, 1},
			},
			want: []int{3, 3},
		},
		{
			name: "minimum is in bottom-left corner",
			input: [][]int{
				{13, 14, 15, 16},
				{9, 10, 11, 12},
				{5, 6, 7, 8},
				{1, 2, 3, 4},
			},
			want: []int{3, 0},
		},
		{
			name: "minimum is in top row",
			input: [][]int{
				{2, 1, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: []int{0, 1},
		},
		{
			name: "minimum is in bottom row",
			input: [][]int{
				{13, 14, 15, 16},
				{9, 10, 11, 12},
				{5, 6, 7, 8},
				{2, 1, 3, 4},
			},
			want: []int{3, 1},
		},
		{
			name: "minimum is in leftmost column",
			input: [][]int{
				{17, 18, 19, 20},
				{5, 6, 7, 8},
				{1, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: []int{2, 0},
		},
		{
			name: "minimum is in rightmost column",
			input: [][]int{
				{26, 27, 28, 13},
				{16, 15, 14, 1},
				{20, 19, 18, 17},
				{25, 24, 23, 22},
			},
			want: []int{1, 3},
		},
		{
			name: "minimum is in the middle",
			input: [][]int{
				{120, 89, 88, 81, 82},
				{118, 87, 41, 40, 83},
				{116, 90, 84, 39, 85},
				{114, 99, 95, 91, 86},
				{112, 110, 108, 106, 104},
			},
			want: []int{2, 3},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := DivideAndConquer(tc.input); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

func Test_RollDownhill(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{
			name:  "len(input) == 0",
			input: nil,
			want:  nil,
		},
		{
			name: "len(input) == 1",
			input: [][]int{
				{1},
			},
			want: []int{0, 0},
		},
		{
			name: "minimum is in top-left corner",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: []int{0, 0},
		},
		{
			name: "minimum is in top-right corner",
			input: [][]int{
				{4, 3, 2, 1},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: []int{0, 3},
		},
		{
			name: "minimum is in bottom-right corner",
			input: [][]int{
				{13, 14, 15, 16},
				{9, 10, 11, 12},
				{5, 6, 7, 8},
				{4, 3, 2, 1},
			},
			want: []int{3, 3},
		},
		{
			name: "minimum is in bottom-left corner",
			input: [][]int{
				{13, 14, 15, 16},
				{9, 10, 11, 12},
				{5, 6, 7, 8},
				{1, 2, 3, 4},
			},
			want: []int{3, 0},
		},
		{
			name: "minimum is in top row",
			input: [][]int{
				{2, 1, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: []int{0, 1},
		},
		{
			name: "minimum is in bottom row",
			input: [][]int{
				{13, 14, 15, 16},
				{9, 10, 11, 12},
				{5, 6, 7, 8},
				{2, 1, 3, 4},
			},
			want: []int{3, 1},
		},
		{
			name: "minimum is in leftmost column",
			input: [][]int{
				{17, 18, 19, 20},
				{5, 6, 7, 8},
				{1, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: []int{2, 0},
		},
		{
			name: "minimum is in rightmost column",
			input: [][]int{
				{26, 27, 28, 13},
				{16, 15, 14, 1},
				{20, 19, 18, 17},
				{25, 24, 23, 22},
			},
			want: []int{1, 3},
		},
		{
			name: "when two or more minima exist, it follows the path of smallest neighbours from the starting element",
			input: [][]int{
				{17, 18, 19, 20},
				{10, 11, 12, 1},
				{5, 6, 7, 8},
				{13, 14, 15, 16},
			},
			want: []int{2, 0},
		},
		{
			name: "path to minmum follows a spiral",
			input: [][]int{
				{89, 88, 81, 82},
				{87, 41, 40, 83},
				{35, 84, 39, 85},
				{36, 37, 38, 86},
			},
			want: []int{2, 0},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := RollDownhill(tc.input); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}
