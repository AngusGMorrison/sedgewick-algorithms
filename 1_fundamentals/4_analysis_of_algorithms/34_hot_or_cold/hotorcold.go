package hotorcold

type solver struct {
	n         int
	target    int
	guesses   int
	lastGuess int
}

func (s *solver) Solve() int {
	s.guesses++
	mid := s.n / 2
	if mid == s.target {
		return mid
	}
	s.lastGuess = mid
	s.guesses++
	if mid-1 == s.target {
		return mid - 1
	}
	if s.warmer(mid - 1) {
		return s.solve(1, mid-2)
	}

	return s.solve(mid+1, s.n)
}

func (s *solver) solve(lo, hi int) int {
	// Our new guess is the furthest item in the current range from the previous guess, allowing us
	// to eliminate either half of the range depending on whether the guess is warmer or colder.
	s.guesses++
	var guess int
	if max(abs(lo-s.lastGuess), abs(hi-s.lastGuess)) == abs(lo-s.lastGuess) {
		guess = lo
	} else {
		guess = hi
	}
	// fmt.Println("Lo: ", lo, "Hi: ", hi, "Guess: ", guess)
	if guess == s.target {
		return guess
	}
	maxOfGuessAndLastGuess := max(guess, s.lastGuess)
	// fmt.Println("Max of guess and last guess: ", maxOfGuessAndLastGuess)
	minOfGuessAndLastGuess := min(guess, s.lastGuess)
	// fmt.Println("Min of guess and last guess: ", minOfGuessAndLastGuess)
	mid := minOfGuessAndLastGuess + (maxOfGuessAndLastGuess-minOfGuessAndLastGuess)/2
	// fmt.Println("Mid: ", mid)
	if s.warmer(guess) {
		// fmt.Println("guess is warmer")
		if guess == maxOfGuessAndLastGuess {
			// fmt.Println("guess is max of guess and last guess")
			s.lastGuess = guess
			return s.solve(mid, guess-1)
		}
		// fmt.Println("guess is min of guess and last guess")
		s.lastGuess = guess
		return s.solve(guess+1, mid)
	} else if s.colder(guess) {
		// fmt.Println("guess is colder")
		if guess == maxOfGuessAndLastGuess {
			// fmt.Println("guess is max of guess and last guess")
			s.lastGuess = guess
			return s.solve(lo, mid)
		}
		// fmt.Println("guess is min of guess and last guess")
		s.lastGuess = guess
		return s.solve(mid, hi)
	} else {
		// If we're neither warmer nor colder, the answer must lie exactly halfway between the last
		// guess and the current guess.
		// fmt.Println("guess is mid")
		return mid
	}
}

func (s *solver) warmer(guess int) bool {
	return abs(s.target-guess) < abs(s.target-s.lastGuess)
}

func (s *solver) colder(guess int) bool {
	return abs(s.target-guess) > abs(s.target-s.lastGuess)
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
