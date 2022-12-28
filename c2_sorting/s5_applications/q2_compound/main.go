package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	words := readWords(os.Stdin)
	writeCompoundWords(os.Stdout, words)
}

func readWords(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words
}

func writeCompoundWords(w io.Writer, words []string) {
	// By sorting, we guarantee that words with the same prefix are grouped together.
	sort.Strings(words)

	for i := 1; i < len(words); i++ { // while there is at least one pair of words remaining...
		for prefix, next := words[i-1], words[i]; strings.HasPrefix(next, prefix); next = words[i] { // for each word that starts with the current prefix...
			// Use binary search to check if the second half of the word exists.
			suffix := next[:len(prefix)]
			if _, ok := slices.BinarySearch(words, suffix); ok {
				fmt.Fprintln(w, next)
			}

			i++
			if i >= len(words) {
				break
			}
		}
	}
}
