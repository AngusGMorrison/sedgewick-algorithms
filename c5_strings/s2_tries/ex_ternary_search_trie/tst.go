package ex_ternary_search_trie

type TST[V any] struct {
	root *node[V]
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
	return t.KeysWithPrefix("")
}

func (t *TST[V]) KeysWithPrefix(prefix string) []string {
	runePrefix := []rune(prefix)
	prefixRoot := t.root.get(runePrefix, 0)
	return prefixRoot.collect(runePrefix)
}

func (t *TST[V]) LongestPrefixOf(s string) string {
	maxLen := t.root.maxPrefixLen([]rune(s), 0, 0)
	return string([]rune(s)[:maxLen])
}

func (t *TST[V]) Delete(key string) {
	t.root.delete([]rune(key), 0)
}

type node[V any] struct {
	r                rune
	val              V
	hasVal           bool
	left, mid, right *node[V]
}

func (n *node[V]) get(key []rune, idx int) *node[V] {
	if n == nil {
		return nil
	}

	r := key[idx]
	if r < n.r {
		return n.left.get(key, idx+1)
	} else if r > n.r {
		return n.right.get(key, idx+1)
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
		n.left.put(key, val, idx+1)
	} else if r > n.r {
		n.right.put(key, val, idx+1)
	} else if idx < len(key)-1 {
		n.mid = n.mid.put(key, val, idx+1)
	} else {
		n.val = val
		n.hasVal = true
	}

	return n
}

func (n *node[V]) collect(prefix []rune) []string {
	if n == nil {
		return nil
	}

	prefix = append(prefix, n.r)

	var keys []string
	if n.hasVal {
		keys = append(keys, string(prefix))
	}
	keys = append(keys, n.left.collect(prefix)...)
	keys = append(keys, n.mid.collect(prefix)...)
	keys = append(keys, n.right.collect(prefix)...)
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
		return n.left.maxPrefixLen(key, idx+1, length)
	} else if r > n.r {
		return n.right.maxPrefixLen(key, idx+1, length)
	} else if idx < len(key)-1 {
		return n.mid.maxPrefixLen(key, idx, length)
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
		n.left = n.left.delete(key, idx+1)
	} else if r > n.r {
		n.right = n.right.delete(key, idx+1)
	} else if idx < len(key)-1 {
		n.mid = n.mid.delete(key, idx+1)
	} else {
		n.val = *new(V)
		n.hasVal = false
	}

	if n.left == nil && n.mid == nil && n.right == nil {
		return nil
	}
	return n
}
