# tstable

Go package for tables with a simple API

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorsphere/tstable)](https://pkg.go.dev/mod/github.com/thorsphere/tstable)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorsphere/tstable)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorsphere/tstable)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorsphere/tstable)
[![CodeFactor](https://www.codefactor.io/repository/github/thorsphere/tstable/badge)](https://www.codefactor.io/repository/github/thorsphere/tstable)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorsphere/tstable)

The Go package `tstable` provides a simple interface for generating customizable ASCII tables. Initialize a table using `New` with a slice of header strings, and append data using `AddRow`. The visual output can be configured by modifying padding (`SetPadding`) and borders (`SetGrid`) using either built-in presets or a custom `Grid` configuration. Specific columns can be configured for multi-line text wrapping using `SetMultiline` and the maximum width for wrapped lines can be set with `SetMultilineWidth`. Horizontal separator lines can be inserted between rows using `AddSeparator`. Tables are not sorted by default, which can be overridden via `SortBy`. The final text representation is generated using `Print` or `String()`.

- **Simple**: Without configuration, just function calls
- **Easy to use**: Just define the header of a table and add rows
- **Multiline support**: Configure columns for automatic text wrapping
- **Separator lines**: Insert horizontal lines between groups of rows
- **Tested**: Unit tests with high code coverage.
- **Dependencies**: Only depends on the [Go Standard Library](https://pkg.go.dev/std), [tserr](https://github.com/thorsphere/tserr) and [tsfio](https://github.com/thorsphere/tsfio)

````
┌─────────────────────┬────────────────────────────────┬────────────────┐
│  Fellowship member  │  Title                         │  Weapon        │
├─────────────────────┼────────────────────────────────┼────────────────┤
│  Aragorn            │  King of Gondor                │  Sword         │
│  Boromir            │  Captain of the White Tower    │  Sword         │
│  Gandalf            │  The Grey                      │  Wizard staff  │
│  Gimli              │  Lord of the Glittering Caves  │  Axe           │
│  Legolas            │  Prince of the Woodland Realm  │  Bow           │
└─────────────────────┴────────────────────────────────┴────────────────┘
````

## Usage

The package is installed with 

````go
go get github.com/thorsphere/tstable
````

In the Go app, the package is imported with

````go
import "github.com/thorsphere/tstable"
````

## Table grid

A table grid has an outside border. The header row is separated from the table rows by a horizontal grid line. Table rows do not have a grid line between the rows. Columns are divided by an inside grid line. The package provides a set of grids for table string representation. A grid can be used by providing its reference to SetGrid, for example:

````go
tbl.SetGrid(&tstable.DoubleBorderGrid)
````

<details>
  <summary>See all included grids</summary>
	
  ````
DoubleBorderGrid
  ╔═════════════════════╤════════════════════════════════╤════════════════╗
  ║  Fellowship member  │  Title                         │  Weapon        ║
  ╟─────────────────────┼────────────────────────────────┼────────────────╢
  ║  Aragorn            │  King of Gondor                │  Sword         ║
  ║  Boromir            │  Captain of the White Tower    │  Sword         ║
  ║  Gandalf            │  The Grey                      │  Wizard staff  ║
  ║  Gimli              │  Lord of the Glittering Caves  │  Axe           ║
  ║  Legolas            │  Prince of the Woodland Realm  │  Bow           ║
  ╚═════════════════════╧════════════════════════════════╧════════════════╝
DoubleHorizontalGrid
  ╒═════════════════════╤════════════════════════════════╤════════════════╕
  │  Fellowship member  │  Title                         │  Weapon        │
  ╞═════════════════════╪════════════════════════════════╪════════════════╡
  │  Aragorn            │  King of Gondor                │  Sword         │
  │  Boromir            │  Captain of the White Tower    │  Sword         │
  │  Gandalf            │  The Grey                      │  Wizard staff  │
  │  Gimli              │  Lord of the Glittering Caves  │  Axe           │
  │  Legolas            │  Prince of the Woodland Realm  │  Bow           │
  ╘═════════════════════╧════════════════════════════════╧════════════════╛
DoubleGrid
  ╔═════════════════════╦════════════════════════════════╦════════════════╗
  ║  Fellowship member  ║  Title                         ║  Weapon        ║
  ╠═════════════════════╬════════════════════════════════╬════════════════╣
  ║  Aragorn            ║  King of Gondor                ║  Sword         ║
  ║  Boromir            ║  Captain of the White Tower    ║  Sword         ║
  ║  Gandalf            ║  The Grey                      ║  Wizard staff  ║
  ║  Gimli              ║  Lord of the Glittering Caves  ║  Axe           ║
  ║  Legolas            ║  Prince of the Woodland Realm  ║  Bow           ║
  ╚═════════════════════╩════════════════════════════════╩════════════════╝
RoundGrid
  ╭─────────────────────┬────────────────────────────────┬────────────────╮
  │  Fellowship member  │  Title                         │  Weapon        │
  ├─────────────────────┼────────────────────────────────┼────────────────┤
  │  Aragorn            │  King of Gondor                │  Sword         │
  │  Boromir            │  Captain of the White Tower    │  Sword         │
  │  Gandalf            │  The Grey                      │  Wizard staff  │
  │  Gimli              │  Lord of the Glittering Caves  │  Axe           │
  │  Legolas            │  Prince of the Woodland Realm  │  Bow           │
  ╰─────────────────────┴────────────────────────────────┴────────────────╯
SimpleGrid
  ┌─────────────────────┬────────────────────────────────┬────────────────┐
  │  Fellowship member  │  Title                         │  Weapon        │
  ├─────────────────────┼────────────────────────────────┼────────────────┤
  │  Aragorn            │  King of Gondor                │  Sword         │
  │  Boromir            │  Captain of the White Tower    │  Sword         │
  │  Gandalf            │  The Grey                      │  Wizard staff  │
  │  Gimli              │  Lord of the Glittering Caves  │  Axe           │
  │  Legolas            │  Prince of the Woodland Realm  │  Bow           │
  └─────────────────────┴────────────────────────────────┴────────────────┘
BoldGrid
  ┏━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━┓
  ┃  Fellowship member  │  Title                         │  Weapon        ┃
  ┠─────────────────────┼────────────────────────────────┼────────────────┨
  ┃  Aragorn            │  King of Gondor                │  Sword         ┃
  ┃  Boromir            │  Captain of the White Tower    │  Sword         ┃
  ┃  Gandalf            │  The Grey                      │  Wizard staff  ┃
  ┃  Gimli              │  Lord of the Glittering Caves  │  Axe           ┃
  ┃  Legolas            │  Prince of the Woodland Realm  │  Bow           ┃
  ┗━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━┛
EmptyGrid
  
    Fellowship member    Title                           Weapon        
  
    Aragorn              King of Gondor                  Sword         
    Boromir              Captain of the White Tower      Sword         
    Gandalf              The Grey                        Wizard staff  
    Gimli                Lord of the Glittering Caves    Axe           
    Legolas              Prince of the Woodland Realm    Bow           
  
DoubleVerticalGrid
  ╓─────────────────────╥────────────────────────────────╥────────────────╖
  ║  Fellowship member  ║  Title                         ║  Weapon        ║
  ╟─────────────────────╫────────────────────────────────╫────────────────╢
  ║  Aragorn            ║  King of Gondor                ║  Sword         ║
  ║  Boromir            ║  Captain of the White Tower    ║  Sword         ║
  ║  Gandalf            ║  The Grey                      ║  Wizard staff  ║
  ║  Gimli              ║  Lord of the Glittering Caves  ║  Axe           ║
  ║  Legolas            ║  Prince of the Woodland Realm  ║  Bow           ║
  ╙─────────────────────╨────────────────────────────────╨────────────────╜
InterruptedGrid
  ┏╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┯╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┯╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┓
  ╏  Fellowship member  ╎  Title                         ╎  Weapon        ╏
  ┠╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┨
  ╏  Aragorn            ╎  King of Gondor                ╎  Sword         ╏
  ╏  Boromir            ╎  Captain of the White Tower    ╎  Sword         ╏
  ╏  Gandalf            ╎  The Grey                      ╎  Wizard staff  ╏
  ╏  Gimli              ╎  Lord of the Glittering Caves  ╎  Axe           ╏
  ╏  Legolas            ╎  Prince of the Woodland Realm  ╎  Bow           ╏
  ┗╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┷╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┷╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┛
DashedGrid
  ┏┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┯┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┯┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┓
  ┇  Fellowship member  ┆  Title                         ┆  Weapon        ┇
  ┠┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┼┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┼┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┨
  ┇  Aragorn            ┆  King of Gondor                ┆  Sword         ┇
  ┇  Boromir            ┆  Captain of the White Tower    ┆  Sword         ┇
  ┇  Gandalf            ┆  The Grey                      ┆  Wizard staff  ┇
  ┇  Gimli              ┆  Lord of the Glittering Caves  ┆  Axe           ┇
  ┇  Legolas            ┆  Prince of the Woodland Realm  ┆  Bow           ┇
  ┗┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┷┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┷┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┛
DottedGrid
  ┏┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┯┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┯┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┓
  ┋  Fellowship member  ┊  Title                         ┊  Weapon        ┋
  ┠┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┼┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┼┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┨
  ┋  Aragorn            ┊  King of Gondor                ┊  Sword         ┋
  ┋  Boromir            ┊  Captain of the White Tower    ┊  Sword         ┋
  ┋  Gandalf            ┊  The Grey                      ┊  Wizard staff  ┋
  ┋  Gimli              ┊  Lord of the Glittering Caves  ┊  Axe           ┋
  ┋  Legolas            ┊  Prince of the Woodland Realm  ┊  Bow           ┋
  ┗┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┷┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┷┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┛
  ````
</details>

A custom grid can also be provided to SetGrid. A custom grid is defined with the Grid struct type. The Grid struct type contains the runes to define the grid format of a table. A table grid is defined by thirteen runes. A rune is allowed to be empty.

````go
type Grid struct {
	Hi, Hb, Vi, Vb, Hvi, Hvl, Hvr, Hvt, Hvb, Hvtl, Hvbl, Hvtr, Hvbr rune
}
//	Hi:   	horizontal inside, separation between header and the rest of the table rows
//	Hb:	horizontal border, at the top and bottom of the table
//	Vi:	vertical inside, separation between table columns
//	Vb:	vertical border, at the left and right side of the table
//	Hvi:	horizontal vertical inside
//	Hvl:	horizontal vertical left
//	Hvr:	horizontal vertical right
//	Hvt:	horizontal vertical top
//	Hvb:	horizontal vertical bottom
//	Hvtl:	horizontal vertical top left
//	Hvbl:	horizontal vertical bottom left
//	Hvtr:	horizontal vertcial top right
//	Hvbr:	horizontal vertcial bottom right
````

| `Hvtl` | `Hb`       | `Hvt` | `Hb`       | `Hvtr` |
|------|----------|-----|----------|------|
| `Vb`   | header_1 | `Vi`  | header_2 | `Vb`   |
| `Hvl`  | `Hi`       | `Hvi` | `Hi`       | `Hvr`  |
| `Vb`   | cell_11  | `Vi`  | cell_12  | `Vb`   |
| `Vb`   | cell_21  | `Vi` | cell_22  | `Vb` |
| `Hvbl` | `Hb`       | `Hvb` | `Hb`       | `Hvbr` |

An example with a custom table grid is included in [example/example.go](https://github.com/thorsphere/tstable/blob/main/example/example.go)

## Multiline Columns

By default, all columns display content in a single line. However, you can mark specific columns for multi-line text wrapping. This is useful for columns that contain long descriptions or text that should wrap to multiple lines.

Use `SetMultiline` to mark a column for multi-line wrapping by providing its header string:

```go
tbl.SetMultiline("Description")
```

The maximum width for wrapped lines can be set with `SetMultilineWidth`:

```go
tbl.SetMultilineWidth(30) // Default is 20 characters
```

When a column is marked for multiline wrapping:
- Long text is split into multiple lines at word boundaries when possible
- If a single word exceeds the maximum width, it's forcibly broken
- The table automatically adjusts row heights to accommodate wrapped content
- All other columns maintain their single-line format

## Separator Lines

You can add horizontal separator lines between rows to visually group related data. Call `AddSeparator` immediately after the row that should be followed by a separator line:

```go
tbl.AddRow([]string{"Gandalf", "The Grey", "Wizard staff"})
tbl.AddRow([]string{"Aragorn", "King of Gondor", "Sword"})
tbl.AddSeparator() // Separator line after Aragorn's row
tbl.AddRow([]string{"Legolas", "Prince of the Woodland Realm", "Bow"})
```

Multiple separator lines can be added. Separator lines are only supported for unsorted tables. Calling AddSeparator on a table that has SortBy enabled, or calling SortBy on a table that already has separators, will return an error.

## Example

````go
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
	tbl, _ := tstable.New(header)
	for _, r := range rows {
		tbl.AddRow(r)
	}
	tbl.SortBy(sortby)
	tbl.SetMultiline("Title")
	tbl.SetMultilineWidth(15)
	for n, g := range tstable.AllGrids {
		tbl.SetGrid(g)
		fmt.Println(n)
		fmt.Print(tbl)
	}
}

````
[Go Playground](https://go.dev/play/p/_9-ZmlEVPDS)

This produces a table where the "Description" column wraps long text into multiple lines while other columns maintain single-line format.

## Documentation & Resources

- [Go Package Documentation](https://pkg.go.dev/github.com/thorsphere/tstable) — Complete API reference
- [Open Source Insights](https://deps.dev/go/github.com%2Fthorsphere%2Ftstable) — Dependency analysis

## ⚖️ License & Commercial Usage

Copyright (c) 2023-2026 thorsphere. All rights reserved.

This project is licensed under the **Functional Source License v1.1 (FSL-1.1-ALv2)**. 

* The use, modification, and redistribution of this Go package is completely free for private, educational, non-commercial, and internal purposes. 
* If you are a company or institution looking to use this package in a commercial product, service, or business environment, you must secure a commercial license.
* Each version of this software automatically converts to the fully open-source Apache License, Version 2.0 on the second anniversary of its release.

For full details, please see the [LICENSE](LICENSE) file.

### 💼 Commercial Licensing & Inquiries

To purchase a commercial license or discuss support options, please reach out directly:

* 📩 **Contact:** business at thorsphere dot com
* 💬 **Response Time:** Usually within a couple of business days.

*Please include your company name and a brief overview of your use case so I can provide the right licensing details.*
