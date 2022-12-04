package q16_17_rational

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func Test_New(t *testing.T) {
	t.Parallel()

	t.Run("denominator is non-zero", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name        string
			numerator   int64
			denominator int64
			want        Rational
		}{
			{
				name:        "numerator is zero",
				numerator:   0,
				denominator: 2,
				want: Rational{
					sign:        1,
					numerator:   0,
					denominator: 1, // simplified form
				},
			},
			{
				name:        "numerator and denominator are positive",
				numerator:   1,
				denominator: 2,
				want: Rational{
					sign:        1,
					numerator:   1,
					denominator: 2,
				},
			},
			{
				name:        "numerator and denominator are negative",
				numerator:   -1,
				denominator: -2,
				want: Rational{
					sign:        1,
					numerator:   1,
					denominator: 2,
				},
			},
			{
				name:        "numerator is positive and denominator is negative",
				numerator:   1,
				denominator: -2,
				want: Rational{
					sign:        -1,
					numerator:   1,
					denominator: 2,
				},
			},
			{
				name:        "numerator is negative and denominator is positive",
				numerator:   -1,
				denominator: 2,
				want: Rational{
					sign:        -1,
					numerator:   1,
					denominator: 2,
				},
			},
			{
				name:        "input can be simplified",
				numerator:   2,
				denominator: 4,
				want: Rational{
					sign:        1,
					numerator:   1,
					denominator: 2,
				},
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				gotRational, gotErr := New(tc.numerator, tc.denominator)
				if gotErr != nil {
					t.Errorf("want no error, got %v", gotErr)
				}
				if gotRational != tc.want {
					t.Errorf("want %+v, got %+v", tc.want, gotRational)
				}
			})
		}
	})

	t.Run("denominator is zero", func(t *testing.T) {
		t.Parallel()

		gotRational, gotErr := New(2, 0)
		if !errors.Is(gotErr, ErrZeroDenominator) {
			t.Errorf("want ErrZeroDenominator, got %v", gotErr)
		}
		if (gotRational != Rational{}) {
			t.Errorf("want empty Rational, got %+v", gotRational)
		}
	})
}

func Test_Rational_Plus(t *testing.T) {
	t.Parallel()

	t.Run("addition succeeds", func(t *testing.T) {
		testCases := []struct {
			name string
			a, b Rational
			want Rational
		}{
			{
				name: "operands are both positive",
				a:    Rational{sign: 1, numerator: 5, denominator: 6},
				b:    Rational{sign: 1, numerator: 3, denominator: 8},
				want: Rational{sign: 1, numerator: 29, denominator: 24},
			},
			{
				name: "one operand is negative",
				a:    Rational{sign: -1, numerator: 3, denominator: 8},
				b:    Rational{sign: 1, numerator: 1, denominator: 9},
				want: Rational{sign: -1, numerator: 19, denominator: 72},
			},
			{
				name: "both operands are negative",
				a:    Rational{sign: -1, numerator: 3, denominator: 8},
				b:    Rational{sign: -1, numerator: 5, denominator: 6},
				want: Rational{sign: -1, numerator: 29, denominator: 24},
			},
			{
				name: "sign change",
				a:    Rational{sign: -1, numerator: 3, denominator: 8},
				b:    Rational{sign: 1, numerator: 5, denominator: 6},
				want: Rational{sign: 1, numerator: 11, denominator: 24},
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				gotRational, gotErr := tc.a.Plus(tc.b)
				if gotErr != nil {
					t.Errorf("want no error, got %v", gotErr)
				}
				if gotRational != tc.want {
					t.Errorf("want %+v, got %+v", tc.want, gotRational)
				}
			})
		}
	})

	t.Run("addition overflows", func(t *testing.T) {
		t.Parallel()

		a := Rational{sign: 1, numerator: math.MaxInt64, denominator: 1}
		b := Rational{sign: 1, numerator: 1, denominator: 1}

		gotRational, gotErr := a.Plus(b)
		if !errors.Is(gotErr, ErrInt64Overflow) {
			t.Errorf("want ErrInt64Overflow, got %v", gotErr)
		}
		if (gotRational != Rational{}) {
			t.Errorf("want empty Rational, got %+v", gotRational)
		}
	})
}

func Test_Rational_Minus(t *testing.T) {
	t.Parallel()

	t.Run("subtraction succeeds", func(t *testing.T) {
		testCases := []struct {
			name string
			a, b Rational
			want Rational
		}{
			{
				name: "operands are both positive",
				a:    Rational{sign: 1, numerator: 5, denominator: 6},
				b:    Rational{sign: 1, numerator: 3, denominator: 8},
				want: Rational{sign: 1, numerator: 11, denominator: 24},
			},
			{
				name: "one operand is negative",
				a:    Rational{sign: -1, numerator: 3, denominator: 8},
				b:    Rational{sign: 1, numerator: 1, denominator: 9},
				want: Rational{sign: -1, numerator: 35, denominator: 72},
			},
			{
				name: "both operands are negative",
				a:    Rational{sign: -1, numerator: 5, denominator: 6},
				b:    Rational{sign: -1, numerator: 3, denominator: 8},
				want: Rational{sign: -1, numerator: 11, denominator: 24},
			},
			{
				name: "sign change",
				a:    Rational{sign: -1, numerator: 3, denominator: 8},
				b:    Rational{sign: -1, numerator: 5, denominator: 6},
				want: Rational{sign: 1, numerator: 11, denominator: 24},
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				gotRational, gotErr := tc.a.Minus(tc.b)
				if gotErr != nil {
					t.Errorf("want no error, got %v", gotErr)
				}
				if gotRational != tc.want {
					t.Errorf("want %+v, got %+v", tc.want, gotRational)
				}
			})
		}
	})

	t.Run("subtraction overflows", func(t *testing.T) {
		t.Parallel()

		a := Rational{sign: 1, numerator: math.MinInt64, denominator: 1}
		b := Rational{sign: 1, numerator: 1, denominator: 1}

		gotRational, gotErr := a.Minus(b)
		if !errors.Is(gotErr, ErrInt64Overflow) {
			t.Errorf("want ErrInt64Overflow, got %v", gotErr)
		}
		if (gotRational != Rational{}) {
			t.Errorf("want empty Rational, got %+v", gotRational)
		}
	})
}

func Test_Rational_Times(t *testing.T) {
	t.Parallel()

	t.Run("multiplication succeeds", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name string
			a, b Rational
			want Rational
		}{
			{
				name: "operands are both positive",
				a:    Rational{sign: 1, numerator: 5, denominator: 6},
				b:    Rational{sign: 1, numerator: 3, denominator: 8},
				want: Rational{sign: 1, numerator: 5, denominator: 16},
			},
			{
				name: "one operand is negative",
				a:    Rational{sign: -1, numerator: 5, denominator: 6},
				b:    Rational{sign: 1, numerator: 3, denominator: 8},
				want: Rational{sign: -1, numerator: 5, denominator: 16},
			},
			{
				name: "both operands are negative",
				a:    Rational{sign: -1, numerator: 5, denominator: 6},
				b:    Rational{sign: -1, numerator: 3, denominator: 8},
				want: Rational{sign: 1, numerator: 5, denominator: 16},
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				gotRational, gotErr := tc.a.Times(tc.b)
				if gotErr != nil {
					t.Errorf("want no error, got %v", gotErr)
				}
				if gotRational != tc.want {
					t.Errorf("want %+v, got %+v", tc.want, gotRational)
				}
			})
		}
	})

	t.Run("mulitplication overflows", func(t *testing.T) {
		t.Parallel()

		a := Rational{sign: 1, numerator: math.MinInt64, denominator: 1}
		b := Rational{sign: -1, numerator: 1, denominator: 1}

		gotRational, gotErr := a.Times(b)
		if !errors.Is(gotErr, ErrInt64Overflow) {
			t.Errorf("want ErrInt64Overflow, got %v", gotErr)
		}
		if (gotRational != Rational{}) {
			t.Errorf("want empty Rational, got %+v", gotRational)
		}
	})
}

func Test_Rational_DivideBy(t *testing.T) {
	t.Parallel()

	t.Run("division succeeds", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name string
			a, b Rational
			want Rational
		}{
			{
				name: "operands are both positive",
				a:    Rational{sign: 1, numerator: 5, denominator: 6},
				b:    Rational{sign: 1, numerator: 3, denominator: 8},
				want: Rational{sign: 1, numerator: 20, denominator: 9},
			},
			{
				name: "one operand is negative",
				a:    Rational{sign: -1, numerator: 5, denominator: 6},
				b:    Rational{sign: 1, numerator: 3, denominator: 8},
				want: Rational{sign: -1, numerator: 20, denominator: 9},
			},
			{
				name: "both operands are negative",
				a:    Rational{sign: -1, numerator: 5, denominator: 6},
				b:    Rational{sign: -1, numerator: 3, denominator: 8},
				want: Rational{sign: 1, numerator: 20, denominator: 9},
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				gotRational, gotErr := tc.a.DivideBy(tc.b)
				if gotErr != nil {
					t.Errorf("want no error, got %v", gotErr)
				}
				if gotRational != tc.want {
					t.Errorf("want %+v, got %+v", tc.want, gotRational)
				}
			})
		}
	})

	t.Run("division overflows", func(t *testing.T) {
		t.Parallel()

		a := Rational{sign: 1, numerator: math.MinInt64, denominator: 1}
		b := Rational{sign: -1, numerator: 1, denominator: 1}

		gotRational, gotErr := a.DivideBy(b)
		if !errors.Is(gotErr, ErrInt64Overflow) {
			t.Errorf("want ErrInt64Overflow, got %v", gotErr)
		}
		if (gotRational != Rational{}) {
			t.Errorf("want empty Rational, got %+v", gotRational)
		}
	})
}

func Test_Rational_simplify(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input Rational
		want  Rational
	}{
		{
			name:  "input is in its simplest form",
			input: Rational{sign: 1, numerator: 1, denominator: 3},
			want:  Rational{sign: 1, numerator: 1, denominator: 3},
		},
		{
			name:  "input can be simplified",
			input: Rational{sign: 1, numerator: 70, denominator: 1400},
			want:  Rational{sign: 1, numerator: 1, denominator: 20},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.input.simplify(); got != tc.want {
				t.Errorf("want %+v, got %+v", tc.want, got)
			}
		})
	}
}

func Test_Rational_Equals(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		a, b Rational
		want bool
	}{
		{
			name: "operands are identical",
			a:    Rational{sign: 1, numerator: 1, denominator: 3},
			b:    Rational{sign: 1, numerator: 1, denominator: 3},
			want: true,
		},
		{
			name: "operands are identical when simplified",
			a:    Rational{sign: 1, numerator: 1, denominator: 3},
			b:    Rational{sign: 1, numerator: 3, denominator: 9},
			want: true,
		},
		{
			name: "operands are unequal",
			a:    Rational{sign: 1, numerator: 1, denominator: 3},
			b:    Rational{sign: 1, numerator: 2, denominator: 3},
			want: false,
		},
		{
			name: "operands differ in sign only",
			a:    Rational{sign: 1, numerator: 1, denominator: 3},
			b:    Rational{sign: -1, numerator: 1, denominator: 3},
			want: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.a.Equals(tc.b); got != tc.want {
				t.Errorf("(%s).Equals(%s): want %t, got %t", tc.a, tc.b, tc.want, got)
			}
		})
	}
}

func Test_Rational_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		r    Rational
		want string
	}{
		{
			name: "Rational is positive",
			r:    Rational{sign: 1, numerator: 1, denominator: 3},
			want: "1/3",
		},
		{
			name: "Rational is negative",
			r:    Rational{sign: -1, numerator: 1, denominator: 3},
			want: "-1/3",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.r.String(); got != tc.want {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}

func Test_Rational_IsNegative(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		r    Rational
		want bool
	}{
		{
			name: "Rational is negative",
			r:    Rational{sign: -1, numerator: 1, denominator: 1},
			want: true,
		},
		{
			name: "Rational is positive",
			r:    Rational{sign: 1, numerator: 1, denominator: 1},
			want: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.r.IsNegative(); got != tc.want {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		p, q int64
		want int64
	}{
		{
			name: "p and q are both > 0",
			p:    20,
			q:    8,
			want: 4,
		},
		{
			name: "p is zero",
			p:    0,
			q:    8,
			want: 8,
		},
		{
			name: "q is zero",
			p:    8,
			q:    0,
			want: 8,
		},
		{
			name: "p and q are zero",
			p:    0,
			q:    0,
			want: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := gcd(tc.p, tc.q); got != tc.want {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func Test_additionOverflows(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b int64
		want bool
	}{
		{
			a:    math.MaxInt64,
			b:    1,
			want: true,
		},
		{
			a:    1,
			b:    math.MaxInt64,
			want: true,
		},
		{
			a:    math.MinInt64,
			b:    -1,
			want: true,
		},
		{
			a:    -1,
			b:    math.MinInt64,
			want: true,
		},
		{
			a:    0,
			b:    math.MinInt64,
			want: false,
		},
		{
			a:    math.MinInt64,
			b:    0,
			want: false,
		},
		{
			a:    0,
			b:    math.MaxInt64,
			want: false,
		},
		{
			a:    math.MaxInt64,
			b:    0,
			want: false,
		},
		{
			a:    1,
			b:    2,
			want: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("additionOverflows(%d, %d)", tc.a, tc.b), func(t *testing.T) {
			t.Parallel()

			if got := additionOverflows(tc.a, tc.b); got != tc.want {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}
}

func Test_multiplicationOverflows(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b int64
		want bool
	}{
		{
			a:    math.MaxInt64,
			b:    0,
			want: false,
		},
		{
			a:    0,
			b:    math.MaxInt64,
			want: false,
		},
		{
			a:    math.MinInt64,
			b:    0,
			want: false,
		},
		{
			a:    0,
			b:    math.MinInt64,
			want: false,
		},
		{
			a:    math.MaxInt64,
			b:    1,
			want: false,
		},
		{
			a:    1,
			b:    math.MaxInt64,
			want: false,
		},
		{
			a:    math.MinInt64,
			b:    1,
			want: false,
		},
		{
			a:    1,
			b:    math.MinInt64,
			want: false,
		},
		{
			a:    2,
			b:    math.MaxInt64,
			want: true,
		},
		{
			a:    math.MaxInt64,
			b:    2,
			want: true,
		},
		{
			a:    2,
			b:    math.MinInt64,
			want: true,
		},
		{
			a:    math.MinInt64,
			b:    2,
			want: true,
		},
		{
			a:    3,
			b:    math.MinInt64 / 2,
			want: true,
		},
		{
			a:    math.MinInt64 / 2,
			b:    3,
			want: true,
		},
		{
			a:    3,
			b:    math.MaxInt64 / 2,
			want: true,
		},
		{
			a:    math.MaxInt64 / 2,
			b:    3,
			want: true,
		},
		{
			a:    2,
			b:    (math.MinInt64 / 2) - 1,
			want: true,
		},
		{
			a:    (math.MinInt64 / 2) - 1,
			b:    2,
			want: true,
		},
		{
			a:    2,
			b:    (math.MaxInt64 / 2) + 1,
			want: true,
		},
		{
			a:    (math.MaxInt64 / 2) + 1,
			b:    2,
			want: true,
		},
		{
			a:    -2,
			b:    (math.MinInt64 / 2),
			want: true,
		},
		{
			a:    (math.MinInt64 / 2),
			b:    -2,
			want: true,
		},
		{
			a:    -2,
			b:    (math.MaxInt64 / 2),
			want: false,
		},
		{
			a:    (math.MaxInt64 / 2),
			b:    -2,
			want: false,
		},
		{
			a:    5,
			b:    2,
			want: false,
		},
		{
			a:    2,
			b:    5,
			want: false,
		},
		{
			a:    -2,
			b:    5,
			want: false,
		},
		{
			a:    5,
			b:    -2,
			want: false,
		},
		{
			a:    2,
			b:    -5,
			want: false,
		},
		{
			a:    -5,
			b:    2,
			want: false,
		},
		{
			a:    -5,
			b:    -2,
			want: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("multiplicationOverflows(%d, %d)", tc.a, tc.b), func(t *testing.T) {
			t.Parallel()

			if got := multiplicationOverflows(tc.a, tc.b); got != tc.want {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}
}

func Test_commonDenominatorOverflows(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		a, b Rational
		want bool
	}{
		{
			name: "a.numerator * b.denominator overflows",
			a: Rational{
				sign:        1,
				numerator:   3,
				denominator: 1,
			},
			b: Rational{
				sign:        1,
				numerator:   1,
				denominator: math.MaxInt64 / 2,
			},
			want: true,
		},
		{
			name: "-(a.numerator) * b.denominator overflows",
			a: Rational{
				sign:        -1,
				numerator:   3,
				denominator: 1,
			},
			b: Rational{
				sign:        1,
				numerator:   1,
				denominator: math.MaxInt64 / 2,
			},
			want: true,
		},
		{
			name: "b.numerator * a.denominator overflows",
			a: Rational{
				sign:        1,
				numerator:   1,
				denominator: math.MaxInt64 / 2,
			},
			b: Rational{
				sign:        1,
				numerator:   3,
				denominator: 1,
			},
			want: true,
		},
		{
			name: "-(b.numerator) * a.denominator overflows",
			a: Rational{
				sign:        1,
				numerator:   1,
				denominator: math.MaxInt64 / 2,
			},
			b: Rational{
				sign:        -1,
				numerator:   3,
				denominator: 1,
			},
			want: true,
		},
		{
			name: "a.denominator * b.denominator overflows",
			a: Rational{
				sign:        1,
				numerator:   1,
				denominator: math.MaxInt64 / 2,
			},
			b: Rational{
				sign:        1,
				numerator:   1,
				denominator: 3,
			},
			want: true,
		},
		{
			name: "no overflow occurs",
			a: Rational{
				sign:        1,
				numerator:   1,
				denominator: 2,
			},
			b: Rational{
				sign:        1,
				numerator:   1,
				denominator: 3,
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprintf("commonDenominatorOverflows(%d, %d)", tc.a, tc.b), func(t *testing.T) {
			t.Parallel()

			if got := commonDenominatorOverflows(tc.a, tc.b); got != tc.want {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}
}

func Test_addingCommonRationalsOverflows(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		a, b Rational
		want bool
	}{
		{
			name: "common denominator overflows",
			a: Rational{
				sign:        1,
				numerator:   3,
				denominator: 1,
			},
			b: Rational{
				sign:        1,
				numerator:   1,
				denominator: math.MaxInt64 / 2,
			},
			want: true,
		},
		{
			name: "denominators are equal and adding common fractions overflows",
			a: Rational{
				sign:        1,
				numerator:   math.MaxInt64,
				denominator: 1,
			},
			b: Rational{
				sign:        1,
				numerator:   1,
				denominator: 1,
			},
			want: true,
		},
		{
			name: "denominators are different and adding common fractions overflows",
			a: Rational{
				sign:        1,
				numerator:   math.MaxInt64 / 3,
				denominator: 1,
			},
			b: Rational{
				sign:        1,
				numerator:   2 * math.MaxInt64 / 3,
				denominator: 2,
			},
			want: true,
		},
		{
			name: "no overflow",
			a: Rational{
				sign:        1,
				numerator:   1,
				denominator: 3,
			},
			b: Rational{
				sign:        1,
				numerator:   1,
				denominator: 2,
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := addingCommonRationalsOverflows(tc.a, tc.b); got != tc.want {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}
}
