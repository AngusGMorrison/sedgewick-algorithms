package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	n := flag.Int("n", 5, "Generate an nxn grid.")
	flag.Parse()

	conns := generate(*n)
	for conns.len() > 0 {
		fmt.Println(conns.sample())
	}
}

type connections struct {
	data []connection
}

// sample returns a random connection, removing it from connections.
func (c *connections) sample() connection {
	idx := rand.Intn(len(c.data))
	conn := c.data[idx]
	c.data[idx], c.data[len(c.data)-1] = c.data[len(c.data)-1], c.data[idx]
	c.data = c.data[:len(c.data)-1]
	return conn
}

func (c *connections) add(conn connection) {
	c.data = append(c.data, conn)
}

func (c *connections) len() int {
	return len(c.data)
}

type connection struct {
	p, q int
}

func generate(n int) connections {
	conns := connections{data: make([]connection, 0, n*n)}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			conn := connection{p: i, q: j}
			if rand.Int()%2 == 0 { // randomly orient the connection
				conn.p, conn.q = conn.q, conn.p
			}
			conns.add(conn)
		}
	}

	return conns
}
