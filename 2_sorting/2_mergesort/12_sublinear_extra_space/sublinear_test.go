package sublinear

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
	}{
		// {
		// 	name:  "len(input) == 0",
		// 	input: []int{},
		// },
		// {
		// 	name:  "len(input) < blockSize",
		// 	input: randomIntSlice(blockSize - 1),
		// },
		// {
		// 	name:  "len(input) == blockSize",
		// 	input: randomIntSlice(blockSize),
		// },
		{
			name:  "blockSize | len(input)",
			input: randomIntSlice(3 * blockSize),
		},
		// {
		// 	name:  "blockSize ∤ len(input)",
		// 	input: randomIntSlice(3*blockSize + (blockSize - 1)),
		// },
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			sortedCopy := make([]int, len(tc.input))
			copy(sortedCopy, tc.input)
			sort.Ints(sortedCopy)

			mergeSort(tc.input)
			if !reflect.DeepEqual(tc.input, sortedCopy) {
				t.Errorf("want sorted slice, got\n\t%v", tc.input)
			}
		})
	}
}

func Test_selectionSortBlocks(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "slice is empty",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "len(slice) == 1",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "len(slice) < blockSize",
			input: []int{2, 1},
			want:  []int{2, 1},
		},
		{
			name:  "len(slice) == blockSize",
			input: []int{2, 1, 4, 7, 3},
			want:  []int{2, 1, 4, 7, 3},
		},
		{
			name:  "blockSize ∣ len(slice), first block has smaller key",
			input: []int{2, 1, 4, 7, 3, 8, 5, 6, 9, 0},
			want:  []int{2, 1, 4, 7, 3, 8, 5, 6, 9, 0},
		},
		{
			name:  "blockSize ∣ len(slice), first block has larger key",
			input: []int{2, 1, 4, 7, 3, 0, 5, 6, 9, 8},
			want:  []int{0, 5, 6, 9, 8, 2, 1, 4, 7, 3},
		},
		{
			name:  "blockSize ∤ len(slice), first block has smaller key",
			input: []int{2, 1, 4, 7, 3, 8, 5, 6},
			want:  []int{2, 1, 4, 7, 3, 8, 5, 6},
		},
		{
			name:  "blockSize ∤ len(slice), first block has larger key",
			input: []int{2, 1, 4, 7, 3, 0, 5, 6},
			// input: []int{0, 5, 6, 7, 3, 2, 1, 4},
			want: []int{0, 5, 6, 2, 1, 4, 7, 3},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			selectionSortBlocks(tc.input, 0, len(tc.input)-1, blockSize)

			if !reflect.DeepEqual(tc.input, tc.want) {
				t.Errorf("want\n\t%v\ngot\n\t%v", tc.want, tc.input)
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
