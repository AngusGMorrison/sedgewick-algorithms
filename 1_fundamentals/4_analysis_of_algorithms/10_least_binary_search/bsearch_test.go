package bsearch

import "testing"

func Test_BinarySearch(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		a         []int
		key       int
		wantIndex int
	}{
		{
			name:      "slice is empty",
			a:         nil,
			key:       1,
			wantIndex: -1,
		},
		{
			name:      "key is not present in slice",
			a:         []int{2},
			key:       1,
			wantIndex: -1,
		},
		{
			name:      "key is present in slice",
			a:         []int{1, 2, 3},
			key:       2,
			wantIndex: 1,
		},
		{
			name:      "key is repeated an even number of times",
			a:         []int{1, 2, 2, 2, 2, 3},
			key:       2,
			wantIndex: 1,
		},
		{
			name:      "key is repeated an odd number of times",
			a:         []int{1, 2, 2, 2, 2, 2, 3},
			key:       2,
			wantIndex: 1,
		},
		{
			name:      "slice starts with repeated key",
			a:         []int{2, 2, 2, 2, 2, 2, 3},
			key:       2,
			wantIndex: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if gotIndex := BinarySearch(tc.a, tc.key); gotIndex != tc.wantIndex {
				t.Errorf("want %d, got %d", tc.wantIndex, gotIndex)
			}
		})
	}
}
