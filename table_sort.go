// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tstable

// Import Go standard packages and tserr
import (
	"sort" // sort

	"github.com/thorsphere/tserr" // tserr
)

// byRow implements sort.Interface fo Table based on
// the selected row, which is given by the row index in field key. Initially,
// the first row is selected. The selected row can be changed with function SortBy.
type byRow Table

// Len returns the number of rows in Table
func (t byRow) Len() int {
	// Return zero in case rows is nil
	if t.rows == nil {
		return 0
	}
	// Otherwise, return the length of rows
	return len(t.rows)
}

// Swap swaps the rows with indexes i and j.
func (t byRow) Swap(i, j int) {
	// Return without swap if rows is nil
	if t.rows == nil {
		return
	}
	// Return without swap if i or j is higher or equal to the length of rows
	if (i >= len(t.rows)) || (j >= len(t.rows)) {
		return
	}
	// Swap rows with indexes i and j
	t.rows[i], t.rows[j] = t.rows[j], t.rows[i]
}

// Less reports whether the row with index i must
// sort before the row with index j.
func (t byRow) Less(i, j int) bool {
	// Return false in case rows is nil
	if t.rows == nil {
		return false
	}
	// Return false in case i or j is equal or greater than the number of rows
	if (i >= len(t.rows)) || (j >= len(t.rows)) {
		return false
	}
	// Return false in case key is equal or greater than the number of elements in rows with indexes i and j
	if (t.key >= len(t.rows[i])) || (t.key >= len(t.rows[j])) {
		return false
	}
	// Return whether the row with index i must sort before the row with index j
	// based on alphabetical order
	return t.rows[i][t.key] < t.rows[j][t.key]
}

// sort sorts Table t by the selected row, which is given by the row index in field key.
// Initially, the first row is selected. The selected row can be changed with function SortBy.
func (t *Table) sort() error {
	// Return error in case t is nil
	if t == nil {
		return tserr.NilPtr()
	}
	// Sort Table t by the selected row
	sort.Sort(byRow(*t))
	// Return nil
	return nil
}
