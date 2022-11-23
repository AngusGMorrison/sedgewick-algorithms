package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func Test_Shell(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		len int
	}{
		{len: 3},
		{len: 12},
		{len: 39},
		{len: 120},
		{len: 363},
		{len: 1092},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("len(s) == %d", tc.len), func(t *testing.T) {
			t.Parallel()

			src := rand.NewSource(123456789)
			rng := rand.New(src)
			s := make([]int, tc.len)
			for i := range s {
				s[i] = rng.Int()
			}

			Shell(s)
			if !sort.IntsAreSorted(s) {
				t.Errorf("want sorted array, got\n\t%v", s)
			}
		})
	}
}
