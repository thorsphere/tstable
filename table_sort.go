// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tstable

// Import Go standard packages and tserr
import (
	"sort" // sort
	"strings" // strings

	"github.com/thorsphere/tserr" // tserr
)

// sort sorts Table t by the selected row, which is given by the row index in field key.
// Initially, the first row is selected. The selected row can be changed with function SortBy.
func (t *Table) sort() error {
	// Return error in case t is nil
	if t == nil {
		return tserr.NilPtr()
	}
	// Return nil in case rows is nil, because there are no rows to sort
	if t.rows == nil {
		return nil
	}
	// Sort Table t by the selected row using sort.Slice
	sort.Slice(t.rows, func(i, j int) bool {
		// Return false in case i or j is equal or greater than the number of rows
		if (i >= len(t.rows)) || (j >= len(t.rows)) {
			return false
		}
		// Return false in case key is equal or greater than the number of elements in rows with indexes i and j
		if (t.key >= len(t.rows[i])) || (t.key >= len(t.rows[j])) {
			return false
		}
		// Return whether row i must sort before row j alphabetically
		return strings.ToLower(t.rows[i][t.key]) < strings.ToLower(t.rows[j][t.key])
	})
	// Return nil
	return nil
}
