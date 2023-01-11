package testutil

import (
	"math/rand"
	"testing"
	"time"
)

// Rand extends *rand.Rand with randomizing test utilities.
type Rand struct {
	*rand.Rand
}

// RandOption can be passed to Rand's constructor to override its default behavior.
type RandOption interface {
	apply(opts *randOptions)
}

type randOptions struct {
	seed int64
}

// RandOptionSeed instnatiates the randomizer with the given seed.
type RandOptionSeed struct {
	seed int64
}

func (r RandOptionSeed) apply(opts *randOptions) {
	opts.seed = r.seed
}

// NewRand instantiates a Rand with the given options and logs the seed using the test logger to
// support debugging.
func NewRand(t *testing.T, opts ...RandOption) *Rand {
	t.Helper()

	options := randOptions{
		seed: time.Now().UnixNano(),
	}
	for _, opt := range opts {
		opt.apply(&options)
	}

	src := rand.NewSource(options.seed)
	rng := rand.New(src)
	t.Logf("randomizing with seed %d", options.seed)

	return &Rand{rng}
}

// IntSlice returns a slice with the given length filled with random integers.
func (r *Rand) IntSlice(length int) []int {
	ints := make([]int, length)
	for i := range ints {
		ints[i] = r.Int()
	}
	return ints
}

const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// ASCIIAlphanumString returns a string of the given length constructed from randomly-chosen
// lowercase and uppercase ASCII letters and the digits 0-9.
func (r *Rand) ASCIIAlphanumString(length int) string {
	str := make([]byte, length)
	for i := range str {
		l := r.Intn(len(letters))
		str[i] = letters[l]
	}
	return string(str)
}
