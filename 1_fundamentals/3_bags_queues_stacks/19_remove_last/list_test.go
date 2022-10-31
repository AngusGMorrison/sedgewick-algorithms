package list

import "testing"

func Test_removeLast(t *testing.T) {
	t.Parallel()

	t.Run("list is empty", func(t *testing.T) {
		t.Parallel()

		if got := removeLast[int](nil); got != nil {
			t.Errorf("want nil list, got first node %+v", got)
		}
	})

	t.Run("list has one node", func(t *testing.T) {
		t.Parallel()

		if got := removeLast[int](&node[int]{data: 1}); got != nil {
			t.Errorf("want nil list, got first node %+v", got)
		}
	})

	t.Run("list has two nodes", func(t *testing.T) {
		t.Parallel()

		second := &node[int]{data: 2}
		first := &node[int]{
			data: 1,
			next: second,
		}

		got := removeLast[int](first)
		if got != first {
			t.Errorf("want first node\n\t%+v,\ngot\n\t%+v", first, got)
		}
		if got.next != nil {
			t.Errorf("want only one node, got second node %+v", got.next)
		}
	})

	t.Run("list has more than two nodes", func(t *testing.T) {
		t.Parallel()

		third := &node[int]{data: 3}
		second := &node[int]{
			data: 2,
			next: third,
		}
		first := &node[int]{
			data: 1,
			next: second,
		}

		got := removeLast[int](first)
		if got != first {
			t.Errorf("want first node\n\t%+v,\ngot\n\t%+v", first, got)
		}
		if got.next != second {
			t.Errorf("want second node\t\t%+v,\ngot\n\t%+v", second, got.next)
		}
		if got.next.next != nil {
			t.Errorf("want only two nodes, got third node %+v", got.next.next)
		}
	})
}
