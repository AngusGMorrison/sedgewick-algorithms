package q11_improvements

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	// t.Parallel()

	testCases := []struct {
		name string
		s    []int
	}{
		{
			name: "length is 0",
			s:    []int{},
		},
		{
			name: "length is 1",
			s:    randomIntSlice(1),
		},
		{
			name: "length is 2",
			s:    []int{2, 1},
		},
		{
			name: "length is insertionSortThreshold",
			s:    randomIntSlice(insertionSortThreshold),
		},
		{
			name: "length is insertionSortThreshold + 1",
			s:    randomIntSlice(insertionSortThreshold + 1),
		},
		{
			name: "length is even",
			s:    randomIntSlice(100),
		},
		{
			name: "length is odd",
			s:    randomIntSlice(101),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

			stdlib := make([]int, len(tc.s))
			copy(stdlib, tc.s)
			sort.Ints(stdlib)
			MergeSort(tc.s)

			if !reflect.DeepEqual(tc.s, stdlib) {
				t.Errorf("want sorted slice, got\n\t%v", tc.s)
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
