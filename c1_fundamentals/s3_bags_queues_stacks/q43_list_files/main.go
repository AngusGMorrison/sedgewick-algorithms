package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/queue"
)

func main() {
	q, err := enqueueFiles("../..")
	if err != nil {
		log.Fatalln(err)
	}

	var builder strings.Builder
	q.Each(func(elem string) {
		_, _ = fmt.Fprintf(&builder, "%s\n", elem)
	})
	fmt.Println(builder.String())
}

func enqueueFiles(dir string) (*queue.SliceQueue[string], error) {
	q := queue.NewSliceQueue[string]()
	q.Enqueue(dir) // enqueue root directory

	var enqueueDirectory func(path string, depth int) error
	enqueueDirectory = func(path string, depth int) error {
		dirEntries, err := os.ReadDir(path)
		if err != nil {
			return err
		}
		nextDepth := depth + 1
		for _, dirEntry := range dirEntries {
			q.Enqueue(fmt.Sprintf("%s%s", strings.Repeat("\t", depth), dirEntry.Name()))
			if dirEntry.IsDir() {
				if err := enqueueDirectory(filepath.Join(path, dirEntry.Name()), nextDepth); err != nil {
					return err
				}
			}
		}

		return nil
	}

	if err := enqueueDirectory(dir, 1); err != nil {
		return nil, err
	}

	return q, nil
}
