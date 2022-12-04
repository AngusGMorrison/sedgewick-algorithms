package q20_bitonic_search

import "testing"

func Test_BitonicSearch(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
		key   int
		want  int
	}{
		{
			name:  "len(input) == 0",
			input: nil,
			key:   0,
			want:  -1,
		},
		{
			name:  "len(input) == 1, element does not match key",
			input: []int{1},
			key:   0,
			want:  -1,
		},
		{
			name:  "len(input) == 1, element matches key",
			input: []int{1},
			key:   1,
			want:  0,
		},
		{
			name:  "key is found in increasing part of input",
			input: []int{2, 4, 6, 8, 10, 9, 7, 5, 3, 1},
			key:   4,
			want:  1,
		},
		{
			name:  "key is found in decreasing part of input",
			input: []int{2, 4, 6, 8, 10, 9, 7, 5, 3, 1},
			key:   3,
			want:  8,
		},
		{
			name:  "key is first element",
			input: []int{2, 4, 6, 8, 10, 9, 7, 5, 3, 1},
			key:   2,
			want:  0,
		},
		{
			name:  "key is last element",
			input: []int{2, 4, 6, 8, 10, 9, 7, 5, 3, 1},
			key:   1,
			want:  9,
		},
		{
			name:  "key is not present",
			input: []int{2, 4, 6, 8, 10, 9, 7, 5, 3, 1},
			key:   11,
			want:  -1,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := BitonicSearch(tc.input, tc.key); got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func Test_transitionPoint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "len(input) == 0",
			input: nil,
			want:  -1,
		},
		{
			name:  "len(input) == 1",
			input: []int{1},
			want:  -1,
		},
		{
			name:  "len(input) == 2",
			input: []int{1, -1},
			want:  0,
		},
		{
			name:  "transition point is at start of slice",
			input: []int{8, -3, -5, -11, -14},
			want:  0,
		},
		{
			name:  "transition point is at end of slice",
			input: []int{8, 12, 16, 17, 20},
			want:  4,
		},
		{
			name:  "transition point is in middle of slice",
			input: []int{8, 12, 16, 13, 9, -2, -5},
			want:  2,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := transitionPoint(tc.input); got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
