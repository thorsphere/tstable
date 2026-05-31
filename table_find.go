// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tstable

// Import package tserr
import (
	"github.com/thorsphere/tserr" // tserr
)

// find returns the index i of the column matching provided header h. It returns -1 and
// an error if the provided header h is not found.
func (t *Table) find(h string) (int, error) {
	// Return -1 and an error if t is nil
	if t == nil {
		return -1, tserr.NilPtr()
	}
	// Return -1 and an error if header is nil or empty
	if len(t.header) == 0 {
		return -1, tserr.Empty("header")
	}
	// Iterate all elements of header
	for i, v := range t.header {
		// Return index i if provided h matches element of header
		if v == h {
			return i, nil
		}
	}
	// Return -1 and an error, if provided h is not found
	return -1, tserr.NotExistent(h)
}
