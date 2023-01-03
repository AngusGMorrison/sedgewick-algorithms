package ex_binary_tree

import "golang.org/x/exp/constraints"

type BST[K constraints.Ordered, V any] struct {
	root *node[K, V]
}

func (bst *BST[K, V]) Size() int {
	return bst.root.size()
}

func (bst *BST[K, V]) Get(key K) (V, bool) {
	return bst.root.get(key)
}

func (bst *BST[K, V]) Put(key K, val V) {
	bst.root = bst.root.put(key, val)
}

func (bst *BST[K, V]) Min() (K, bool) {
	min := bst.root.min()
	if min == nil {
		return *new(K), false
	}

	return min.key, true
}

func (bst *BST[K, V]) Max() (K, bool) {
	max := bst.root.max()
	if max == nil {
		return *new(K), false
	}

	return max.key, true
}

func (bst *BST[K, V]) Floor(key K) (K, bool) {
	return bst.root.floor(key)
}

// Select returns the key of the m+1th node in the tree.
func (bst *BST[K, V]) Select(m int) (K, bool) {
	return bst.root.slct(m)
}

// Rank returns the number of keys less than K in the BST.
func (bst *BST[K, V]) Rank(key K) (int, bool) {
	return bst.root.rank(key)
}

func (bst *BST[K, V]) DeleteMin() {
	bst.root = bst.root.deleteMin()
}

func (bst *BST[K, V]) DeleteMax() {
	bst.root = bst.root.deleteMax()
}

func (bst *BST[K, V]) Delete(key K) {
	bst.root = bst.root.delete(key)
}

func (bst *BST[K, V]) Each(fn func(val V)) {
	min, ok := bst.Min()
	if !ok {
		return
	}
	max, ok := bst.Max()
	if !ok {
		return
	}

	bst.root.eachInRange(min, max, fn)
}

type node[K constraints.Ordered, V any] struct {
	key         K
	val         V
	left, right *node[K, V]
	sz          int
}

func newNode[K constraints.Ordered, V any](key K, val V) *node[K, V] {
	return &node[K, V]{
		key: key,
		val: val,
		sz:  1,
	}
}

func (n *node[K, V]) size() int {
	if n == nil {
		return 0
	}

	return n.sz
}

func (n *node[K, V]) get(key K) (V, bool) {
	if n == nil {
		return *new(V), false
	}

	if n.key == key {
		return n.val, true
	} else if n.key < key {
		return n.right.get(key)
	} else {
		return n.left.get(key)
	}
}

func (n *node[K, V]) put(key K, val V) *node[K, V] {
	if n == nil {
		return newNode(key, val)
	}

	if n.key == key {
		n.val = val
	} else if n.key < key {
		n.right = n.right.put(key, val)
	} else {
		n.left = n.left.put(key, val)
	}

	n.sz = n.left.size() + n.right.size() + 1
	return n
}

func (n *node[K, V]) min() *node[K, V] {
	if n == nil {
		return nil
	}

	if n.left == nil {
		return n
	}

	return n.left.min()
}

func (n *node[K, V]) max() *node[K, V] {
	if n == nil {
		return nil
	}

	if n.right == nil {
		return n
	}

	return n.right.max()
}

func (n *node[K, V]) floor(key K) (K, bool) {
	if n == nil {
		return *new(K), false
	}

	if n.key == key {
		return key, true
	} else if n.key > key {
		return n.left.floor(key)
	} else {
		f, ok := n.right.floor(key)
		if ok {
			return f, ok
		}

		// If the right child is larger than key, the current node is the floor of the key
		return n.key, true
	}
}

func (n *node[K, V]) slct(m int) (K, bool) {
	if n == nil {
		return *new(K), false
	}

	leftSize := n.left.size()
	if leftSize == m { // n is the m+1th node
		return n.key, true
	} else if leftSize > m { // the m+1th node is somewhere in the right tree
		return n.left.slct(m)
	} else { // the m+1th node is somewhere in the right tree
		return n.right.slct(m - leftSize - 1) // exclude the size of the left subtree and the root node
	}
}

func (n *node[K, V]) rank(key K) (int, bool) {
	if n == nil {
		return 0, false
	}

	if n.key == key {
		return n.left.size(), true
	} else if n.key > key {
		return n.left.rank(key)
	} else {
		if rightRank, ok := n.right.rank(key); ok {
			// The rank of a node found in the right tree is the size of the left tree plus the
			// current node plus its rank in the right tree alone.
			return n.left.size() + 1 + rightRank, true
		}

		return 0, false
	}
}

func (n *node[K, V]) deleteMin() *node[K, V] {
	if n == nil {
		return nil
	}

	if n.left == nil { // n is the min node
		return n.right
	}

	n.left = n.left.deleteMin()
	n.sz = n.left.size() + n.right.size() + 1
	return n
}

func (n *node[K, V]) deleteMax() *node[K, V] {
	if n == nil {
		return nil
	}

	if n.right == nil { // n is the max node
		return n.left
	}

	n.right = n.right.deleteMin()
	n.sz = n.left.size() + n.right.size() + 1
	return n
}

func (n *node[K, V]) delete(key K) *node[K, V] {
	if n == nil {
		return nil
	}

	if n.key > key {
		n.left = n.left.delete(key)
	} else if n.key < key {
		n.right = n.right.delete(key)
	} else { // delete n
		if n.right == nil {
			return n.left
		}
		if n.left == nil {
			return n.right
		}

		deleted := n
		n = deleted.right.min()
		n.right = deleted.right.deleteMin()
		n.left = deleted.left
	}

	return n
}

func (n *node[K, V]) eachInRange(min, max K, fn func(val V)) {
	if n == nil {
		return
	}

	if n.key > min {
		n.left.eachInRange(min, max, fn)
	}
	if n.key >= min && n.key <= max {
		fn(n.val)
	}
	if n.key < max {
		n.right.eachInRange(min, max, fn)
	}

}
