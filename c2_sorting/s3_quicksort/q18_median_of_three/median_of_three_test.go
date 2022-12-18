package q18_median_of_three

import (
	"fmt"
	"sort"
	"testing"

	"github.com/angusgmorrison/sedgewick_algorithms/testutil"
)

func Test_QuickSort(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		inputLen int
	}{
		{0},
		{1},
		{3},
		{7},
		{13},
		{256},
		{1024},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("slice len %d", tc.inputLen), func(t *testing.T) {
			t.Parallel()

			s := testutil.RandomIntSlice(tc.inputLen)
			QuickSort(s)
			if !sort.IntsAreSorted(s) {
				t.Errorf("want sorted slice, got %v", s)
			}
		})
	}
}
