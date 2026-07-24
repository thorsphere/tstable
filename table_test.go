// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
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

// TestTablePrint tests the implementation of Print for a Table. The test fails if the string representation
// of the test table does not equal the contents of the test data golden file.
func TestSeparator(t *testing.T) {
	// Set name of test table
	n := "Separator"
	// Set name of grid
	g := "SimpleGrid"
	// Retrieve test table
	tbl := testTableSeparator(t)
	// Retrieve Grid for name
	grid, ok := tstable.AllGrids[g]
	// The test fails if Grid is not found
	if !ok {
		t.Fatal(tserr.NotExistent(g))
	}
	// Set the Grid
	if e := tbl.SetGrid(grid); e != nil {
		// The test fails if SetGrid returns an error
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetGrid", Fn: "table", Err: e}))
	}
	// Sprintln table
	s := fmt.Sprint(tbl)
	e := tsfio.EvalGoldenFile(&tsfio.Testcase{Name: n, Data: s})
	// The test fails if the retrieved string representation of the test table does not equal to the contents of the test data golden file
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "EvalGoldenFile", Fn: n, Err: e}))
	}
	// Try to sort the table though it has a separator, which is not allowed.
	// The test fails if SortBy does not return an error.
	if e := tbl.SortBy(sortby); e == nil {
		t.Error(tserr.NilFailed("SortBy"))
	}
}

// TestSeparatorNil tests AddSeparator to return an error in case the provided table is nil. The test fails
// if AddSeparator returns a nil error.
func TestSeparatorNil(t *testing.T) {
	// Table is nil
	var tbl *tstable.Table = nil
	// Add separator to nil table
	if e := tbl.AddSeparator(); e == nil {
		// The test fails if AddSeparator returns a nil error
		t.Error(tserr.NilFailed("AddSeparator"))
	}
}

// TestSeparatorNoRows tests AddSeparator to return an error in case the provided table has no rows.
// The test fails if AddSeparator returns a nil error.
func TestSeparatorNoRows(t *testing.T) {
	// Create test table with test header
	tbl, e := tstable.New(header)
	// The test fails, if NewTable returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "NewTable", Fn: "table", Err: e}))
	}
	// Add a separator to the table though it has no rows, which is not allowed.
	// The test fails if AddSeparator does not return an error.
	if e := tbl.AddSeparator(); e == nil {
		t.Error(tserr.NilFailed("AddSeparator"))
	}
}

// TestSeparatorSorted tests that adding a separator to a sorted table is not allowed. The test fails
// if AddSeparator does not return an error.
func TestSeparatorSorted(t *testing.T) {
	// Retrieve test table
	tbl := testTable(t)
	// Try to add a separator to the table though it is sorted, which is not allowed.
	// The test fails if AddSeparator does not return an error.
	if e := tbl.AddSeparator(); e == nil {
		t.Error(tserr.NilFailed("AddSeparator"))
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

// TestAddRowNil tests AddRow to return an error in case the provided table is nil. The test fails
// if AddRow returns a nil error.
func TestAddRowNil(t *testing.T) {
	// Table is nil
	var tbl *tstable.Table = nil
	// Add nil row
	if e := tbl.AddRow(gandalf); e == nil {
		// The test fails if AddRow returns a nil error
		t.Error(tserr.NilFailed("AddRow"))
	}
}

// TestAddRowWidth tests AddRow to return an error in case the provided row has a different number of
// elements than the table header. The test fails if AddRow returns a nil error.
func TestAddRowWidth(t *testing.T) {
	// Retrieve test table
	tbl := testTable(t)
	// Add row to the table that has a different number of elements than the table header
	if e := tbl.AddRow(gandalf[:len(gandalf)-1]); e == nil {
		// The test fails if AddRow returns a nil error
		t.Error(tserr.NilFailed("AddRow"))
	}
}

// TestStringNil tests String to return an error in case the provided table is nil. The test fails
// if String returns a nil error.
func TestStringNil(t *testing.T) {
	// Table is nil
	var tbl *tstable.Table = nil
	// String nil table
	if e := fmt.Sprint(tbl); e != tserr.NilPtr().Error() {
		// The test fails if String returns a nil error
		t.Error(tserr.NilFailed("String"))
	}
}

// TestPrintNil tests Print to return an error in case the provided table is nil. The test fails
// if Print returns a nil error.
func TestPrintNil(t *testing.T) {
	// Table is nil
	var tbl *tstable.Table = nil
	// Print nil table
	if _, e := tbl.Print(); e == nil {
		// The test fails if Print returns a nil error
		t.Error(tserr.NilFailed("Print"))
	}
}

// TestPrintEmptyHeader tests Print to return an error in case the provided table has an empty header. The test fails
// if Print returns a nil error.
func TestPrintEmptyHeader(t *testing.T) {
	// Retrieve test table
	tbl := &tstable.Table{}
	// Print empty header table
	if _, e := tbl.Print(); e == nil {
		// The test fails if Print returns a nil error
		t.Error(tserr.NilFailed("Print"))
	}
}

// TestSortByNil tests SortBy to return an error in case the provided table is nil. The test fails
// if SortBy returns a nil error.
func TestSortByNil(t *testing.T) {
	// Table is nil
	var tbl *tstable.Table = nil
	// Sort by nil column
	if e := tbl.SortBy(""); e == nil {
		// The test fails if SortBy returns a nil error
		t.Error(tserr.NilFailed("SortBy"))
	}
}

// TestSortByEmptyHeader tests SortBy to return an error in case the provided column header is empty. The test fails
// if SortBy returns a nil error.
func TestSortByEmptyHeader(t *testing.T) {
	// Retrieve test table
	var tbl *tstable.Table = new(tstable.Table)
	// Sort by empty column
	if e := tbl.SortBy(""); e == nil {
		// The test fails if SortBy returns a nil error
		t.Error(tserr.NilFailed("SortBy"))
	}
}

// TestSetMultilineNil tests SetMultiline to return an error in case the provided table is nil. The test fails
// if SetMultiline returns a nil error.
func TestSetMultilineNil(t *testing.T) {
	// Table is nil
	var tbl *tstable.Table = nil
	// Set multiline nil table
	if e := tbl.SetMultiline(""); e == nil {
		// The test fails if SetMultiline returns a nil error
		t.Error(tserr.NilFailed("SetMultiline"))
	}
}

// TestSetMultilineEmptyHeader tests SetMultiline to return an error in case the provided column header is empty.
// The test fails if SetMultiline returns a nil error.
func TestSetMultilineEmptyHeader(t *testing.T) {
	// Retrieve test table
	var tbl *tstable.Table = new(tstable.Table)
	// Set multiline empty header table
	if e := tbl.SetMultiline(""); e == nil {
		// The test fails if SetMultiline returns a nil error
		t.Error(tserr.NilFailed("SetMultiline"))
	}
}

// TestSetPaddingNil tests SetPadding to return an error in case the provided table is nil. The test fails
// if SetPadding returns a nil error.
func TestSetPaddingNil(t *testing.T) {
	// Table is nil
	var tbl *tstable.Table = nil
	// Set padding nil table
	if e := tbl.SetPadding(1); e == nil {
		// The test fails if SetPadding returns a nil error
		t.Error(tserr.NilFailed("SetPadding"))
	}
}

// TestSetGridNil tests SetGrid to return an error in case the provided table is nil. The test fails
// if SetGrid returns a nil error.
func TestSetGridNil(t *testing.T) {
	// Table is nil
	var tbl *tstable.Table = nil
	// Set grid nil table
	if e := tbl.SetGrid(&tstable.DoubleBorderGrid); e == nil {
		// The test fails if SetGrid returns a nil error
		t.Error(tserr.NilFailed("SetGrid"))
	}
}

// TestSetMultilineNil tests SetMultiline to return an error in case the provided table is nil. The test fails
// if SetMultiline returns a nil error.
func TestSetMultilineWidthNil(t *testing.T) {
	// Table is nil
	var tbl *tstable.Table = nil
	// Set multiline width nil table
	if e := tbl.SetMultilineWidth(15); e == nil {
		// The test fails if SetMultilineWidth returns a nil error
		t.Error(tserr.NilFailed("SetMultilineWidth"))
	}
}

// TestSetMultilineWidthNegative tests SetMultilineWidth to return an error in case the provided width is negative.
// The test fails if SetMultilineWidth returns a nil error.
func TestSetMultilineWidthNegative(t *testing.T) {
	// Retrieve test table
	tbl := testTable(t)
	// Set multiline width to a negative value
	if e := tbl.SetMultilineWidth(-1); e == nil {
		// The test fails if SetMultilineWidth returns a nil error
		t.Error(tserr.NilFailed("SetMultilineWidth"))
	}
}

// TestSetMultilineWidth tests SetMultilineWidth to return nil to indicate success. The test fails
// if SetMultilineWidth returns an error.
func TestSetMultilineWidth(t *testing.T) {
	// Retrieve test table
	tbl := testTable(t)
	// Set multiline width
	if e := tbl.SetMultilineWidth(15); e != nil {
		// The test fails if SetMultilineWidth returns a nil error
		t.Error(tserr.NilExpected("SetMultilineWidth"))
	}
}

// TestSetMultilineWrongColumn tests SetMultiline to return an error in case the provided column header is wrong.
// The test fails if SetMultiline returns a nil error.
func TestSetMultilineWrongColumn(t *testing.T) {
	// Retrieve test table
	tbl := testTable(t)
	// Set multiline to a column which does not exist
	if e := tbl.SetMultiline("Wrong"); e == nil {
		// The test fails if SetMultiline returns a nil error
		t.Error(tserr.NilFailed("SetMultiline"))
	}
}

// TestSetMultiline tests SetMultiline to return nil to indicate success. The test fails
// if SetMultiline returns an error.
func TestSetMultiline(t *testing.T) {
	// Retrieve test table
	tbl := testTable(t)
	// Set multiline to a column which does not exist
	if e := tbl.SetMultiline("Title"); e != nil {
		// The test fails if SetMultiline returns a nil error
		t.Error(tserr.NilExpected("SetMultiline"))
	}
	evalTable("SimpleGridMultiline", tbl, t)
}

func TestSetMultilineLong(t *testing.T) {
	// Retrieve test table
	tbl := testMultilineTable(t)
	// Set multiline to a column which does not exist
	if e := tbl.SetMultiline("Quote"); e != nil {
		// The test fails if SetMultiline returns a nil error
		t.Error(tserr.NilExpected("SetMultiline"))
	}
	evalTable("SimpleGridMultilineLong", tbl, t)
}
