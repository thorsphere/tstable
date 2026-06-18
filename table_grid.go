// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tstable

// Import Go standard library package strings as well as tsfio and tserr
import (
	"strings" // strings

	"github.com/thorsphere/tserr" // tserr
	"github.com/thorsphere/tsfio" // tsfio
)

// spaces returns the spaces as string for padding. It returns an empty string and an error, if any.
func (t *Table) spaces() (string, error) {
	// Return an empty string and an error if t is nil
	if t == nil {
		return "", tserr.NilPtr()
	}
	// Return an empty string and an error if padding is negative
	if t.padding < 0 {
		return "", tserr.Higher(&tserr.HigherArgs{Var: "padding", Actual: int64(t.padding), LowerBound: 0})
	}
	// Return spaces for padding
	return strings.Repeat(" ", t.padding), nil
}

// hline returns a horizontal grid line (on top) of row with index r for table t as a string. It returns
// an empty string and an error, if any.
func (t *Table) writeHLine(b *strings.Builder, r int) error {
	// Return an error if t is nil
	if t == nil {
		return tserr.NilPtr()
	}
	// Return an error if b is nil
	if b == nil {
		return tserr.NilPtr()
	}
	// Return an error if width is nil
	if t.width == nil {
		return tserr.NilPtr()
	}
	// Return an error if rows is nil
	if t.rows == nil {
		return tserr.NilPtr()
	}
	// Add initial padding to the horizontal line
	spaces, e := t.spaces()
	// Return an error if spaces fails
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "spaces", Fn: "table", Err: e})
	}
	// Retrieve the horizontal line rune
	h, e := t.h_rune(r)
	// Return an error if h_rune fails
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "h_rune", Fn: "table", Err: e})
	}
	// Add initial padding to the horizontal line
	b.WriteString(spaces)
	// Iterate width column by column
	for i, w := range t.width {
		// Return an error if width is negative
		if w < 0 {
			return tserr.Higher(&tserr.HigherArgs{Var: "width", Actual: int64(w), LowerBound: 0})
		}
		// Retrieve the horizontal vertical rune
		hv, e := t.hv_rune(r, i)
		// Return an error if hv_rune fails
		if e != nil {
			return tserr.Op(&tserr.OpArgs{Op: "hv_rune", Fn: "table", Err: e})
		}
		// Set width of column i to multiline width if multiline
		if t.multiline[i] {
			// Set multiline width
			w = t.mlwidth
		}
		// Add a horizontal line runes for each column of the table
		b.WriteString(hv)
		b.WriteString(strings.Repeat(h, t.padding+t.padding+w))
	}
	// Retrieve the horizontal vertical rune at the end of the horizontal line
	hv, e := t.hv_rune(r, len(t.header))
	// Return an error if hv_rune fails
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "hv_rune", Fn: "table", Err: e})
	}
	// End rune of the horizontal line
	b.WriteString(hv)
	b.WriteString("\n")
	// Return nil to indicate success
	return nil
}

// h_rune returns the horizontal grid line rune for row r. The returned rune depends on whether
// it is the top line with r = 0, the bottom line with r = rmax or a line in between. It returns an
// empty string and an error if r is negative or r is higher than rmax.
func (t *Table) h_rune(r int) (string, error) {
	// Return an empty string and an error if t is nil
	if t == nil {
		return "", tserr.NilPtr()
	}
	// Return an empty string and an error if r is negative
	if r < 0 {
		return "", tserr.Higher(&tserr.HigherArgs{Var: "row index", Actual: int64(r), LowerBound: 0})
	}
	// Return an empty string and an error if rows is nil
	if t.rows == nil {
		return "", tserr.NilPtr()
	}
	// Maximum number of horizontal grid lines
	rmax := len(t.rows) + 1
	// Return an empty string and an error if r is higher tham rmax
	if r > rmax {
		return "", tserr.Lower(&tserr.LowerArgs{Var: "row index", Actual: int64(r), HigherBound: int64(rmax)})
	}
	// Return the horizontal border grid line rune for the top line or the bottom line
	if (r == 0) || (r == rmax) {
		return tsfio.RuneToPrintable(t.grid.Hb), nil
	}
	// Return the horizontal grid line rune
	return tsfio.RuneToPrintable(t.grid.Hi), nil
}

// hv_rune returns the horizontal vertical grid line for row r and column c. It returns an
// an empty string and an error if r or c are negative as well as if r or c are higher than rmax or cmax.
func (t *Table) hv_rune(r, c int) (string, error) {
	// Return an empty string and an error if t is nil
	if t == nil {
		return "", tserr.NilPtr()
	}
	// Return an empty string and an error if r is negative
	if r < 0 {
		return "", tserr.Higher(&tserr.HigherArgs{Var: "row index", Actual: int64(r), LowerBound: 0})
	}
	// Return an empty string and an error if c is negative
	if c < 0 {
		return "", tserr.Higher(&tserr.HigherArgs{Var: "column index", Actual: int64(c), LowerBound: 0})
	}
	// Return an empty string and an error if rows is nil
	if t.rows == nil {
		return "", tserr.NilPtr()
	}
	// Return an empty string and an error if header is nil
	if t.header == nil {
		return "", tserr.NilPtr()
	}
	// Maximum number of horizontal grid lines
	rmax := len(t.rows) + 1
	// Maximum number of vertical grid lines
	cmax := len(t.header)
	// Return an empty string and an error if r is higher than rmax
	if r > rmax {
		return "", tserr.Lower(&tserr.LowerArgs{Var: "row index", Actual: int64(r), HigherBound: int64(rmax)})
	}
	// Return an empty string and an error if c is higher than cmax
	if c > cmax {
		return "", tserr.Lower(&tserr.LowerArgs{Var: "column index", Actual: int64(c), HigherBound: int64(cmax)})
	}
	// First horizontal grid line
	if r == 0 {
		// First column
		if c == 0 {
			return tsfio.RuneToPrintable(t.grid.Hvtl), nil
		}
		// Last column
		if c == cmax {
			return tsfio.RuneToPrintable(t.grid.Hvtr), nil
		}
		// Any other column
		return tsfio.RuneToPrintable(t.grid.Hvt), nil
	}
	// Last horizontal grid line
	if r == rmax {
		// First column
		if c == 0 {
			return tsfio.RuneToPrintable(t.grid.Hvbl), nil
		}
		// Last column
		if c == cmax {
			return tsfio.RuneToPrintable(t.grid.Hvbr), nil
		}
		// Any other column
		return tsfio.RuneToPrintable(t.grid.Hvb), nil
	}
	// First column of any other horizontal line
	if c == 0 {
		return tsfio.RuneToPrintable(t.grid.Hvl), nil
	}
	// Last grid line of any other horizontal line
	if c == cmax {
		return tsfio.RuneToPrintable(t.grid.Hvr), nil
	}
	// Any other horizontal vertical grid line
	return tsfio.RuneToPrintable(t.grid.Hvi), nil
}

// vline returns a vertical grid line for table t as a string. It returns an empty string and an error, if any.
func (t *Table) vline(c int) (string, error) {
	// Return an empty string and an error if t is nil
	if t == nil {
		return "", tserr.NilPtr()
	}
	// Return an empty string and an error if c is negative
	if c < 0 {
		return "", tserr.Higher(&tserr.HigherArgs{Var: "column index", Actual: int64(c), LowerBound: 0})
	}
	// Return an empty string and an error if header is nil
	if t.header == nil {
		return "", tserr.NilPtr()
	}
	spaces, e := t.spaces()
	if e != nil {
		return "", tserr.Op(&tserr.OpArgs{Op: "spaces", Fn: "table", Err: e})
	}
	// Maximum number of vertical grid lines
	cmax := len(t.header)
	// Return an empty string and an error if c is higher than cmax
	if c > cmax {
		return "", tserr.Lower(&tserr.LowerArgs{Var: "column index", Actual: int64(c), HigherBound: int64(cmax)})
	}
	// Set vertical grid line
	v_rune := t.grid.Vi
	// Update vertical grid line for the first or last column
	if (c == 0) || (c == cmax) {
		v_rune = t.grid.Vb
	}
	// Return vertical grid line with padding and nil
	var b strings.Builder
	b.WriteString(spaces)
	b.WriteString(tsfio.RuneToPrintable(v_rune))
	return b.String(), nil
}
