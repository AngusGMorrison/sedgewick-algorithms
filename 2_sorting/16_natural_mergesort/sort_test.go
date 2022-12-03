package sort

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	// t.Parallel()

	testCases := []struct {
		len int
	}{
		{0},
		{1},
		{2},
		{3},
		{11},
		{103},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(strconv.Itoa(tc.len), func(t *testing.T) {
			t.Parallel()

			s := randomIntSlice(tc.len)
			MergeSort(s)
			if !sort.IntsAreSorted(s) {
				t.Errorf("want sorted slice, got\n\t%v", s)
			}
		})
	}
}

func randomIntSlice(len int) []int {
	ints := make([]int, len)
	for i := range ints {
		ints[i] = rand.Intn(100)
	}
	return ints
}
