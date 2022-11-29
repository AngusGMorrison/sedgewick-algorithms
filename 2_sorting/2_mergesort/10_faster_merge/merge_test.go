package merge

import (
	"math/rand"
	gosort "sort"
	"testing"
)

func Test_Sort(t *testing.T) {
	t.Parallel()

	s := make([]int, 100)
	for i := 0; i < len(s)-1; i++ {
		s[i] = rand.Int()
	}

	Sort(s)

	if !gosort.IntsAreSorted(s) {
		t.Errorf("want sorted slice, got\n\t%v", s)
	}
}
