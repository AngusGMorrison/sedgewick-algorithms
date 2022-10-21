package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var errUsage = errors.New("usage: bsearch path/to/whitelist")

func run() error {
	printWhitelisted := flag.Bool("printWhitelisted", false, "printWhitelisted prints all keys that appear in the whitelist")
	flag.Parse()

	whitelist, err := parseWhitelist()
	if err != nil {
		return err
	}

	if err := scanAndCompare(os.Stdin, whitelist, *printWhitelisted); err != nil {
		return err
	}

	return nil
}

func parseWhitelist() ([]int, error) {
	args := flag.Args()
	if len(args) < 1 {
		return nil, errUsage
	}

	bytes, err := ioutil.ReadFile(args[0])
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(string(bytes))

	var whitelist []int
	for _, f := range fields {
		i, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}

		whitelist = append(whitelist, i)
	}

	sort.IntSlice(whitelist).Sort()

	return whitelist, nil
}

func scanAndCompare(r io.Reader, whitelist []int, printWhitelisted bool) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		raw := scanner.Text()
		key, err := strconv.Atoi(raw)
		if err != nil {
			return err
		}

		index := rank(key, whitelist)
		if index == -1 {
			if !printWhitelisted {
				fmt.Println(key)
			}
		} else {
			if printWhitelisted {
				fmt.Println(key)
			}
		}
	}

	return nil
}

func rank(key int, a []int) int {
	if len(a) == 0 {
		return -1
	}

	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
