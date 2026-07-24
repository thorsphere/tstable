// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package tstable_test

// Import Go standard library package testing as well as tstable, tserr and tsfio
import (
	"testing" // testing

	"github.com/thorsphere/tserr"   // tserr
	"github.com/thorsphere/tsfio"   // tsfio
	"github.com/thorsphere/tstable" // tstable
)

// Test table definition
var (
	header  = []string{"Fellowship member", "Title", "Weapon"}           // Test table header
	gandalf = []string{"Gandalf", "The Grey", "Wizard staff"}            // Test table row
	aragorn = []string{"Aragorn", "King of Gondor", "Sword"}             // Test table row
	legolas = []string{"Legolas", "Prince of the Woodland Realm", "Bow"} // Test table row
	gimli   = []string{"Gimli", "Lord of the Glittering Caves", "Axe"}   // Test table row
	boromir = []string{"Boromir", "Captain of the White Tower", "Sword"} // Test table row
)

// Test constants
const (
	sortby    = "Weapon"        // Test table sort by column
	padding   = 1               // Test table padding
	testStrP  = "Andúril"       // Test string with only printable runes
	testStrNp = testStrP + "\n" // Test string with a non-printable rune
)

// Test multiline tables
var (
	header_long  = append(header, "Quote")
	gandalf_long = append(gandalf, "You cannot pass, I am a servant of the Secret Fire, wielder of the flame of Anor. You cannot pass. The dark fire will not avail you, flame of Udûn. Go back to the Shadow! You cannot pass.")
	aragorn_long = append(aragorn, "But I am the real Strider, fortunately. I am Aragorn son of Arathorn; and if by life or death I can save you, I will.")
	legolas_long = append(legolas, "Yes, they are elves, and they say that you breathe so loud they could shoot you in the dark.")
	gimli_long   = append(gimli, "Far, far below the deepest delvings of the dwarves, the world is gnawed by nameless things.")
	boromir_long = append(boromir, "I_tried_to_take_the_ring_from_Frodo,_I_am_sorry._I_have_paid.")
)

// testTable creates the test table and returns a pointer to the test table.
func testTable(t *testing.T) *tstable.Table {
	// Panic if t is nil
	if t == nil {
		panic(tserr.NilPtr())
	}
	// Create test table with test header
	tbl, e := tstable.New(header)
	// The test fails, if NewTable returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "NewTable", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(gandalf); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(aragorn); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(legolas); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(gimli); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Set padding of the test table
	if e := tbl.SetPadding(padding); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: e}))
	}
	// Set sort by column
	if e := tbl.SortBy(sortby); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SortBy", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(boromir); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Return the test table
	return tbl
}

func testTableSeparator(t *testing.T) *tstable.Table {
	// Panic if t is nil
	if t == nil {
		panic(tserr.NilPtr())
	}
	// Create test table with test header
	tbl, e := tstable.New(header)
	// The test fails, if NewTable returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "NewTable", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(gandalf); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(aragorn); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Add a separator to the test table
	if e := tbl.AddSeparator(); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddSeparator", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(legolas); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(gimli); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Set padding of the test table
	if e := tbl.SetPadding(padding); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(boromir); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Return the test table
	return tbl
}

// testMultilineTable creates the test table with multiline columns and returns a pointer to the test table.
func testMultilineTable(t *testing.T) *tstable.Table {
	// Panic if t is nil
	if t == nil {
		panic(tserr.NilPtr())
	}
	// Create test table with test header
	tbl, e := tstable.New(header_long)
	// The test fails, if NewTable returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "NewTable", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(gandalf_long); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(aragorn_long); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(legolas_long); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(gimli_long); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Set padding of the test table
	if e := tbl.SetPadding(padding); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: e}))
	}
	// Set sort by column
	if e := tbl.SortBy(sortby); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SortBy", Fn: "table", Err: e}))
	}
	// Add a row to the test table
	if e := tbl.AddRow(boromir_long); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	// Set multiline columns
	if e := tbl.SetMultiline("Quote"); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetMultiline", Fn: "table", Err: e}))
	}
	// Set multiline width of the test table
	if e := tbl.SetMultilineWidth(50); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetMultilineWidth", Fn: "table", Err: e}))
	}
	// Return the test table
	return tbl
}

// evalTable evaluates a table if it equals the test data from the golden file provided by the table name. The test fails
// if the string representation of the table does not equal to the contents of the golden file.
func evalTable(name string, tbl *tstable.Table, t *testing.T) {
	// Panic if t is nil
	if t == nil {
		panic(tserr.NilPtr())
	}
	// Retrieve the string representation of the test table
	s, e := tbl.Print()
	// The test fails if Print fails.
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Print", Fn: name, Err: e}))
	}
	/*
		if e = tsfio.CreateGoldenFile(&tsfio.Testcase{Name: name, Data: s}); e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{Op: "CreateGoldenFile", Fn: name, Err: e}))
		}
	*/
	// Retrieve the test data golden file contents for the grid
	e = tsfio.EvalGoldenFile(&tsfio.Testcase{Name: name, Data: s})
	// The test fails if the retrieved string representation of the test table does not equal to the contents of the test data golden file
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "EvalGoldenFile", Fn: name, Err: e}))
	}
}
