package localminimum

import "testing"

func Test_LocalMinimumRecursive(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "input is empty",
			input: nil,
			want:  -1,
		},
		{
			name:  "len(input) == 1",
			input: []int{1},
			want:  0,
		},
		{
			name:  "len(input) == 2, min is first element",
			input: []int{1, 2},
			want:  0,
		},
		{
			name:  "len(input) == 2, min is second element",
			input: []int{2, 1},
			want:  1,
		},
		{
			name:  "local minimum is the first element",
			input: []int{15, 20, 25, 30, 35},
			want:  0,
		},
		{
			name:  "local minimum is the last element",
			input: []int{35, 30, 25, 20, 15},
			want:  4,
		},
		{
			name:  "local minimum is a middle element",
			input: []int{35, 30, 25, 26, 27},
			want:  2,
		},
		{
			name:  "multiple local minima",
			input: []int{35, 30, 34, 25, 21, 18, 27},
			want:  5,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := LocalMinimumRecursive(tc.input); got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func Test_LocalMinimumLoop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "input is empty",
			input: nil,
			want:  -1,
		},
		{
			name:  "len(input) == 1",
			input: []int{1},
			want:  0,
		},
		{
			name:  "len(input) == 2, min is first element",
			input: []int{1, 2},
			want:  0,
		},
		{
			name:  "len(input) == 2, min is second element",
			input: []int{2, 1},
			want:  1,
		},
		{
			name:  "local minimum is the first element",
			input: []int{15, 20, 25, 30, 35},
			want:  0,
		},
		{
			name:  "local minimum is the last element",
			input: []int{35, 30, 25, 20, 15},
			want:  4,
		},
		{
			name:  "local minimum is a middle element",
			input: []int{35, 30, 25, 26, 27},
			want:  2,
		},
		{
			name:  "multiple local minima",
			input: []int{35, 30, 34, 25, 21, 18, 27},
			want:  5,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := LocalMinimumLoop(tc.input); got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
