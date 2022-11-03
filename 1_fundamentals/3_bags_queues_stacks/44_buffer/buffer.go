package buffer

import "github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/bqs/stack"

// Buffer represents a buffer in a text editor, where the cursor is positioned at the top of the
// preCursor stack.
type Buffer struct {
	preCursor, postCursor *stack.SliceStack[rune]
}

func (b *Buffer) Len() int {
	return b.preCursor.Len() + b.postCursor.Len()
}

func (b *Buffer) Insert(r rune) {
	b.preCursor.Push(r)
}

func (b *Buffer) Delete() (rune, bool) {
	return b.preCursor.Pop()
}

func (b *Buffer) Left(k int) {
	for i := 0; i < k; i++ {
		r, ok := b.preCursor.Pop()
		if !ok { // leftmost position reached
			return
		}
		b.postCursor.Push(r)
	}
}

func (b *Buffer) Right(k int) {
	for i := 0; i < k; i++ {
		r, ok := b.postCursor.Pop()
		if !ok { // right position reached
			return
		}
		b.preCursor.Push(r)
	}
}
