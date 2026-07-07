// Package tstable provides a simple interface for generating customizable ASCII tables.
// Initialize a table using New with a slice of header strings, and append data using AddRow.
// The visual output can be configured by modifying padding (SetPadding) and borders (SetGrid)
// using either built-in presets or a custom Grid configuration.
// Tables are automatically sorted alphabetically by the first column by default, which can
// be overridden via SortBy. Specific columns can be configured for multi-line text wrapping
// using SetMultiline and the maximum width for wrapped lines can be set with SetMultilineWidth.
// The final text representation is generated using Print or String.
//
// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tstable

// Import Go standard packages, lpstats, tsfio and tserr
import (
	"strings"      // strings
	"unicode/utf8" // utf8

	"github.com/thorsphere/tserr" // tserr
	"github.com/thorsphere/tsfio" // tsfio
)

// Table holds the header of the table and all rows of the table. It also contains
// information on the width of each column, the row index for sorting, padding and the table grid.
// Per default, a table has padding 2, a simple grid and is sorted by its first row.
type Table struct {
	header     []string     // Header as a slice of strings
	rows       [][]string   // Rows as a slice of slices of strings
	width      []int        // Width of each row
	multiline  []bool       // Whether a row contains multiple lines
	separators map[int]bool // Map of row indices that should have a separator line after them
	key        int          // Row index for sorting (default first column)
	padding    int          // Padding (default 2)
	mlwidth    int          // Maximum width of multiline columns
	grid       *Grid        // Table grid
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
		padding:    2,                    // default padding
		mlwidth:    20,                   // default maximum width of multiline columns
		grid:       &SimpleGrid,          // with a simple table grid
		header:     h,                    // set header
		rows:       make([][]string, 0),  // allocate and initialize rows
		width:      make([]int, len(h)),  // allocate and initialize width
		multiline:  make([]bool, len(h)), // allocate and initialize multiline
		separators: make(map[int]bool),   // allocate and initialize separators
		key:        -1,                   // set sort key to -1 (no sorting)
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
		return tserr.EqualInt(&tserr.EqualIntArgs{Var: "row", Actual: int64(len(r)), Want: int64(len(t.header))})
	}
	// Return an error if the number of elements in r does not equal to the number of elements of width
	if len(r) != len(t.width) {
		return tserr.EqualInt(&tserr.EqualIntArgs{Var: "row", Actual: int64(len(r)), Want: int64(len(t.width))})
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
	// Return an empty string and an error, if t is nil
	if t == nil {
		return "", tserr.NilPtr()
	}
	// Return an empty string and an error, if header or rows are nil
	if (t.header == nil) || (t.rows == nil) {
		return "", tserr.NilPtr()
	}
	// Return an empty string and an error, if the number of elements in header does not equal the number of elements in width
	if len(t.header) != len(t.width) {
		return "", tserr.EqualInt(&tserr.EqualIntArgs{Var: "table width slice", Actual: int64(len(t.width)), Want: int64(len(t.header))})
	}
	// Sort table by selected row, which is given by the row index in struct field key
	if err := t.sort(); err != nil {
		// Return an empty string and an error, if sorting fails
		return "", tserr.Op(&tserr.OpArgs{Op: "sort", Fn: "table", Err: err})
	}
	// Use strings.Builder to build the return value as a string
	var b strings.Builder
	// Pre-calculate estimated capacity to minimize buffer array allocations
	estRowLen := 1 // Account for newline
	for _, w := range t.width {
		estRowLen += w + (t.padding * 2) + 1
	}
	b.Grow(estRowLen * (len(t.rows) + len(t.separators) + 3)) // Allocating for rows, header, and separating grid lines
	// Add top horizontal grid line to string builder
	e := t.writeHLine(&b, 0)
	// Return an empty string and an error, if hline fails
	if e != nil {
		return "", tserr.Op(&tserr.OpArgs{Op: "writeHLine", Fn: "table", Err: e})
	}
	// Write the header row to the string builder
	e = t.writeMultiRow(&b, t.header)
	// Return an empty string and an error, if printRow fails
	if e != nil {
		return "", tserr.Op(&tserr.OpArgs{Op: "printRow", Fn: "table", Err: e})
	}
	// dd horizontal grid line to return string
	e = t.writeHLine(&b, 1)
	// Return an empty string and an error, if hline fails
	if e != nil {
		return "", tserr.Op(&tserr.OpArgs{Op: "writeHLine", Fn: "table", Err: e})
	}
	// Return string representation if the table does not have rows
	if len(t.rows) == 0 {
		return b.String(), nil
	}
	// Print rows
	for i, r := range t.rows {
		// Write row r to the string builder
		e := t.writeMultiRow(&b, r)
		// Return an empty string and an error, if printRow fails
		if e != nil {
			return "", tserr.Op(&tserr.OpArgs{Op: "printRow", Fn: "table", Err: e})
		}
		// Write separator line after this row if configured
		if t.separators[i] {
			// Add horizontal grid line to return string
			e = t.writeHLine(&b, i+2) // +2 because row 0 of data is at position 2 (after top border and header section)
			// Return an empty string and an error, if hline fails
			if e != nil {
				return "", tserr.Op(&tserr.OpArgs{Op: "writeHLine", Fn: "table", Err: e})
			}
		}
	}
	// Add horizontal grid line to return string
	e = t.writeHLine(&b, len(t.rows)+1)
	// Return an empty string and an error, if hline fails
	if e != nil {
		return "", tserr.Op(&tserr.OpArgs{Op: "writeHLine", Fn: "table", Err: e})
	}
	return b.String(), nil
}

// writeRow writes a row to a string builder
func (t *Table) writeMultiRow(b *strings.Builder, r []string) error {
	// Return an error, if t is nil
	if t == nil {
		return tserr.NilPtr()
	}
	// Return an error, if r is nil
	if r == nil {
		return tserr.NilPtr()
	}
	// Return an error, if b is nil
	if b == nil {
		return tserr.NilPtr()
	}
	// Return an error, if the number of elements in r does not equal the number of elements in the table header
	if len(r) != len(t.header) {
		return tserr.EqualInt(&tserr.EqualIntArgs{Var: "row", Actual: int64(len(r)), Want: int64(len(t.header))})
	}
	// Return an error, if the number of elements in r does not equal the number of elements in width
	if len(r) != len(t.width) {
		return tserr.EqualInt(&tserr.EqualIntArgs{Var: "row", Actual: int64(len(r)), Want: int64(len(t.width))})
	}
	// Define type colLines to hold a slice of strings
	type colLines []string
	// Create a slice of colLines to hold the wrapped lines of each column
	columns := make([]colLines, len(r))
	// Maximum number of lines in any column
	maxLines := 0
	// Iterate over columns
	for i, cell := range r {
		// Wrap cell in columns[i] if multiline
		if t.multiline[i] && t.width[i] > 0 {
			// Wrap cell in columns[i]
			columns[i] = wrapText(cell, t.mlwidth)
		} else { // Otherwise, add cell to columns[i]
			// Add cell to columns[i]
			columns[i] = []string{cell}
		}
		// Update maxLines if column i has more lines than maxLines
		if len(columns[i]) > maxLines {
			// Update maxLines
			maxLines = len(columns[i])
		}
	}
	// Retrieve spaces for padding
	spaces, e := t.spaces()
	// Return an error, if spaces returns an error
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "spaces", Fn: "table", Err: e})
	}
	// Retrieve vertical grid line
	vrline, e := t.vline(len(r))
	// Return an error, if vline fails
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "vline", Fn: "table", Err: e})
	}
	// Iterate all lines of row r
	for l := 0; l < maxLines; l++ {
		// Iterate all elements of row r
		for i := range r {
			// Retrieve top vertical grid line
			vline, e := t.vline(i)
			// Return an error, if vline fails
			if e != nil {
				return tserr.Op(&tserr.OpArgs{Op: "vline", Fn: "table", Err: e})
			}
			// Add vertical grid line to return string and start new line
			b.WriteString(vline)
			// Add padding to return string
			b.WriteString(spaces)
			// Retrieve text of line l
			lineText := ""
			// If line l is within bounds of columns[i]
			if l < len(columns[i]) {
				// Retrieve text of line l
				lineText = columns[i][l]
			}
			// Add line text to return string
			b.WriteString(lineText)
			// Calculate cell width
			cellWidth := 0
			// If column i is multiline
			if t.multiline[i] {
				// Set cell width to multiline width
				cellWidth = t.mlwidth
			} else { // Otherwise, set cell width to width of column i
				// Set cell width to width of column i
				cellWidth = t.width[i]
			}
			// Calculate padding
			padRunes := cellWidth - utf8.RuneCountInString(lineText)
			// Return an error, if padding is negative
			if padRunes < 0 {
				tserr.Higher(&tserr.HigherArgs{Var: "width", Actual: int64(cellWidth), LowerBound: int64(utf8.RuneCountInString(lineText))})
			}
			// Add padding to return string
			b.WriteString(strings.Repeat(" ", padRunes))
		}
		// Add vertical grid line to return string and start new row
		b.WriteString(vrline)
		// Add newline to return string
		b.WriteByte('\n')
	}
	// Return nil to indicate success
	return nil
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
	// Return error in case t has separator(s), because sorting is not allowed for tables with separator(s)
	if len(t.separators) > 0 {
		return tserr.MethodNotAllowed(&tserr.MethodNotAllowedArgs{Method: "SortBy", Resource: "table with separator(s)"})
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

// SetMultiline marks specific columns for multi-line wrapping. Each string in `cols` should match
// a column header; those columns will then wrap long text to the next line when printed.
// Columns not specified are left unchanged.
func (t *Table) SetMultiline(h string) error {
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
	// Set multiline to true for column i
	t.multiline[i] = true
	// Return nil to indicate success
	return nil
}

// wrapText splits s into lines of at most maxRunes runes, breaking at word
// boundaries when possible. If a single word exceeds maxRunes, it is
// forcibly broken.
func wrapText(s string, maxRunes int) []string {
	// Return empty slice if s is empty
	if s == "" {
		return []string{""}
	}
	// Return single line if maxRunes is zero or negative
	if maxRunes <= 0 {
		return []string{s}
	}
	// Split s into lines
	var lines []string
	// Split s into words
	words := strings.Fields(s)
	// Return single line if no words
	if len(words) == 0 {
		return []string{s}
	}
	// Build current line
	var current strings.Builder
	// Rune count of current line (without spaces)
	currentLen := 0
	// Iterate over words
	for i, w := range words {
		// Rune count of word w
		wLen := utf8.RuneCountInString(w)
		if currentLen+1+wLen <= maxRunes {
			// Add space if not first word
			if i != 0 {
				// Add space
				current.WriteByte(' ')
			}
			// Word fits on current line, add it
			current.WriteString(w)
			// Rune count of current line
			currentLen += 1 + wLen
		} else {
			// Word does not fit on current line, start new line
			lines = append(lines, current.String())
			// Reset current line
			current.Reset()
			// If word w is too long, break it
			if wLen+1 > maxRunes {
				// Split w into lines
				for wLen+1 > maxRunes {
					// Append first maxRunes runes of w to lines
					lines = append(lines, w[:maxRunes])
					// Remove first maxRunes runes of w
					w = w[maxRunes:]
					// Rune count of remaining of w
					wLen = utf8.RuneCountInString(w)
				}
			}
			// Append w to current line
			current.WriteString(w)
			// Rune count of current line (without spaces)
			currentLen = wLen
		}
	}
	// Append remaining text to current line
	if current.Len() > 0 {
		lines = append(lines, current.String())
	}
	// Return lines
	return lines
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
		return tserr.Higher(&tserr.HigherArgs{Var: "padding", Actual: int64(p), LowerBound: 0})
	}
	// Set table padding to p
	t.padding = p
	// Return nil
	return nil
}

// SetMultilineWidth sets the multiline width of table t.
// The default multiline width is 20. It returns an error if w is negative.
func (t *Table) SetMultilineWidth(w int) error {
	// Return an error, if t is nil
	if t == nil {
		return tserr.NilPtr()
	}
	// Return an error, if w is negative
	if w <= 0 {
		return tserr.Higher(&tserr.HigherArgs{Var: "width", Actual: int64(w), LowerBound: 0})
	}
	// Set table multiline width to w
	t.mlwidth = w
	// Return nil
	return nil
}

// SetGrid sets the grid for table t when printed. Per default, a new table has a simple grid enabled.
func (t *Table) SetGrid(g *Grid) error {
	// Return an error if t is nil
	if t == nil {
		return tserr.NilPtr()
	}
	// Return an error if g is nil
	if g == nil {
		return tserr.NilPtr()
	}
	// Set table grid to g
	t.grid = g
	// Return nil
	return nil
}

// AddSeparator adds a horizontal separator line after the most recently added row.
// It returns an error if t is nil or if no rows have been added yet.
func (t *Table) AddSeparator() error {
	// Return an error if t is nil
	if t == nil {
		return tserr.NilPtr()
	}
	// Return an error if sorting is enabled, because adding separators is not allowed for sorted tables
	if t.key >= 0 {
		// Return an error if sorting is enabled, because adding separators is not allowed for sorted tables
		return tserr.MethodNotAllowed(&tserr.MethodNotAllowedArgs{Method: "AddSeparator", Resource: "sorted table"})
	}
	// Return an error if no rows have been added yet
	if len(t.rows) == 0 {
		return tserr.Empty("rows")
	}
	// Add a separator after the last row
	t.separators[len(t.rows)-1] = true
	// Return nil to indicate success
	return nil
}
