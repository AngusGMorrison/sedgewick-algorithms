package q28_stack_queue

import (
	"reflect"
	"testing"
)

func Test_stack(t *testing.T) {
	t.Parallel()

	s := newStack[int]()
	input := []int{1, 2, 3}
	for _, elem := range input {
		s.push(elem)
	}

	want := []int{3, 2, 1}
	var got []int
	for elem, ok := s.pop(); ok; elem, ok = s.pop() {
		got = append(got, elem)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
