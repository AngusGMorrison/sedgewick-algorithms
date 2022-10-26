package rational

import (
	"errors"
	"fmt"
	"math"
)

// Rational represents a rational number.
type Rational struct {
	sign                   int64
	numerator, denominator int64
}

var (
	ErrZeroDenominator = errors.New("denominator can't be 0")
	ErrInt64Overflow   = errors.New("operation would cause uint64 overflow")
)

// New returns a new rational number in its simplest form.
func New(numerator, denominator int64) (Rational, error) {
	if denominator == 0 {
		return Rational{}, ErrZeroDenominator
	}

	var sign int64 = 1
	if numerator < 0 {
		sign = -sign
	}
	if denominator < 0 {
		sign = -sign
	}

	r := Rational{
		sign:        sign,
		numerator:   abs(numerator),
		denominator: abs(denominator),
	}

	return r.simplify(), nil
}

func (r Rational) Plus(s Rational) (Rational, error) {
	if addingCommonRationalsOverflows(r, s) {
		return Rational{}, ErrInt64Overflow
	}

	var num, den int64
	if r.denominator == s.denominator {
		num = r.SignedNumerator() + s.SignedNumerator()
		den = r.denominator
	} else {
		num = r.SignedNumerator()*s.denominator + s.SignedNumerator()*r.denominator
		den = r.denominator * s.denominator
	}

	t := Rational{
		sign:        signOf(num),
		numerator:   abs(num),
		denominator: den,
	}

	return t.simplify(), nil
}

func (r Rational) Minus(s Rational) (Rational, error) {
	s.sign = -s.sign
	return r.Plus(s)
}

func (r Rational) Times(s Rational) (Rational, error) {
	if multiplicationOverflows(r.SignedNumerator(), s.SignedNumerator()) ||
		multiplicationOverflows(r.denominator, s.denominator) {
		return Rational{}, ErrInt64Overflow
	}

	num := r.SignedNumerator() * s.SignedNumerator()
	den := r.denominator * s.denominator
	t := Rational{
		sign:        signOf(num),
		numerator:   abs(num),
		denominator: den,
	}

	return t.simplify(), nil
}

func (r Rational) DivideBy(s Rational) (Rational, error) {
	s.numerator, s.denominator = s.denominator, s.numerator
	return r.Times(s)
}

func (r Rational) Equals(s Rational) bool {
	t := r.simplify()
	u := s.simplify()

	return t.sign == u.sign &&
		t.numerator == u.numerator &&
		t.denominator == u.denominator
}

func (r Rational) String() string {
	var sign string
	if r.sign < 0 {
		sign = "-"
	}
	return fmt.Sprintf("%s%d/%d", sign, r.numerator, r.denominator)
}

func (r Rational) IsNegative() bool {
	return r.sign == -1
}

func (r Rational) SignedNumerator() int64 {
	return r.sign * r.numerator
}

func (r Rational) simplify() Rational {
	divisor := gcd(r.numerator, r.denominator)
	r.numerator /= divisor
	r.denominator /= divisor

	return r
}

func gcd(p, q int64) int64 {
	for q != 0 {
		temp := p
		p = q
		q = temp % q
	}

	return int64(math.Abs(float64(p)))
}

func abs(i int64) int64 {
	if i < 0 {
		return -i
	}

	return i
}

func signOf(i int64) int64 {
	if i < 0 {
		return -1
	}

	return 1
}

// addingCommonRationalsOverflows returns true if the process of calculating a common denominator for a and b or
// adding a and b in that form would cause integer overflow.
func addingCommonRationalsOverflows(a, b Rational) bool {
	if a.denominator == b.denominator {
		return additionOverflows(a.numerator, b.numerator)
	}

	if commonDenominatorOverflows(a, b) {
		return true
	}

	left := a.SignedNumerator() * b.denominator
	right := b.SignedNumerator() * a.denominator
	return additionOverflows(left, right)
}

func additionOverflows(a, b int64) bool {
	if b > 0 {
		// If b is positive, a must be smaller than or equal to the differnce between the maximum integer and b.
		if a > math.MaxInt64-b {
			return true
		}
	} else {
		// If b is negative, a must be smaller than the difference between the minimum integer and
		// b.
		if a < math.MinInt64-b {
			return true
		}
	}

	return false
}

// commonDenominatorOverflows returns true if any of the steps required to convert a and b to
// fractions over a common denominator would overflow.
func commonDenominatorOverflows(a, b Rational) bool {
	return multiplicationOverflows(a.SignedNumerator(), b.denominator) ||
		multiplicationOverflows(b.SignedNumerator(), a.denominator) ||
		multiplicationOverflows(a.denominator, b.denominator)
}

func multiplicationOverflows(a, b int64) bool {
	// Multiplying either a or b by 0 or 1 cannot cause an overflow.
	if a == 0 || b == 0 || a == 1 || b == 1 {
		return false
	}

	// Multiplying the minimum value by anything other than 0 or 1 will cause an overflow.
	if a == math.MinInt64 || b == math.MinInt64 {
		return true
	}

	return (a*b)/b != a
}
