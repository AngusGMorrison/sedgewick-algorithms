package q27_queue_two_stacks

import (
	"reflect"
	"testing"
)

func Test_queue(t *testing.T) {
	t.Parallel()

	t.Run("items are popped in enqueued order", func(t *testing.T) {
		t.Parallel()

		q := newQueue[int]()
		input := []int{1, 2, 3}
		for _, elem := range input {
			q.enqueue(elem)
		}

		var output []int
		for elem, ok := q.dequeue(); ok; elem, ok = q.dequeue() {
			output = append(output, elem)
		}

		if !reflect.DeepEqual(input, output) {
			t.Errorf("want %v, got %v", input, output)
		}
	})
}
