package q14_foursum

import "testing"

func Test_FourSum(t *testing.T) {
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
			name:      "no fours sum to zero",
			input:     []int{1, 2, 3, 4},
			wantCount: 0,
		},
		{
			name:      "valid fours are present",
			input:     []int{1, 2, 3, -6, 5, 6, 7, -18},
			wantCount: 2,
		},
		{
			name:      "double couting does not occur",
			input:     []int{2, 2, -2, 3}, // -2 should not be counted twice to form a four
			wantCount: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if gotCount := FourSum(tc.input); gotCount != tc.wantCount {
				t.Errorf("want %d, got %d", tc.wantCount, gotCount)
			}
		})
	}
}
