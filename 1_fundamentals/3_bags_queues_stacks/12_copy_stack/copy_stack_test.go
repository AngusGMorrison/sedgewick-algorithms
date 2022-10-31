package stackutils

import (
	"reflect"
	"testing"

	"github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/bqs/stack"
)

func Test_CopySliceStack(t *testing.T) {
	t.Parallel()

	original := stack.NewSliceStack[int]()
	original.Push(1)
	original.Push(2)

	cp := CopySliceStack[int](original)
	if !reflect.DeepEqual(original, cp) {
		t.Errorf("want copy\n\t%+v,\ngot\n\t%+v", original, cp)
	}
}

func Test_CopyListStack(t *testing.T) {
	t.Parallel()

	original := stack.NewListStack[int]()
	original.Push(1)
	original.Push(2)

	cp := CopyListStack[int](original)
	if !reflect.DeepEqual(original, cp) {
		t.Errorf("want copy\n\t%+v,\ngot\n\t%+v", original, cp)
	}
}
