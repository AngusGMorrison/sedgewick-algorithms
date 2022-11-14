package fibonacci

import "testing"

func Test_FibonacciSearch(t *testing.T) {
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
			name:  "len(input) == 1, key present",
			input: []int{0},
			key:   0,
			want:  0,
		},
		{
			name:  "len(input) == 1, key not present",
			input: []int{1},
			key:   0,
			want:  -1,
		},
		{
			name:  "len(input) == 2, key is first element",
			input: []int{0, 1},
			key:   0,
			want:  0,
		},
		{
			name:  "len(input) == 2, key is last element",
			input: []int{0, 1},
			key:   1,
			want:  1,
		},
		{
			name:  "len(input) == 2, key not present",
			input: []int{0, 1},
			key:   2,
			want:  -1,
		},
		{
			name:  "len(input) == 3, key is first element",
			input: []int{0, 1, 2},
			key:   0,
			want:  0,
		},
		{
			name:  "len(input) == 3, key is middle element",
			input: []int{0, 1, 2},
			key:   1,
			want:  1,
		},
		{
			name:  "len(input) == 3, key is last element",
			input: []int{0, 1, 2},
			key:   2,
			want:  2,
		},
		{
			name:  "len(input) == 3, key is not present",
			input: []int{0, 1, 2},
			key:   3,
			want:  -1,
		},
		{
			name:  "len(input) > 3, key is present",
			input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
			key:   23,
			want:  23,
		},
		{
			name:  "len(input) > 3, key is not present",
			input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
			key:   27,
			want:  -1,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := FibonacciSearch(tc.input, tc.key); got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}
