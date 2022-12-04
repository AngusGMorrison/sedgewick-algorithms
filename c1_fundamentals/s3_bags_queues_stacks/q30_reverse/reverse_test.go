package q30_reverse

import (
	"reflect"
	"testing"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/list"
)

func Test_reverse(t *testing.T) {
	t.Parallel()

	// Convenience variables. Copy before use.
	node1 := list.Node[int]{Data: 1}
	node2 := list.Node[int]{Data: 2}
	node3 := list.Node[int]{Data: 3}

	t.Run("list is nil", func(t *testing.T) {
		t.Parallel()

		if got := reverse[int](nil); got != nil {
			t.Errorf("want nil, got %+v", got)
		}
	})

	t.Run("list has one node", func(t *testing.T) {
		t.Parallel()

		list := node1
		wantList := node1

		if got := reverse[int](&list); !reflect.DeepEqual(got, &wantList) {
			t.Errorf("want\n\t%s\ngot\n\t%s", &wantList, got)
		}
	})

	t.Run("list has two nodes", func(t *testing.T) {
		t.Parallel()

		initialSecond := node2
		initialFirst := node1
		initialFirst.Next = &initialSecond

		wantSecond := node1
		wantFirst := node2
		wantFirst.Next = &wantSecond

		if got := reverse[int](&initialFirst); !reflect.DeepEqual(got, &wantFirst) {
			t.Errorf("want\n\t%s\ngot\n\t%s", &wantFirst, got)
		}
	})

	t.Run("list has more than two nodes", func(t *testing.T) {
		t.Parallel()

		initialThird := node3
		initialSecond := node2
		initialSecond.Next = &initialThird
		initialFirst := node1
		initialFirst.Next = &initialSecond

		wantThird := node1
		wantSecond := node2
		wantSecond.Next = &wantThird
		wantFirst := node3
		wantFirst.Next = &wantSecond

		if got := reverse[int](&initialFirst); !reflect.DeepEqual(got, &wantFirst) {
			t.Errorf("want\n\t%s\ngot\n\t%s", &wantFirst, got)
		}
	})
}
