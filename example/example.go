// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed by the Functional Source License v1.1
// (FSL-1.1-ALv2) that can be found in the LICENSE file.
package main

import (
	"fmt"

	"github.com/thorsphere/tstable"
)

var (
	header = []string{"Fellowship member", "Title", "Weapon"}
	rows   = [][]string{
		{"Gandalf", "The Grey", "Wizard staff"},
		{"Aragorn", "King of Gondor", "Sword"},
		{"Legolas", "Prince of the Woodland Realm", "Bow"},
		{"Gimli", "Lord of the Glittering Caves", "Axe"},
		{"Boromir", "Captain of the White Tower", "Sword"},
	}
	sortby = "Weapon"
)

func main() {
	example1()
	example2()
}

func example1() {
	tbl, _ := tstable.New(header)
	for _, r := range rows {
		tbl.AddRow(r)
	}
	tbl.SetGrid(&tstable.DoubleBorderGrid)
	fmt.Print(tbl)
}

// Customized grid
func example2() {
	var fancyGrid = tstable.Grid{
		Hi:   '\u2504',
		Hb:   '\u2504',
		Vi:   '\u2506',
		Vb:   '\u2506',
		Hvi:  '\u253C',
		Hvl:  '\u251C',
		Hvr:  '\u2524',
		Hvt:  '\u252C',
		Hvb:  '\u2534',
		Hvtl: '\u256D',
		Hvbl: '\u2570',
		Hvtr: '\u256E',
		Hvbr: '\u256F',
	}

	tbl, _ := tstable.New(header)
	for _, r := range rows {
		tbl.AddRow(r)
	}
	tbl.SetGrid(&fancyGrid)
	tbl.SortBy(sortby)
	fmt.Print(tbl)
}
