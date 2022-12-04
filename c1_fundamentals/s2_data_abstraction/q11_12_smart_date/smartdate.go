package q11_12_smart_date

import (
	"errors"
	"math"
)

type Date struct {
	Year, Day uint
	Month     Month
}

var ErrInvalidDate = errors.New("date is invalid")

func New(year, month, day uint) (Date, error) {
	if !validMonth(month) || !validDayOfMonth(year, Month(month), day) {
		return Date{}, ErrInvalidDate
	}

	return Date{Year: year, Month: Month(month), Day: day}, nil
}

// DayOfWeek uses the formula described at https://cs.uwaterloo.ca/~alopez-o/math-faq/node73.html to
// return the day of the week the date corresponds to.
func (d Date) DayOfWeek() Day {
	// Formula requires March to be the first month of the year.
	year := float64(d.Year)
	month := float64(d.Month) - 2
	if month < 1 {
		month = 12 - math.Abs(month)
		year--
	}
	// Sunday = 0, ..., Saturday = 6
	day := math.Mod(float64(d.Day-1)+math.Floor(2.6*month-0.2)-2*20+year+year/4+20/4, 7)
	return Day(day + 1)
}

func validMonth(month uint) bool {
	if month < 1 || month > 12 {
		return false
	}

	return true
}

func validDayOfMonth(year uint, month Month, day uint) bool {
	if day < 1 {
		return false
	}

	if day > month.Days(year) {
		return false
	}

	return true
}

type Month uint

const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var _daysInMonths = [14]uint{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31, 29}

func (m Month) Days(year uint) uint {
	if m == February && isLeap(year) {
		return _daysInMonths[13]
	}

	return _daysInMonths[m]
}

func isLeap(year uint) bool {
	if year%100 == 0 && year%400 != 0 {
		return false
	}

	return year%4 == 0
}

type Day uint

var _days = [8]string{"", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

const (
	Sunday Day = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d Day) String() string {
	return _days[d]
}
