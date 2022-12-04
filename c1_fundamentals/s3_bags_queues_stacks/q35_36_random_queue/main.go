package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []string{"C", "D", "H", "S"}
	deck := NewRandomQueue[string]()
	for _, v := range values {
		for _, s := range suits {
			deck.Enqueue(v + s)
		}
	}

	hands := make([][]string, 4)
	for !deck.IsEmpty() {
		hand := deck.Len() % 4
		card, err := deck.Dequeue()
		if err != nil {
			log.Fatalln(err)
		}
		hands[hand] = append(hands[hand], card)
	}

	for _, hand := range hands {
		fmt.Println(hand)
	}
}

var ErrEmpty = errors.New("queue was empty")

type RandomQueue[E any] struct {
	rng *rand.Rand
	s   []E
}

func NewRandomQueue[E any]() *RandomQueue[E] {
	src := rand.NewSource(time.Now().UnixNano())
	return &RandomQueue[E]{
		rng: rand.New(src),
	}
}

func (rq *RandomQueue[E]) Len() int {
	return len(rq.s)
}

func (rq *RandomQueue[E]) IsEmpty() bool {
	return len(rq.s) == 0
}

func (rq *RandomQueue[E]) Enqueue(elem E) {
	rq.s = append(rq.s, elem)
}

func (rq *RandomQueue[E]) Dequeue() (E, error) {
	if rq.IsEmpty() {
		return *new(E), ErrEmpty
	}

	idx := rq.rng.Intn(len(rq.s))
	elem := rq.s[idx]
	last := len(rq.s) - 1
	rq.s[last], rq.s[idx] = rq.s[idx], rq.s[last]
	rq.s = rq.s[:len(rq.s)-1]

	return elem, nil
}

func (rq *RandomQueue[E]) Sample() (E, error) {
	if rq.IsEmpty() {
		return *new(E), ErrEmpty
	}

	idx := rq.rng.Intn(len(rq.s))
	return rq.s[idx], nil
}

func (rq *RandomQueue[E]) Each(f func(elem E)) {
	rq.shuffle()
	for _, elem := range rq.s {
		f(elem)
	}
}

func (rq *RandomQueue[E]) shuffle() {
	length := len(rq.s)
	for i := 0; i < length; i++ {
		j := i + rq.rng.Intn(length-i)
		rq.s[i], rq.s[j] = rq.s[j], rq.s[i]
	}
}
