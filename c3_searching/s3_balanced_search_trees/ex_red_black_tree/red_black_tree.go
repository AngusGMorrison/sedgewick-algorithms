package ex_red_black_tree

import "golang.org/x/exp/constraints"

type RedBlackTree[K constraints.Ordered, V any] struct {
	root *node[K, V]
}

func (rbt *RedBlackTree[K, V]) Size() int {
	return rbt.root.size()
}

func (rbt *RedBlackTree[K, V]) Put(key K, val V) {
	rbt.root = rbt.root.put(key, val)
	rbt.root.color = black
}

type color bool

const (
	red   color = true
	black color = false
)

type node[K constraints.Ordered, V any] struct {
	key         K
	val         V
	color       color
	left, right *node[K, V]
	sz          int
}

func newNode[K constraints.Ordered, V any](key K, val V) *node[K, V] {
	return &node[K, V]{
		key:   key,
		val:   val,
		color: red,
		sz:    1,
	}
}

func (n *node[K, V]) size() int {
	if n == nil {
		return 0
	}

	return n.sz
}

func (n *node[K, V]) isRed() bool {
	return n.color == red
}

func (n *node[K, V]) rotateLeft() *node[K, V] {
	// Rotation.
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	// Recolor.
	newRoot.color = n.color // preserve color of link to parent
	n.color = red           // left-leaning link is now red

	// Update sizes.
	newRoot.sz = n.sz
	n.sz = 1 + n.left.size() + n.right.size()

	return newRoot
}

func (n *node[K, V]) rotateRight() *node[K, V] {
	// Rotation.
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	// Recolor.
	newRoot.color = n.color
	n.color = red

	// Update sizes.
	newRoot.sz = n.sz
	n.sz = 1 + n.left.size() + n.right.size()

	return newRoot
}

func (n *node[K, V]) flipColors() {
	n.color = red
	n.left.color = black
	n.right.color = black
}

func (n *node[K, V]) put(key K, val V) *node[K, V] {
	if n == nil {
		return newNode(key, val)
	}

	if key < n.key {
		n.left = n.left.put(key, val)
	} else if key > n.key {
		n.right = n.right.put(key, val)
	} else {
		n.val = val
	}

	if n.right.isRed() && !n.left.isRed() {
		n = n.rotateLeft()
	}
	if n.left.isRed() && n.left.left.isRed() {
		n = n.rotateRight()
	}
	if n.left.isRed() && n.right.isRed() {
		n.flipColors()
	}

	n.sz = 1 + n.left.size() + n.right.size()
	return n
}
