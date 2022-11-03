package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	list := newUniqueListFromReader(os.Stdin)
	fmt.Println(list)
}

func newUniqueListFromReader(r io.Reader) *uniqueList[string] {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	list := &uniqueList[string]{}
	for scanner.Scan() {
		list.insertAtFront(scanner.Text())
	}

	return list
}

type uniqueList[D comparable] struct {
	len   int
	first *node[D]
}

func (l *uniqueList[D]) insertAtFront(data D) {
	if l.len == 0 {
		l.first = &node[D]{data: data}
		l.len++
		return
	}

	if l.first.data == data {
		return // element to insert is already the head of the list
	}

	var parent *node[D]
	for parent = l.first; parent.next != nil && parent.next.data != data; parent = parent.next {
	}
	if parent.next == nil { // element not already present
		l.first = &node[D]{
			data: data,
			next: l.first,
		}
		l.len++
		return
	}

	match := parent.next
	parent.next = parent.next.next
	match.next = l.first
	l.first = match
}

func (l *uniqueList[D]) String() string {
	if l.len == 0 {
		return "{}"
	}

	var builder strings.Builder
	_, _ = fmt.Fprintf(&builder, "%v", l.first.data)
	for cur := l.first.next; cur != nil; cur = cur.next {
		_, _ = fmt.Fprintf(&builder, " -> %v", cur.data)
	}

	return builder.String()
}

type node[D comparable] struct {
	data D
	next *node[D]
}
