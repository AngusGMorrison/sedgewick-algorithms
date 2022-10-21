package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	t, err := readTable(os.Stdin)
	if err != nil {
		log.Fatalf("%v\nusage:\n\tname<string> first_score<int> second_score<int>\n\t[name2<string> first_score2<int> second_score2<int>, ...]", err)
	}

	_, err = t.WriteTo(os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}

// readTable attempts to construct a table by reading from r, returning an error if the data
// present in r is not as expected.
func readTable(r io.Reader) (table, error) {
	scanner := bufio.NewScanner(r)

	var t table
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 3 {
			return nil, errInsufficientFields
		}

		score1, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, atoiError{score: fields[1], cause: err}
		}

		score2, err := strconv.Atoi(fields[2])
		if err != nil {
			return nil, atoiError{score: fields[2], cause: err}
		}

		r := row{
			name:   fields[0],
			score1: score1,
			score2: score2,
		}
		t = append(t, r)
	}

	return t, nil
}

type row struct {
	name           string
	score1, score2 int
}

func (r row) ratio() float64 {
	return float64(r.score1) / float64(r.score2)
}

const tabSize = 4

type table []row

// WriteTo writes the string representation of t to w.
func (t table) WriteTo(w io.Writer) (int64, error) {
	bytes, err := w.Write([]byte(t.String()))
	return int64(bytes), err
}

func (t table) String() string {
	var builder strings.Builder

	// Render heading row.
	nameHeading := "Name"
	builder.WriteString(nameHeading)
	tabsUsedByNameHeading := (len(nameHeading) / tabSize) + 1
	totalTabsNeededForNameColumn := t.tabsNeededForNameColumn()
	needNMoreTabsForNameHeading := minTabs(totalTabsNeededForNameColumn, tabsUsedByNameHeading)
	builder.WriteString(strings.Repeat("\t", needNMoreTabsForNameHeading))
	builder.WriteString("Score 1\tScore 2\tRatio\n")

	// Number of tabs added when name needs 0 more tabs is 0. 0 is a special case.

	// Render rows.
	for _, row := range t {
		builder.WriteString(row.name)
		tabsUsedByName := (len(row.name) / tabSize) + 1
		needNMoreTabsForName := minTabs(totalTabsNeededForNameColumn, tabsUsedByName)
		builder.WriteString(strings.Repeat("\t", needNMoreTabsForName))
		fmt.Fprintf(&builder, "%d\t%d\t%.3f\n", row.score1, row.score2, row.ratio())
	}

	return builder.String()
}

// minTabs minimizes the total tabs required to space table columns. We require a minimum of one tab
// character after each entry, and minTabs handles the special case where a field partially or fully
// occupies the final tab allotted for its column, where simply subtracting the number of tabs we
// want from the number currently occupied would result in one tab too few, but naively adding one
// to the number of tabs we want would result in too many tabs for shorter fields.
func minTabs(want, have int) int {
	return int(math.Max(1, float64(want-have)))
}

// tabesNeededForNameColumn returns the number of tabs occupied by the longest name in the table,
// plus one to leave adequate space between columns.
func (t table) tabsNeededForNameColumn() int {
	var maxNameLen int
	for _, r := range t {
		if len(r.name) > maxNameLen {
			maxNameLen = len(r.name)
		}
	}

	return (maxNameLen / tabSize) + 1
}

// Defining custom errors makes asserting the nature of the errors returned when testing readTable
// much easier.
var errInsufficientFields = errors.New("too few fields to tabulate")

type atoiError struct {
	score string
	cause error
}

func (e atoiError) Error() string {
	return fmt.Sprintf("convert score %s to int: %v", e.score, e.cause)
}

func (e atoiError) Is(target error) bool {
	other, ok := target.(atoiError)
	if !ok || e.score != other.score {
		return false
	}

	return true
}
