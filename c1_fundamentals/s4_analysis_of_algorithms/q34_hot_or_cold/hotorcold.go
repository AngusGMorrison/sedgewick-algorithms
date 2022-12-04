package q34_hot_or_cold

type solver struct {
	n       int
	target  int
	guesses int
}

// Solve guesses target in loga(n)+b guesses, where a is somewhere between 1 and 2. It recursively
// eliminates half the range at each guess, by testing whether the element furthest from the
// previous guess is warmer or colder than the previous guess. If warmer, we know that the target is
// in the half of the range closest to the new guess. If colder, we know it's in the half of the
// range closest to the previous guess.
func (s *solver) Solve() int {
	// The initial guess is the middle of the range. Gives us no warmer or colder information since
	// there is no previous guess.
	mid := s.n / 2
	s.guesses++
	if mid == s.target {
		return mid
	}

	// The next guess is immediately next to the middle of the range. The choice of side to test is
	// arbitrary. This tells us which half of the range the target is in.
	s.guesses++
	if mid-1 == s.target {
		return mid - 1
	}
	if s.warmer(mid-1, mid) {
		// If the guess was closer to the target, solve in the range between 1 and the middle of the
		// range, excluding the elements already tested.
		return s.solve(1, mid-2, mid)
	}
	// Our guess was further from the target, so solve between the middle and the end of the range,
	// excluding the middle, which has already been tested.
	return s.solve(mid+1, s.n, mid)
}

func (s *solver) solve(lo, hi, lastGuess int) int {
	// Our new guess is the furthest item in the current range from the previous guess, allowing us
	// to eliminate either half of the range depending on whether the guess is warmer or colder.
	guess := furthestFromLastGuess(lo, hi, lastGuess)
	s.guesses++
	if guess == s.target {
		return guess
	}

	min := min(guess, lastGuess)
	max := max(guess, lastGuess)
	mid := min + (max-min)/2

	if s.warmer(guess, lastGuess) { // target is in half of range closest to guess
		if guess == max { // shrink the range based on whether we're at the low or high end of it
			return s.solve(mid, guess-1, guess)
		}
		return s.solve(guess+1, mid, guess)
	}

	if s.colder(guess, lastGuess) { // target is in half of range furthest from guess
		if guess == max {
			return s.solve(lo, mid, guess)
		}
		return s.solve(mid, hi, guess)
	}

	// If we're neither warmer nor colder, the answer must lie exactly halfway between the last
	// guess and the current guess.
	return mid
}

func (s *solver) warmer(guess, lastGuess int) bool {
	return abs(s.target-guess) < abs(s.target-lastGuess)
}

func (s *solver) colder(guess, lastGuess int) bool {
	return abs(s.target-guess) > abs(s.target-lastGuess)
}

func furthestFromLastGuess(lo, hi, lastGuess int) int {
	loDiff := abs(lo - lastGuess)
	if max(loDiff, abs(hi-lastGuess)) == loDiff {
		return lo
	} else {
		return hi
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
