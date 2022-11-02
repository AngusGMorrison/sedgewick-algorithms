package reverse

import "github.com/angusgmorrison/sedgewick_algorithms/1_fundamentals/bqs/list"

func reverse[E comparable](n *list.Node[E]) *list.Node[E] {
	first := n
	var reverse *list.Node[E]
	for first != nil {
		second := first.Next
		first.Next = reverse
		reverse = first
		first = second
	}

	return reverse
}

func reverseRecursive[E comparable](first *list.Node[E]) *list.Node[E] {
	if first == nil {
		return nil
	}
	if first.Next == nil {
		return first
	}
	second := first.Next
	rest := reverse(second)
	second.Next = first
	first.Next = nil
	return rest
}
