package ex_ternary_search_trie

import (
	"reflect"
	"sort"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_TST_Put(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []string
		wantTree func() *node[int]
	}{
		{
			name:  "single key",
			input: []string{"it"},
			wantTree: func() *node[int] {
				t := &node[int]{
					r:      't',
					val:    0,
					hasVal: true,
				}
				i := &node[int]{
					r:   'i',
					mid: t,
				}

				return i
			},
		},
		{
			name:  "multiple linear keys",
			input: []string{"it", "its"},
			wantTree: func() *node[int] {
				s := &node[int]{
					r:      's',
					val:    1,
					hasVal: true,
				}
				t := &node[int]{
					r:      't',
					val:    0,
					hasVal: true,
					mid:    s,
				}
				i := &node[int]{
					r:   'i',
					mid: t,
				}

				return i
			},
		},
		{
			name:  "multiple branching keys",
			input: []string{"it", "at", "ot"},
			wantTree: func() *node[int] {
				t0 := &node[int]{
					r:      't',
					val:    0,
					hasVal: true,
				}
				t1 := &node[int]{
					r:      't',
					val:    1,
					hasVal: true,
				}
				t2 := &node[int]{
					r:      't',
					val:    2,
					hasVal: true,
				}
				o := &node[int]{
					r:   'o',
					mid: t2,
				}
				a := &node[int]{
					r:   'a',
					mid: t1,
				}
				i := &node[int]{
					r:     'i',
					left:  a,
					mid:   t0,
					right: o,
				}

				return i
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			trie := New[int]()
			for i, str := range tc.input {
				trie.Put(str, i)
			}

			want := tc.wantTree()
			if !reflect.DeepEqual(trie.root, want) {
				t.Errorf("want trie %+v, got %+v", want, trie.root)
			}

		})
	}
}

func Test_TST_Get(t *testing.T) {
	t.Parallel()

	t.Run("search hits", func(t *testing.T) {
		trie := New[int]()
		input := []string{"she", "sells", "sea", "shells", "by", "the", "shore"}
		for i, str := range input {
			trie.Put(str, i)
		}

		for i, str := range input {
			val, ok := trie.Get(str)
			if !ok {
				t.Errorf("key %q not in trie", str)
			}

			if val != i {
				t.Errorf("key %q has value %d, want %d", str, val, i)
			}
		}
	})

	t.Run("search miss", func(t *testing.T) {
		trie := New[int]()
		val, ok := trie.Get("anything")
		if ok {
			t.Errorf("want ok to be false, got true")
		}
		if val != 0 {
			t.Errorf("want empty val, got %d", val)
		}
	})
}

func Test_TST_Keys(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []string
	}{
		{
			name:  "empty trie",
			input: nil,
		},
		{
			name:  "keys present",
			input: []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			trie := New[int]()
			for i, str := range tc.input {
				trie.Put(str, i)
			}

			got := trie.Keys()
			sort.Strings(tc.input)
			want := slices.Compact(tc.input)

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want\n\t%v\ngot\n\t%v", want, got)
			}
		})
	}
}

func Test_TST_Keys_With_Prefix(t *testing.T) {
	t.Parallel()

	trie := New[int]()
	for i, str := range []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"} {
		trie.Put(str, i)
	}

	testCases := []struct {
		name   string
		prefix string
		want   []string
	}{
		{
			name:   "empty prefix",
			prefix: "",
			want:   []string{"by", "sea", "sells", "she", "shells", "shore", "the"},
		},
		{
			name:   "non-empty prefix",
			prefix: "se",
			want:   []string{"sea", "sells"},
		},
		{
			name:   "prefix not in tree",
			prefix: "notfound",
			want:   nil,
		},
		{
			name:   "prefix is a valid key",
			prefix: "she",
			want:   []string{"she", "shells"},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := trie.KeysWithPrefix(tc.prefix)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want\n\t%v\ngot\n\t%v", tc.want, got)
			}
		})
	}
}

func Test_TST_LongestPrefixOf(t *testing.T) {
	t.Parallel()

	trie := New[int]()
	for i, str := range []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"} {
		trie.Put(str, i)
	}

	testCases := []struct {
		name string
		key  string
		want string
	}{
		{
			name: "no prefix of key present in trie",
			key:  "aurochs",
			want: "",
		},
		{
			name: "multiple prefixes of key present in trie",
			key:  "shellsort",
			want: "shells",
		},
		{
			name: "full prefix is present in trie",
			key:  "shells",
			want: "shells",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := trie.LongestPrefixOf(tc.key)
			if tc.want != got {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}

func Test_TST_Delete(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []string
		delete []string
		want   []string
	}{
		{
			name:   "key to delete has len 0",
			input:  []string{"she", "sells", "sea", "shells", "by", "the", "shore"},
			delete: []string{""},
			want:   []string{"by", "sea", "sells", "she", "shells", "shore", "the"},
		},
		{
			name:   "key to delete is not present",
			input:  []string{"she", "sells", "sea", "shells", "by", "the", "shore"},
			delete: []string{"aurochs"},
			want:   []string{"by", "sea", "sells", "she", "shells", "shore", "the"},
		},
		{
			name:   "key to delete is the only key in the trie",
			input:  []string{"she"},
			delete: []string{"she"},
			want:   nil,
		},
		{
			name:   "key to delete has no downstream keys",
			input:  []string{"she", "sells", "sea", "shells", "by", "the", "shore"},
			delete: []string{"shells"},
			want:   []string{"by", "sea", "sells", "she", "shore", "the"},
		},
		{
			name:   "key to delete has downstream keys",
			input:  []string{"she", "sells", "sea", "shells", "by", "the", "shore"},
			delete: []string{"she"},
			want:   []string{"by", "sea", "sells", "shells", "shore", "the"},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			trie := New[int]()
			for i, str := range tc.input {
				trie.Put(str, i)
			}
			for _, str := range tc.delete {
				trie.Delete(str)
			}

			got := trie.Keys()
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want\n\t%v\ngot\n\t%v", tc.want, got)
			}
		})
	}
}
