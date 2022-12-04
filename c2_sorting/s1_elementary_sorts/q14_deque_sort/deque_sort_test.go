package q14_deque_sort

import (
	"math/rand"
	"testing"
)

func Test_Sort(t *testing.T) {
	t.Parallel()

	deck := NewDeck(rand.NewSource(123456789))
	deck.Sort()
	if !deck.Sorted() {
		t.Errorf("want sorted deck, got\n\t%v", deck.data)
	}
}
