// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tstable_test

// Import Go standard library packages as well as tstable, tsfio and tserr
import (
	"fmt"     // fmt
	"testing" // testing

	"github.com/thorsphere/tserr"   // tserr
	"github.com/thorsphere/tsfio"   // tsfio
	"github.com/thorsphere/tstable" // tstable
)

// TestMinTable1 tests the string representation of a table with one column and one row with empty strings as contents. The test fails
// if the retrieved string does not equal to the testdata golden file.
func TestMinTable1(t *testing.T) {
	// Set name of test table
	name := "MinTable1"
	// Retrieve new test table with one column and an empty string in the header
	tbl, e := tstable.New([]string{""})
	// The test fails, if NewTable returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "NewTable", Fn: "table", Err: e}))
	}
	// Add one row with an empty string to the table
	if err := tbl.AddRow([]string{""}); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: err}))
	}
	// Set padding to zero
	if err := tbl.SetPadding(0); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: err}))
	}
	// Evaluate test table
	evalTable(name, tbl, t)
}

// TestMinTable2 tests the string representation of a table with two columns and one row with empty strings as contents. The test fails
// if the retrieved string does not equal to the testdata golden file.
func TestMinTable2(t *testing.T) {
	// Set name of test table
	name := "MinTable2"
	// Retrieve new test table with two columns and empty strings in the header
	tbl, e := tstable.New([]string{"", ""})
	// The test fails, if NewTable returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "NewTable", Fn: "table", Err: e}))
	}
	// Add one row with empty strings to the table
	if err := tbl.AddRow([]string{"", ""}); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: err}))
	}
	// Set padding to zero
	if err := tbl.SetPadding(0); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: err}))
	}
	// Evaluate test table
	evalTable(name, tbl, t)
}

// TestMinTable3 tests the string representation of a table with one columns, only an empty string as header and no rows. The test fails
// if the retrieved string does not equal to the testdata golden file.
func TestMinTable3(t *testing.T) {
	// Set name of test table
	name := "MinTable3"
	// Retrieve new test table with two columns and empty strings in the header
	tbl, e := tstable.New([]string{""})
	// The test fails, if NewTable returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "NewTable", Fn: "table", Err: e}))
	}
	// Set padding to zero
	if err := tbl.SetPadding(0); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: err}))
	}
	// Evaluate test table
	evalTable(name, tbl, t)
}

// TestMinTable4 tests the string representation of a table with two columns, only containing empty strings in the header and no rows. The test fails
// if the retrieved string does not equal to the testdata golden file.
func TestMinTable4(t *testing.T) {
	// Set name of test table
	name := "MinTable4"
	// Retrieve new test table with two columns and empty strings in the header
	tbl, e := tstable.New([]string{"", ""})
	// The test fails, if NewTable returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "NewTable", Fn: "table", Err: e}))
	}
	// Set padding to zero
	if err := tbl.SetPadding(0); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: err}))
	}
	// Evaluate test table
	evalTable(name, tbl, t)
}

// TestSortByErr tests sorting a table by a column which does not exist. The test fails
// if SortBy does not return an error.
func TestSortByErr(t *testing.T) {
	// Retrieve test table
	tbl := testTable(t)
	// Sort by a column which does not exist
	e := tbl.SortBy("Date of Birth")
	// The test fails if SortBy does not return an error
	if e == nil {
		t.Error(tserr.NilFailed("SortBy"))
	}
}

// TestNilHeader tests NewTable to return an error in case the provided header is nil. The test fails
// if NewTable returns a nil error.
func TestNilHeader(t *testing.T) {
	// Table header is nil
	var header []string = nil
	// Retrieve test table
	if _, e := tstable.New(header); e == nil {
		// The test fails if NewTable returns a nil error
		t.Error(tserr.NilFailed("NewTable"))
	}
}

// TestNilRow tests AddRow to return an error in case the provided row is nil. The test fails
// if AddRow returns a nil error.
func TestNilRow(t *testing.T) {
	// Table row is nil
	var row []string = nil
	// Retrieve test table
	tbl := testTable(t)
	// Add row to the table
	if e := tbl.AddRow(row); e == nil {
		// The test fails if AddRow returns a nil error
		t.Error(tserr.NilFailed("AddRow"))
	}
}

// TestEmptyHeader tests NewTable to return an error in case the provided header has zero length. The test fails
// if NewTable returns a nil error.
func TestEmptyHeader(t *testing.T) {
	// Table header with zero length
	var header []string = make([]string, 0)
	// Retrieve test table
	if _, e := tstable.New(header); e == nil {
		// The test fails if NewTable returns a nil error
		t.Error(tserr.NilFailed("NewTable"))
	}
}

// TestEmptyRow tests AddRow to return an error in case the provided row has zero length. The test fails
// if AddRow returns a nil error.
func TestEmptyRow(t *testing.T) {
	// Table row with zero length
	var row []string = make([]string, 0)
	// Retrieve test table
	tbl := testTable(t)
	// Add row to the table
	if e := tbl.AddRow(row); e == nil {
		// The test fails if AddRow returns a nil error
		t.Error(tserr.NilFailed("AddRow"))
	}
}

// TestIncompleteRow tests AddRow to return an error in case the provided row has a different length to the table header.
// The test fails if AddRow returns a nil error.
func TestIncompleteRow(t *testing.T) {
	// Table row with length 2
	row := []string{"Frodo", "Bearer of the One Ring, Sting"}
	// Retrieve test table
	tbl := testTable(t)
	// Add row to the table
	if e := tbl.AddRow(row); e == nil {
		// The test fails if AddRow returns a nil error
		t.Error(tserr.NilFailed("AddRow"))
	}
}

// TestNonPrintableRow tests AddRow to return an error in case the provided row contains non-printable runes.
// The test fails if AddRow returns a nil error.
func TestNonPrintableRow(t *testing.T) {
	// Table row with a non-printable rune
	row := []string{"Frodo", "Bearer of the One Ring", "Sting\n"}
	// Retrieve test table
	tbl := testTable(t)
	// Add row to the table
	if e := tbl.AddRow(row); e == nil {
		// The test fails if AddRow returns a nil error
		t.Error(tserr.NilFailed("AddRow"))
	}
}

// TestTableStringer tests the implementation of String for a Table. The test fails if the string representation
// of the test table does not equal the contents of the test data golden file.
func TestTableStringer(t *testing.T) {
	name := "SimpleGrid"
	// Retrieve test table
	tbl := testTable(t)
	// Retrieve Grid for name
	grid, ok := tstable.AllGrids[name]
	// The test fails if Grid is not found
	if !ok {
		t.Fatal(tserr.NotExistent(name))
	}
	// Set the Grid
	if e := tbl.SetGrid(grid); e != nil {
		// The test fails if SetGrid returns an error
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetGrid", Fn: "table", Err: e}))
	}
	// Sprintln table
	s := fmt.Sprint(tbl)
	// Retrieve the test data golden file contents for the grid
	e := tsfio.EvalGoldenFile(&tsfio.Testcase{Name: name, Data: s})
	// The test fails if the retrieved string representation of the test table does not equal to the contents of the test data golden file
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "EvalGoldenFile", Fn: name, Err: e}))
	}
}

// TestNonPrintableHeader tests NewTable to return an error in case the provided header contains non-printable runes.
// The test fails if NewTable returns a nil error.
func TestNonPrintableHeader(t *testing.T) {
	// Table header with non-printable runes
	header := []string{testStrNp}
	// Retrieve test table
	if _, e := tstable.New(header); e == nil {
		// The test fails if NewTable returns a nil error
		t.Error(tserr.NilFailed("NewTable"))
	}
}

// TestNegativePadding tests SetPadding to return an error in case the provided padding is negative.
// The test fails if SetPadding returns a nil error.
func TestNegativePadding(t *testing.T) {
	// Padding with a negative value
	padding := -1
	// Retrieve test table
	tbl := testTable(t)
	// Set padding
	if e := tbl.SetPadding(padding); e == nil {
		// The test fails if SetPadding returns a nil error
		t.Error(tserr.NilFailed("SetPadding"))
	}
}

// TestNilGrid tests SetGrid to return an error in case the provided Grid is nil.
// The test fails if SetGrid returns a nil error.
func TestNilGrid(t *testing.T) {
	// Grid is nil
	var grid *tstable.Grid = nil
	// Retrieve test table
	tbl := testTable(t)
	// Set grid
	if e := tbl.SetGrid(grid); e == nil {
		// The test fails if SetGrid returns a nil error
		t.Error(tserr.NilFailed("SetGrid"))
	}
}
