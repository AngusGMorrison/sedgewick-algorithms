package bag

import (
	"math/rand"
	"time"
)

type Bag[E any] struct {
	rng *rand.Rand
	s   []E
}

func New[E any]() *Bag[E] {
	src := rand.NewSource(time.Now().UnixNano())
	return &Bag[E]{
		rng: rand.New(src),
	}
}

func (b *Bag[E]) IsEmpty() bool {
	return len(b.s) == 0
}

func (b *Bag[E]) Len() int {
	return len(b.s)
}

func (b *Bag[E]) Add(elem E) {
	b.s = append(b.s, elem)
}

func (b *Bag[E]) Each(f func(elem E)) {
	b.shuffle()
	for _, elem := range b.s {
		f(elem)
	}
}

func (b *Bag[E]) shuffle() {
	for i := 0; i < len(b.s); i++ {
		j := i + b.rng.Intn(len(b.s)-i)
		b.s[i], b.s[j] = b.s[j], b.s[i]
	}
}
