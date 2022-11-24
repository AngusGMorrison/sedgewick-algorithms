/*
This is an insertion sort problem. If you had a hand of cards, you would pick a card
that was not in sorted order and insert it into the correct position. That is, you move the card
around the hand. However, when we have a deck of cards and we can only see the top two, we must move
the deck around the out-of-place card.

When the top card is less than the second card, we consider the top card to be out of place. We move
it to the correct place by flowing the rest of the deck around it, moving the second card to the
bottom of the deck, drawing the next, moving it to the bottom and so on, until we find a card that
is less than the top card, at which point we know the top card is in the correct place relative to
its neighbors and move it to the bottom. The former second card then becomes the new top card.

Since the deck is cyclical (we add at least one card to the bottom of the deck on every comparison),
we must also track whether we have compared the top card to every other card in the deck. If so, we
know it is in the correct place relative to all other cards, and we can move it to the bottom.
*/
package dequesort

import (
	"math/rand"
)

var (
	_suits = []string{"C", "D", "H", "S"}
	_ranks = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
)

type Card struct {
	suit string
	rank int
}

func (c Card) Less(other Card) bool {
	if c.suit < other.suit {
		return true
	}
	if c.suit > other.suit {
		return false
	}
	return c.rank < other.rank
}

type Deck struct {
	Deque[Card]
	rng *rand.Rand
}

func NewDeck(src rand.Source) *Deck {
	deck := &Deck{
		Deque: Deque[Card]{},
		rng:   rand.New(src),
	}
	for _, s := range _suits {
		for _, r := range _ranks {
			deck.Add(Card{suit: s, rank: r})
		}
	}
	deck.Shuffle()

	return deck
}

func (d *Deck) Sort() {
	var (
		cur             Card
		compares        int           // tracks how many cards the current card has been compared to, up to nCards-1
		sortedSubseqLen int           // tracks the length of the current subsequence that we know to be sorted
		nCards          int = d.Len() // d.Len() fluctuates as cards are popped and pushed, so we cache the initial value
	)

	// Loop until the whole deck is one sorted subsequence. We loop to nCards-1, not nCards, because the final
	// card is added without incrementing sortedSubseqLen.
	for sortedSubseqLen < nCards-1 {
		// Loop until the current card has been compared to every other card, or the whole deck is
		// sorted pending the addition of the current card to the bottom of the deck.
		for cur, _ = d.PopRight(); compares < d.Len() && sortedSubseqLen < nCards-1; cur, _ = d.PopRight() {
			next, _ := d.PopRight()
			compares++
			if cur.Less(next) { // cur is out of order. Move next to the bottom and put cur back on top (like an insertion sort swap)
				d.PushRight(cur)
				d.PushLeft(next)
				sortedSubseqLen = 0
			} else { // cur is in the correct order relative to next. Move cur to the bottom and put next on top
				d.PushRight(next)
				d.PushLeft(cur)
				sortedSubseqLen++
				compares = 0
			}
		}

		// cur is in the correct position relative to every other card in the deck. Move it to the
		// bottom.
		d.PushLeft(cur)
		compares = 0
	}
}

func (d *Deck) Sorted() bool {
	cur, ok := d.PopLeft()
	if !ok {
		return true // deck is empty
	}

	for next, ok := d.PopLeft(); ok; next, ok = d.PopLeft() {
		if !cur.Less(next) {
			return false
		}
		cur = next
	}
	return true
}

func (d *Deck) Shuffle() {
	d.rng.Shuffle(d.Len(), func(i, j int) {
		d.data[i], d.data[j] = d.data[j], d.data[i]
	})
}

func (d *Deck) Add(card Card) {
	d.PushLeft(card)
}

type Deque[D any] struct {
	data []D
}

func (sd *Deque[D]) Len() int {
	return len(sd.data)
}

func (sd *Deque[D]) PushLeft(data D) {
	sd.data = append(sd.data, *new(D))
	copy(sd.data[1:], sd.data)
	sd.data[0] = data
}

func (sd *Deque[D]) PushRight(data D) {
	sd.data = append(sd.data, data)
}

func (sd *Deque[D]) PopLeft() (D, bool) {
	if len(sd.data) == 0 {
		return *new(D), false
	}

	data := sd.data[0]
	sd.data = sd.data[1:]

	return data, true
}

func (sd *Deque[D]) PopRight() (D, bool) {
	if len(sd.data) == 0 {
		return *new(D), false
	}

	data := sd.data[len(sd.data)-1]
	sd.data = sd.data[:len(sd.data)-1]

	return data, true
}
