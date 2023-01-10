package ex_ternary_search_trie

type TST[V any] struct {
	root *node[V]
}

func New[V any]() *TST[V] {
	return &TST[V]{}
}

func (t *TST[V]) Get(key string) (V, bool) {
	n := t.root.get([]rune(key), 0)
	if n == nil || !n.hasVal {
		return *new(V), false
	}

	return n.val, true
}

func (t *TST[V]) Put(key string, val V) {
	t.root = t.root.put([]rune(key), val, 0)
}

func (t *TST[V]) Keys() []string {
	return t.root.collect(nil, nil)
}

func (t *TST[V]) KeysWithPrefix(prefix string) []string {
	if prefix == "" {
		return t.Keys()
	}

	runePrefix := []rune(prefix)
	prefixRoot := t.root.get(runePrefix, 0)
	if prefixRoot == nil {
		return nil
	}

	var keys []string
	if prefixRoot.hasVal {
		keys = append(keys, prefix)
	}

	return prefixRoot.mid.collect(runePrefix, keys)
}

func (t *TST[V]) LongestPrefixOf(s string) string {
	maxLen := t.root.maxPrefixLen([]rune(s), 0, 0)
	return string([]rune(s)[:maxLen])
}

func (t *TST[V]) Delete(key string) {
	if len(key) == 0 {
		return
	}
	t.root = t.root.delete([]rune(key), 0)
}

type node[V any] struct {
	r                rune
	val              V
	hasVal           bool
	left, mid, right *node[V]
}

func (n *node[V]) get(key []rune, idx int) *node[V] {
	if n == nil || len(key) == 0 {
		return n
	}

	r := key[idx]
	if r < n.r {
		// For the less-than and greater-than cases, we don't advance the index counter because
		// we've haven't reached the correct spot (alphabetically) for r yet.
		return n.left.get(key, idx)
	} else if r > n.r {
		return n.right.get(key, idx)
	} else if idx < len(key)-1 { // handle double letters
		return n.mid.get(key, idx+1)
	} else {
		return n
	}
}

func (n *node[V]) put(key []rune, val V, idx int) *node[V] {
	r := key[idx]
	if n == nil {
		n = &node[V]{r: r}
	}

	if r < n.r {
		// For the less-than and greater-than cases, we don't advance the index counter because
		// we've haven't reached the correct spot (alphabetically) for r yet.
		n.left = n.left.put(key, val, idx)
	} else if r > n.r {
		n.right = n.right.put(key, val, idx)
	} else if idx < len(key)-1 {
		n.mid = n.mid.put(key, val, idx+1)
	} else {
		n.val = val
		n.hasVal = true
	}

	return n
}

func (n *node[V]) collect(prefix []rune, keys []string) []string {
	if n == nil {
		return keys
	}

	// The current node does not contain a letter belonging to its left and right children. Words
	// that incorporate its left child (or its left children) come alphabetically before words that
	// incorporate the current node. Words that incorporate its right child (or its children) come
	// alphabetically after words that incorporate the current node.

	keys = n.left.collect(prefix, keys) // left children do not include the current node - use the old prefix
	nextPrefix := append(prefix, n.r)
	if n.hasVal {
		keys = append(keys, string(nextPrefix))
	}
	keys = n.mid.collect(nextPrefix, keys) // middle children include the current node - use the new prefix
	keys = n.right.collect(prefix, keys)   // right children do not include the current node - use the old prefix
	return keys
}

func (n *node[V]) maxPrefixLen(key []rune, idx int, length int) int {
	if n == nil {
		return length
	}

	if n.hasVal {
		length = idx + 1
	}

	r := key[idx]
	if r < n.r {
		return n.left.maxPrefixLen(key, idx, length)
	} else if r > n.r {
		return n.right.maxPrefixLen(key, idx, length)
	} else if idx < len(key)-1 {
		return n.mid.maxPrefixLen(key, idx+1, length)
	} else {
		return length
	}
}

func (n *node[V]) delete(key []rune, idx int) *node[V] {
	if n == nil {
		return nil
	}

	r := key[idx]
	if r < n.r {
		// For the less-than and greater-than cases, we don't advance the index counter because
		// we've haven't reached the correct spot (alphabetically) for r yet.
		n.left = n.left.delete(key, idx)
	} else if r > n.r {
		n.right = n.right.delete(key, idx)
	} else if idx < len(key)-1 {
		n.mid = n.mid.delete(key, idx+1)
	} else {
		n.val = *new(V)
		n.hasVal = false
	}

	if n.isUnused() {
		return nil
	}
	return n
}

func (n *node[V]) isUnused() bool {
	return !n.hasVal && n.left == nil && n.mid == nil && n.right == nil
}
