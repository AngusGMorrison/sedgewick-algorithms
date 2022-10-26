package main

import "fmt"

func main() {
	a := []int{10, 11, 12, 16, 18, 23, 29, 33, 48, 54, 57, 68, 77, 84, 98}
	c := counter{}
	rank(11, a, &c)
	fmt.Println(c.count)
}

type counter struct {
	count int
}

func (c *counter) inc() {
	c.count++
}

func rank(key int, a []int, c *counter) int {
	lo := 0
	hi := len(a)
	for lo <= hi {
		c.inc()
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
