package q15_faster_sum

import "testing"

func Test_TwoSumLinear(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []int
		wantCount int
	}{
		{
			name:      "slice is empty",
			input:     nil,
			wantCount: 0,
		},
		{
			name:      "slice contains only negative values",
			input:     []int{-3, -2, -1, -0},
			wantCount: 0,
		},
		{
			name:      "slice contains only positive values",
			input:     []int{0, 1, 2, 3},
			wantCount: 0,
		},
		{
			name:      "slice contains negative and positive values but no pairs",
			input:     []int{-3, -1, 0, 2, 4},
			wantCount: 0,
		},
		{
			name:      "slice contains valid pairs but no duplicates",
			input:     []int{-3, -2, -1, 0, 1, 2, 3},
			wantCount: 3,
		},
		{
			name:      "slice contains valid pairs, including an even number of duplicates",
			input:     []int{-3, -2, -1, 1, 2, 3, 3},
			wantCount: 4,
		},
		{
			name:      "slice contains valid pairs, including an odd number of duplicates",
			input:     []int{-3, -2, -1, 1, 1, 2, 3, 3},
			wantCount: 5,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if gotCount := TwoSumLinear(tc.input); gotCount != tc.wantCount {
				t.Errorf("want %d, got %d", tc.wantCount, gotCount)
			}
		})
	}
}

func Test_ThreeSumQuadratic(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     []int
		wantCount int
	}{
		{
			name:      "slice is empty",
			input:     nil,
			wantCount: 0,
		},
		{
			name:      "slice contains only negative values",
			input:     []int{-3, -2, -1, -0},
			wantCount: 0,
		},
		{
			name:      "slice contains only positive values",
			input:     []int{0, 1, 2, 3},
			wantCount: 0,
		},
		{
			name:      "slice contains negative and positive values but no triples",
			input:     []int{-3, -1, 0, 2, 7},
			wantCount: 0,
		},
		{
			name:      "slice contains valid triples but no duplicates",
			input:     []int{-2, -1, 0, 1, 2},
			wantCount: 2,
		},
		{
			name:      "slice contains valid triples, including duplicates",
			input:     []int{-2, -2, -1, 0, 1, 2, 2},
			wantCount: 5,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if gotCount := ThreeSumQuadratic(tc.input); gotCount != tc.wantCount {
				t.Errorf("want %d, got %d", tc.wantCount, gotCount)
			}
		})
	}
}
