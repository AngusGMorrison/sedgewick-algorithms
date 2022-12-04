package q29_steque_two_stacks

import (
	"reflect"
	"testing"
)

func Test_steque(t *testing.T) {
	t.Parallel()

	s := newSteque[int]()
	s.push(1)
	s.push(2)
	s.enqueue(3)
	s.push(4)
	s.enqueue(5)

	var got []int
	for elem, ok := s.pop(); ok; elem, ok = s.pop() {
		got = append(got, elem)
	}
	want := []int{4, 2, 1, 3, 5}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
