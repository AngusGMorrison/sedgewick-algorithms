package testutil

import "math/rand"

func RandomIntSlice(len int) []int {
	ints := make([]int, len)
	for i := range ints {
		ints[i] = rand.Int()
	}
	return ints
}
