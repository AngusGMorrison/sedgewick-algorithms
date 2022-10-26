package smartdate

import (
	"errors"
	"testing"
)

func Test_New(t *testing.T) {
	t.Parallel()

	t.Run("valid inputs", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name             string
			day, month, year uint
			want             Date
		}{
			{
				name:  "first day of any month",
				day:   1,
				month: 1,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(1),
					Day:   1,
				},
			},
			{
				name:  "any day of any month",
				day:   13,
				month: 7,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(7),
					Day:   13,
				},
			},
			{
				name:  "last day of January",
				day:   31,
				month: 1,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(1),
					Day:   31,
				},
			},
			{
				name:  "last day of February (non-leap year)",
				day:   28,
				month: 2,
				year:  1,
				want: Date{
					Year:  1,
					Month: Month(2),
					Day:   28,
				},
			},
			{
				name:  "last day of February (leap year)",
				day:   29,
				month: 2,
				year:  1992,
				want: Date{
					Year:  1992,
					Month: Month(2),
					Day:   29,
				},
			},
			{
				name:  "last day of March",
				day:   31,
				month: 3,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(3),
					Day:   31,
				},
			},
			{
				name:  "last day of April",
				day:   30,
				month: 4,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(4),
					Day:   30,
				},
			},
			{
				name:  "last day of May",
				day:   31,
				month: 5,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(5),
					Day:   31,
				},
			},
			{
				name:  "last day of June",
				day:   30,
				month: 6,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(6),
					Day:   30,
				},
			},
			{
				name:  "last day of July",
				day:   31,
				month: 7,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(7),
					Day:   31,
				},
			},
			{
				name:  "last day of August",
				day:   31,
				month: 8,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(8),
					Day:   31,
				},
			},
			{
				name:  "last day of September",
				day:   30,
				month: 9,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(9),
					Day:   30,
				},
			},
			{
				name:  "last day of October",
				day:   31,
				month: 10,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(10),
					Day:   31,
				},
			},
			{
				name:  "last day of November",
				day:   30,
				month: 11,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(11),
					Day:   30,
				},
			},
			{
				name:  "last day of December",
				day:   31,
				month: 12,
				year:  0,
				want: Date{
					Year:  0,
					Month: Month(12),
					Day:   31,
				},
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				gotDate, gotErr := New(tc.year, tc.month, tc.day)
				if gotErr != nil {
					t.Errorf("want not error, got %q", gotErr)
				}

				if gotDate != tc.want {
					t.Errorf("want %+v, got %+v", tc.want, gotDate)
				}
			})
		}
	})

	t.Run("invalid inputs", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name             string
			day, month, year uint
		}{
			{
				name:  "zeroth month",
				day:   1,
				month: 0,
				year:  0,
			},
			{
				name:  "zeroth day of any month",
				day:   0,
				month: 1,
				year:  0,
			},
			{
				name:  "after last day of January",
				day:   32,
				month: 1,
				year:  0,
			},
			{
				name:  "after last day of February (non-leap year)",
				day:   29,
				month: 2,
				year:  1,
			},
			{
				name:  "after last day of February (leap year)",
				day:   30,
				month: 2,
				year:  1992,
			},
			{
				name:  "after last day of March",
				day:   32,
				month: 3,
				year:  0,
			},
			{
				name:  "after last day of April",
				day:   31,
				month: 4,
				year:  0,
			},
			{
				name:  "after last day of May",
				day:   32,
				month: 5,
				year:  0,
			},
			{
				name:  "after last day of June",
				day:   31,
				month: 6,
				year:  0,
			},
			{
				name:  "after last day of July",
				day:   32,
				month: 7,
				year:  0,
			},
			{
				name:  "after last day of August",
				day:   32,
				month: 8,
				year:  0,
			},
			{
				name:  "after last day of September",
				day:   31,
				month: 9,
			},
			{
				name:  "after last day of October",
				day:   32,
				month: 10,
				year:  0,
			},
			{
				name:  "after last day of November",
				day:   31,
				month: 11,
				year:  0,
			},
			{
				name:  "after last day of December",
				day:   32,
				month: 12,
				year:  0,
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				gotDate, gotErr := New(tc.year, tc.month, tc.day)
				if !errors.Is(gotErr, ErrInvalidDate) {
					t.Errorf("want ErrInvalidDate, got %q", gotErr)
				}

				if (gotDate != Date{}) {
					t.Errorf("want empty Date, got %+v", gotDate)
				}
			})
		}
	})
}

func Test_Date_dayOfWeek(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		date Date
		want Day
	}{
		{
			name: "Sunday",
			date: Date{
				Year:  2022,
				Month: October,
				Day:   23,
			},
			want: Sunday,
		},
		{
			name: "Monday",
			date: Date{
				Year:  2022,
				Month: October,
				Day:   24,
			},
			want: Monday,
		},
		{
			name: "Tuesday",
			date: Date{
				Year:  2022,
				Month: October,
				Day:   25,
			},
			want: Tuesday,
		},
		{
			name: "Wednesday",
			date: Date{
				Year:  2022,
				Month: October,
				Day:   26,
			},
			want: Wednesday,
		},
		{
			name: "Thursday",
			date: Date{
				Year:  2022,
				Month: October,
				Day:   27,
			},
			want: Thursday,
		},
		{
			name: "Friday",
			date: Date{
				Year:  2022,
				Month: October,
				Day:   28,
			},
			want: Friday,
		},
		{
			name: "Saturday",
			date: Date{
				Year:  2022,
				Month: October,
				Day:   29,
			},
			want: Saturday,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.date.DayOfWeek(); got != tc.want {
				t.Errorf("want %d (%s), got %d (%s)", tc.want, tc.want, got, got)
			}
		})
	}
}
