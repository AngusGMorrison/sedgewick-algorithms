package bs

import (
	"fmt"
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		key  int
		a    []int
		want int
	}{
		{
			key:  23,
			a:    nil,
			want: -1,
		},
		{
			key:  23,
			a:    []int{},
			want: -1,
		},
		{
			key:  23,
			a:    []int{10, 11, 12, 16, 18, 23, 29, 33, 48, 54, 57, 68, 77, 84, 98},
			want: 5,
		},
		{
			key:  50,
			a:    []int{10, 11, 12, 16, 18, 23, 29, 33, 48, 54, 57, 68, 77, 84, 98},
			want: -1,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("BinarySearch(%d, %v)", tc.key, tc.a), func(t *testing.T) {
			t.Parallel()

			if got := BinarySearch(tc.key, tc.a); got != tc.want {
				t.Fatalf("BinarySearch(%d, %v): want %d, got %d", tc.key, tc.a, got, tc.want)
			}
		})
	}
}
