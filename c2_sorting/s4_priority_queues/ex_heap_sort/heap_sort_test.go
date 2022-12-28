package ex_heap_sort

import (
	"sort"
	"testing"

	"github.com/angusgmorrison/sedgewick_algorithms/testutil"
)

func Test_HeapSort(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
	}{
		{
			name:  "empty input",
			input: nil,
		},
		{
			name:  "input with even number of elems",
			input: []int{2, 5, 3, 8, 9, 7, 6, 1, 4, 0},
			// input: testutil.RandomIntSlice(100),
		},
		{
			name:  "input with odd number of elems",
			input: testutil.RandomIntSlice(101),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			HeapSort(tc.input)
			if !sort.IntsAreSorted(tc.input) {
				t.Errorf("want sorted array, got %v", tc.input)
			}
		})
	}
}
