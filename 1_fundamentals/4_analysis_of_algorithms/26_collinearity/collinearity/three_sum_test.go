package collinearity

import "testing"

func Test_threeSum(t *testing.T) {
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
			name:      "slice contains valid triples",
			input:     []int{-2, -1, 0, 1, 2},
			wantCount: 2,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if gotCount := threeSum(tc.input); gotCount != tc.wantCount {
				t.Errorf("want %d, got %d", tc.wantCount, gotCount)
			}
		})
	}
}
