// Package tstable provides a simple interface for tables. A new instance of a table can be retrieved with
// New and providing a table header. Table rows can be added with AddRow. The table visualization can be
// altered with SetGrid and SetPadding. The package provides a set of grids or a grid can be customized.
// The string representation of a table is retrieved with Print. A table is sorted alphabetically by the
// first column. It can be sorted by other columns with SortBy.
//
// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tstable

// Import Go standard packages, lpstats, tsfio and tserr
import (
	"strings"      // strings
	"unicode/utf8" // utf8

	// lpstats
	"github.com/thorsphere/tserr" // tserr
	"github.com/thorsphere/tsfio" // tsfio
)

// Table holds the header of the table and all rows of the table. It also contains
// information on the width of each column, the row index for sorting, padding and the table grid.
// Per default, a table has padding 2, a simple grid and is sorted by its first row.
type Table struct {
	header  []string   // Header as a slice of strings
	rows    [][]string // Rows as a slice of slices of strings
	width   []int      // Width of each row
	key     int        // Row index for sorting (default first column)
	padding int        // Padding (default 2)
	grid    *Grid      // Table grid
}

// New returns a pointer to a new Table. It expects the header of the table
// h as a slice of strings. It returns nil and an error, if h is nil, has
// zero length or contains non-printable runes. The order of the header is fixed.
func New(h []string) (*Table, error) {
	// Return nil and an error if h is nil or h has zero length
	if len(h) == 0 {
		return nil, tserr.Empty("header")
	}
	// Retrieve whether h contains only printable runes with IsPrintable
	p, e := tsfio.IsPrintable(h)
	// Return nil and an error if IsPrintable fails
	if e != nil {
		return nil, tserr.Op(&tserr.OpArgs{Op: "IsPrintable", Fn: "header", Err: e})
	}
	// Return nil and an error if the header contains non-printable runes
	if !p {
		return nil, tserr.NonPrintable("header")
	}
	// Retrieve a new instance of struct Table
	t := &Table{
		padding: 2,                   // default padding
		grid:    &SimpleGrid,         // with a simple table grid
		header:  h,                   // set header
		rows:    make([][]string, 0), // allocate and initialize rows
		width:   make([]int, len(h)), // allocate and initialize width
		key:     0,                   // set sort key to first column
	}
	// Iterate over elements of h
	for i, c := range h {
		// Set width of column to number of runes in element c of h
		t.width[i] = utf8.RuneCountInString(c)
	}
	// Return pointer to Table
	return t, nil
}

// AddRow appends a row r at the end of the rows of table t. The row r is provided by a slice of strings. Row r must
// contain the same number of elements as the table header. The order of elements must match
// the order of columns defined by the table header. It returns
// an error if t is nil, r is nil or empty or if the number of elements in r does not equal the
// number of elements in the table header or if r contains non-printable runes.
func (t *Table) AddRow(r []string) error {
	// Return an error if t is nil
	if t == nil {
		return tserr.NilPtr()
	}
	// Return an error if r is nil or r is empty
	if len(r) == 0 {
		return tserr.Empty("row")
	}
	// Return an error if the number of elements in r does not equal to the number of elements of the table header
	if len(r) != len(t.header) {
		return tserr.Equal(&tserr.EqualArgs{Var: "row", Actual: int64(len(r)), Want: int64(len(t.header))})
	}
	// Return an error if the number of elements in r does not equal to the number of elements of width
	if len(r) != len(t.width) {
		return tserr.Equal(&tserr.EqualArgs{Var: "row", Actual: int64(len(r)), Want: int64(len(t.width))})
	}
	// Retrieve in p whether r only contains printable runes
	p, e := tsfio.IsPrintable(r)
	// If IsPrintable returns an error, return that error
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "IsPrintable", Fn: "row", Err: e})
	}
	// Return an error if row r contains non-printable runes
	if !p {
		return tserr.NonPrintable("row")
	}
	// Append row r at the end of the rows of the table
	t.rows = append(t.rows, r)
	// Iterate all elements of row r
	for i, c := range r {
		// Set the width of column i to the maximum of current width of column i and number or runes in element i of row r
		// This adjusts the width of column i, if c contains more runes than previous elements of column i
		t.width[i] = max(t.width[i], utf8.RuneCountInString(c))
	}
	// Return nil
	return nil
}

// String implements the Stringer interface. It returns the string representation of
// table t. It returns an error text in case of an error.
func (t *Table) String() string {
	// Return an error if t is nil
	if t == nil {
		return tserr.NilPtr().Error()
	}
	// Retrieve the string representation of t
	s, e := t.Print()
	// Return the error text, if Print fails.
	if e != nil {
		return e.Error()
	}
	// Return the string representation of t
	return s
}

// Print returns the contents of table t in a string representation. The formatting
// of the table can be altered by changing the padding with SetPadding or setting a different grid with SetGrid.
// The rows are sorted in alphabetical order according to the selected column with
// SortBy. Per default, it is sorted by the first column.
func (t *Table) Print() (string, error) {
	// Initialize return value with an empty string
	text := ""
	// Return an empty string and an error, if t is nil
	if t == nil {
		return text, tserr.NilPtr()
	}
	// Return an empty string and an error, if header or rows are nil
	if (t.header == nil) || (t.rows == nil) {
		return text, tserr.NilPtr()
	}
	// Return an empty string and an error, if the number of elements in header does not equal the number of elements in width
	if len(t.header) != len(t.width) {
		return text, tserr.Equal(&tserr.EqualArgs{Var: "table width slice", Actual: int64(len(t.width)), Want: int64(len(t.header))})
	}
	// Retrieve spaces for padding
	spaces, e := t.spaces()
	// Return an empty string and an error, if spaces returns an error
	if e != nil {
		return text, tserr.Op(&tserr.OpArgs{Op: "spaces", Fn: "table", Err: e})
	}
	// Sort table by selected row, which is given by the row index in struct field key
	if err := t.sort(); err != nil {
		// Return an empty string and an error, if sorting fails
		return text, tserr.Op(&tserr.OpArgs{Op: "sort", Fn: "table", Err: err})
	}
	// Retrieve top horizontal grid line
	hline, e := t.hline(0)
	// Return an empty string and an error, if hline fails
	if e != nil {
		return text, tserr.Op(&tserr.OpArgs{Op: "hline", Fn: "table", Err: e})
	}
	// Add top horizontal grid line to string
	text += hline
	// Print header
	for i, h := range t.header {
		// Return an empty string and an error, if the difference of width of column i and length of h is negative
		if t.width[i]-len(h) < 0 {
			return "", tserr.Higher(&tserr.HigherArgs{Var: "width", Actual: int64(t.width[i]), LowerBound: int64(len(h))})
		}
		// Retrieve top vertical grid line
		vline, e := t.vline(i)
		// Return an empty string and an error, if vline fails
		if e != nil {
			return "", tserr.Op(&tserr.OpArgs{Op: "vline", Fn: "table", Err: e})
		}
		// Add header of column to return string
		text += vline + spaces + h + strings.Repeat(" ", t.width[i]-len(h))
	}
	// Retrieve vertical grid line at the end of the header
	vrline, e := t.vline(len(t.header))
	// Return an empty string and an error, if vline fails
	if e != nil {
		return "", tserr.Op(&tserr.OpArgs{Op: "vline", Fn: "table", Err: e})
	}
	// Add vertical grid line at the end of the header to the return string
	text += vrline + "\n"
	// Retrieve horizontal grid line below header
	hline, e = t.hline(1)
	// Return an empty string and an error, if hline fails
	if e != nil {
		return "", tserr.Op(&tserr.OpArgs{Op: "hline", Fn: "table", Err: e})
	}
	// Add horizontal grid line to return string
	text += hline
	// Return string representation if the table does not have rows
	if len(t.rows) == 0 {
		return text, nil
	}
	// Print rows
	for _, r := range t.rows {
		// Return an empty string and an error, if the size of row r is not equal to the size of width
		if len(t.width) != len(r) {
			return "", tserr.Higher(&tserr.HigherArgs{Var: "sizte of row", Actual: int64(len(r)), LowerBound: int64(len(t.width))})
		}
		// Print row r
		for j, c := range r {
			// Return an empty string and an error, if the difference of width of column j and length of c is negative
			if t.width[j]-len(c) < 0 {
				return "", tserr.Higher(&tserr.HigherArgs{Var: "width", Actual: int64(t.width[j]), LowerBound: int64(len(c))})
			}
			// Retrieve top vertical grid line
			vline, e := t.vline(j)
			// Return an empty string and an error, if vline fails
			if e != nil {
				return text, tserr.Op(&tserr.OpArgs{Op: "vline", Fn: "table", Err: e})
			}
			// Add cell c of row r to return string
			text += vline + spaces + c + strings.Repeat(" ", t.width[j]-len(c))
		}
		// Add vertical grid line to return string and start new row
		text += vrline + "\n"
	}
	// Retrieve bottom horizontal grid line
	hline, e = t.hline(len(t.rows) + 1)
	// Return an empty string and an error, if hline fails
	if e != nil {
		return text, tserr.Op(&tserr.OpArgs{Op: "hline", Fn: "table", Err: e})
	}
	// Add horizontal grid line to return string
	return text + hline, nil
}

// SortBy sets table t to be sorted by column header h. When printing the table, the table will be sorted by column with header h.
// It returns an error if column header h is empty or cannot be found in the table t.
func (t *Table) SortBy(h string) error {
	// Return an error, if t is bil
	if t == nil {
		return tserr.NilPtr()
	}
	if len(t.header) == 0 {
		return tserr.Empty("header")
	}
	// Retrieve index i of column header h
	i, e := t.find(h)
	// Return an error, if find returns an error
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "find", Fn: h, Err: e})
	}
	// Set sort key to index i
	t.key = i
	// Return nil
	return nil
}

// SetPadding sets the table padding to p. The default padding of a new table is 2. Padding p defines the number
// of spaces between the cell grid edges and the cell content. It returns an error if p is negative.
func (t *Table) SetPadding(p int) error {
	// Return an error, if t is nil
	if t == nil {
		return tserr.NilPtr()
	}
	// Return an error, if padding p is negative
	if p < 0 {
		return tserr.Higher(&tserr.HigherArgs{Var: "padding", Actual: int64(t.padding), LowerBound: 0})
	}
	// Set table padding to p
	t.padding = p
	// Return nil
	return nil
}

// SetGrid sets the grid for table t when printed. Per default, a new table has a simple grid enabled.
func (t *Table) SetGrid(g *Grid) error {
	// Return an error if t or g is nil
	if (t == nil) || (g == nil) {
		return tserr.NilPtr()
	}
	// Set table grid to g
	t.grid = g
	// Return nil
	return nil
}
