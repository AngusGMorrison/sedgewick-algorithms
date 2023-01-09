package ex_rway_trie

const wildcard = '.'

type RWayTrie[E any] struct {
	radix int
	root  *node[E]
}

func (t *RWayTrie[E]) Get(key string) (E, bool) {
	n := t.root.get([]rune(key), 0)
	if n == nil || !n.hasVal {
		return *new(E), false
	}

	return n.val, true
}

func (t *RWayTrie[E]) Put(key string, val E) {
	t.root = t.root.put([]rune(key), val, 0, t.radix)
}

func (t *RWayTrie[E]) Keys() []string {
	return t.KeysWithPrefix("")
}

func (t *RWayTrie[E]) KeysWithPrefix(prefix string) []string {
	// The root of the trie containing all keys starting with prefix is the node containing the
	// final character of prefix.
	root := t.root.get([]rune(prefix), 0)
	return root.collect([]rune(prefix), t.radix)
}

func (t *RWayTrie[E]) KeysMatching(pattern string) []string {
	return t.root.match(nil, []rune(pattern), t.radix)
}

func (t *RWayTrie[E]) LongestPrefixOf(s string) string {
	length := t.root.maxPrefixLen([]rune(s), 0, 0)
	return string([]rune(s)[:length])
}

func (t *RWayTrie[E]) Delete(key string) {
	t.root = t.root.delete([]rune(key), 0, t.radix)
}

type node[E any] struct {
	val    E
	hasVal bool
	next   []*node[E]
}

func (n *node[E]) get(key []rune, idx int) *node[E] {
	if n == nil || idx == len(key) {
		return n // may be nil or contain no value if no such key exists
	}

	r := key[idx]
	return n.next[r].get(key, idx+1)
}

func (n *node[E]) put(key []rune, val E, idx int, radix int) *node[E] {
	if n == nil {
		n = &node[E]{
			next: make([]*node[E], radix),
		}
	}

	if idx == len(key) {
		n.val = val
		n.hasVal = true
		return n
	}

	r := key[idx]
	n.next[r] = n.next[r].put(key, val, idx+1, radix)
	return n
}

// collect iterates through the trie, appending each full word to strings as they are completed.
func (n *node[E]) collect(prefix []rune, radix int) []string {
	if n == nil {
		return nil
	}

	var keys []string
	if n.hasVal {
		// We've found a complete word, since the final letter (corresponding to the current node),
		// was appended to prefix in the previous recursion.
		keys = append(keys, string(prefix))
	}
	// On each iteration, extend the prefix with a different rune in the alphabet and continue the
	// search down the tree.
	for r := 0; r < radix; r++ {
		nextPrefix := append(prefix, rune(r))
		keys = append(keys, n.next[r].collect(nextPrefix, radix)...)
	}

	return keys
}

func (n *node[E]) match(prefix []rune, pattern []rune, radix int) []string {
	if n == nil {
		return nil
	}

	var keys []string
	// Once we reach the end of a pattern, we check to see if we've matched a whole word.
	if len(prefix) == len(pattern) {
		if n.hasVal { // found a full word
			keys = append(keys, string(prefix))
		}
		return keys // end of the pattern reached
	}

	// Continue building prefix according to the pattern. If the next character is a wildcard,
	// enqueue a recursive call for all runes in the alphabet appended to prefix. If the next
	// character is not a wildcard, enqueue one recursive call, appending only that character to the
	// prefix.
	nextRune := pattern[len(prefix)]
	for r := 0; r < radix; r++ {
		if nextRune == wildcard || nextRune == rune(r) {
			nextPrefix := append(prefix, rune(r))
			keys = append(keys, n.next[r].match(nextPrefix, pattern, radix)...)
		}
	}

	return keys
}

func (n *node[E]) maxPrefixLen(word []rune, idx int, length int) int {
	if n == nil {
		return length
	}
	if n.hasVal {
		// A full key has been found, making it the longest prefix yet.
		length = idx
	}
	if length == len(word) {
		// The whole word is present in the tree - it is its own longest prefix.
		return length
	}

	nextRune := word[idx]
	return n.next[nextRune].maxPrefixLen(word, idx+1, length)
}

func (n *node[E]) delete(key []rune, idx int, radix int) *node[E] {
	if n == nil {
		return nil
	}

	if idx == len(key) {
		// The final node in the key has been found.
		n.val = *new(E)
		n.hasVal = false
	} else {
		nextRune := key[idx]
		n.next[nextRune] = n.next[nextRune].delete(key, idx+1, radix)
	}

	if n.hasVal {
		// The node is not terminal in key, but has a value, meaning it is used by another key and
		// must be retained.
		return n
	}

	for _, r := range n.next {
		if r != nil {
			// If the current node does not have a value but has even a single a non-nil child, it
			// is part of a chain of nodes used by other keys and must be retained.
			return n
		}
	}

	// The current node holds no value and is not part of a chain used by any other key.
	return nil
}
