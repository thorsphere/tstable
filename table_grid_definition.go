// Copyright (c) 2023-2026 thorsphere.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tstable

// A Grid contains the runes to define the grid format of a table. A table grid is defined by thirteen runes.
// A rune is allowed to be empty.
//
//	Hi:   	horizontal inside, separation between header and the rest of the table rows
//	Hb:		horizontal border, at the top and bottom of the table
//	Vi:		vertical inside, separation between table columns
//	Vb:		vertical border, at the left and right side of the table
//	Hvi:	horizontal vertical inside
//	Hvl:	horizontal vertical left
//	Hvr:	horizontal vertical right
//	Hvt:	horizontal vertical top
//	Hvb:	horizontal vertical bottom
//	Hvtl:	horizontal vertical top left
//	Hvbl:	horizontal vertical bottom left
//	Hvtr:	horizontal vertcial top right
//	Hvbr:	horizontal vertcial bottom right
//
// A table grid has an outside border. The header row is separated from the table rows by a horizontal grid line.
// Table rows do not have a grid line between the rows. Columns are divided by an inside grid line.
type Grid struct {
	Hi, Hb, Vi, Vb, Hvi, Hvl, Hvr, Hvt, Hvb, Hvtl, Hvbl, Hvtr, Hvbr rune
}

var (
	// The EmptyGrid defines an empty table grid. A table with the EmptyGrid does not have a grid.
	EmptyGrid = Grid{}

	// The DoubleBorderGrid has a double-lined border.
	DoubleBorderGrid = Grid{
		Hi:   '\u2500', // horizontal inside				─
		Hb:   '\u2550', // horizontal border				═
		Vi:   '\u2502', // vertical inside					│
		Vb:   '\u2551', // vertical border					║
		Hvi:  '\u253C', // horizontal vertical inside		┼
		Hvl:  '\u255F', // horizontal vertical left			╟
		Hvr:  '\u2562', // horizontal vertical right		╢
		Hvt:  '\u2564', // horizontal vertical top			╤
		Hvb:  '\u2567', // horizontal vertical bottom		╧
		Hvtl: '\u2554', // horizontal vertical top left		╔
		Hvbl: '\u255A', // horizontal vertical bottom left	╚
		Hvtr: '\u2557', // horizontal vertical top right	╗
		Hvbr: '\u255D', // horizontal vertical bottom right	╝
	}

	// The DoubleHorizontalGrid has double-lined horizontal lines.
	DoubleHorizontalGrid = Grid{
		Hi:   '\u2550', // horizontal inside				─
		Hb:   '\u2550', // horizontal border				═
		Vi:   '\u2502', // vertical inside					│
		Vb:   '\u2502', // vertical border					║
		Hvi:  '\u256A', // horizontal vertical inside		┼
		Hvl:  '\u255E', // horizontal vertical left			╟
		Hvr:  '\u2561', // horizontal vertical right		╢
		Hvt:  '\u2564', // horizontal vertical top			╤
		Hvb:  '\u2567', // horizontal vertical bottom		╧
		Hvtl: '\u2552', // horizontal vertical top left		╔
		Hvbl: '\u2558', // horizontal vertical bottom left	╚
		Hvtr: '\u2555', // horizontal vertical top right	╗
		Hvbr: '\u255B', // horizontal vertical bottom right	╝
	}

	// The DoubleVerticalGrid has double-lined vertical lines.
	DoubleVerticalGrid = Grid{
		Hi:   '\u2500', // horizontal inside				─
		Hb:   '\u2500', // horizontal border				─
		Vi:   '\u2551', // vertical inside					║
		Vb:   '\u2551', // vertical border					║
		Hvi:  '\u256B', // horizontal vertical inside		╫
		Hvl:  '\u255F', // horizontal vertical left			╟
		Hvr:  '\u2562', // horizontal vertical right		╢
		Hvt:  '\u2565', // horizontal vertical top			╥
		Hvb:  '\u2568', // horizontal vertical bottom		╨
		Hvtl: '\u2553', // horizontal vertical top left		╓
		Hvbl: '\u2559', // horizontal vertical bottom left	╙
		Hvtr: '\u2556', // horizontal vertical top right	╖
		Hvbr: '\u255C', // horizontal vertical bottom right	╜
	}

	// The DoubleGrid has a double-lined grid.
	DoubleGrid = Grid{
		Hi:   '\u2550', // horizontal inside				═
		Hb:   '\u2550', // horizontal border				═
		Vi:   '\u2551', // vertical inside					║
		Vb:   '\u2551', // vertical border					║
		Hvi:  '\u256C', // horizontal vertical inside		╬
		Hvl:  '\u2560', // horizontal vertical left			╠
		Hvr:  '\u2563', // horizontal vertical right		╣
		Hvt:  '\u2566', // horizontal vertical top			╦
		Hvb:  '\u2569', // horizontal vertical bottom		╩
		Hvtl: '\u2554', // horizontal vertical top left		╔
		Hvbl: '\u255A', // horizontal vertical bottom left	╚
		Hvtr: '\u2557', // horizontal vertical top right	╗
		Hvbr: '\u255D', // horizontal vertical bottom right	╝
	}

	// The RoundGrid has rounded corners.
	RoundGrid = Grid{
		Hi:   '\u2500', // horizontal inside				─
		Hb:   '\u2500', // horizontal border				─
		Vi:   '\u2502', // vertical inside					│
		Vb:   '\u2502', // vertical border					│
		Hvi:  '\u253C', // horizontal vertical inside		┼
		Hvl:  '\u251C', // horizontal vertical left			├
		Hvr:  '\u2524', // horizontal vertical right		┤
		Hvt:  '\u252C', // horizontal vertical top			┬
		Hvb:  '\u2534', // horizontal vertical bottom		┴
		Hvtl: '\u256D', // horizontal vertical top left		╭
		Hvbl: '\u2570', // horizontal vertical bottom left	╰
		Hvtr: '\u256E', // horizontal vertical top right	╮
		Hvbr: '\u256F', // horizontal vertical bottom right	╯
	}

	// The SimpleGrid has single grid lines
	SimpleGrid = Grid{
		Hi:   '\u2500', // horizontal inside				─
		Hb:   '\u2500', // horizontal border				─
		Vi:   '\u2502', // vertical inside					│
		Vb:   '\u2502', // vertical border					│
		Hvi:  '\u253C', // horizontal vertical inside		┼
		Hvl:  '\u251C', // horizontal vertical left			├
		Hvr:  '\u2524', // horizontal vertical right		┤
		Hvt:  '\u252C', // horizontal vertical top			┬
		Hvb:  '\u2534', // horizontal vertical bottom		┴
		Hvtl: '\u250C', // horizontal vertical top left		┌
		Hvbl: '\u2514', // horizontal vertical bottom left	└
		Hvtr: '\u2510', // horizontal vertical top right	┐
		Hvbr: '\u2518', // horizontal vertical bottom right	┘
	}

	// The BoldGrid has bold border lines.
	BoldGrid = Grid{
		Hi:   '\u2500', // horizontal inside				─
		Hb:   '\u2501', // horizontal border				━
		Vi:   '\u2502', // vertical inside					│
		Vb:   '\u2503', // vertical border					┃
		Hvi:  '\u253C', // horizontal vertical inside		┼
		Hvl:  '\u2520', // horizontal vertical left			┠
		Hvr:  '\u2528', // horizontal vertical right		┨
		Hvt:  '\u252F', // horizontal vertical top			┯
		Hvb:  '\u2537', // horizontal vertical bottom		┷
		Hvtl: '\u250F', // horizontal vertical top left		┏
		Hvbl: '\u2517', // horizontal vertical bottom left	┗
		Hvtr: '\u2513', // horizontal vertical top right	┓
		Hvbr: '\u251B', // horizontal vertical bottom right	┛
	}

	// The InterruptedGrid has interrupted grid lines.
	InterruptedGrid = Grid{
		Hi:   '\u254C', // horizontal inside				╌
		Hb:   '\u254D', // horizontal border				╍
		Vi:   '\u254E', // vertical inside					╎
		Vb:   '\u254F', // vertical border					╏
		Hvi:  '\u253C', // horizontal vertical inside		┼
		Hvl:  '\u2520', // horizontal vertical left			┠
		Hvr:  '\u2528', // horizontal vertical right		┨
		Hvt:  '\u252F', // horizontal vertical top			┯
		Hvb:  '\u2537', // horizontal vertical bottom		┷
		Hvtl: '\u250F', // horizontal vertical top left		┏
		Hvbl: '\u2517', // horizontal vertical bottom left	┗
		Hvtr: '\u2513', // horizontal vertical top right	┓
		Hvbr: '\u251B', // horizontal vertical bottom right	┛
	}

	// The DashedGrid has dashed grid lines.
	DashedGrid = Grid{
		Hi:   '\u2504', // horizontal inside				┄
		Hb:   '\u2505', // horizontal border				┅
		Vi:   '\u2506', // vertical inside					┆
		Vb:   '\u2507', // vertical border					┇
		Hvi:  '\u253C', // horizontal vertical inside		┼
		Hvl:  '\u2520', // horizontal vertical left			┠
		Hvr:  '\u2528', // horizontal vertical right		┨
		Hvt:  '\u252F', // horizontal vertical top			┯
		Hvb:  '\u2537', // horizontal vertical bottom		┷
		Hvtl: '\u250F', // horizontal vertical top left		┏
		Hvbl: '\u2517', // horizontal vertical bottom left	┗
		Hvtr: '\u2513', // horizontal vertical top right	┓
		Hvbr: '\u251B', // horizontal vertical bottom right	┛
	}

	// The DottedGrid has dotted grid lines.
	DottedGrid = Grid{
		Hi:   '\u2508', // horizontal inside				┈
		Hb:   '\u2509', // horizontal border				┉
		Vi:   '\u250A', // vertical inside					┊
		Vb:   '\u250B', // vertical border					┋
		Hvi:  '\u253C', // horizontal vertical inside		┼
		Hvl:  '\u2520', // horizontal vertical left			┠
		Hvr:  '\u2528', // horizontal vertical right		┨
		Hvt:  '\u252F', // horizontal vertical top			┯
		Hvb:  '\u2537', // horizontal vertical bottom		┷
		Hvtl: '\u250F', // horizontal vertical top left		┏
		Hvbl: '\u2517', // horizontal vertical bottom left	┗
		Hvtr: '\u2513', // horizontal vertical top right	┓
		Hvbr: '\u251B', // horizontal vertical bottom right	┛
	}

	// AllGrids is a map which contains all Grids of the package. The map returns the *Grid when using the name of the Grid as key.
	AllGrids = map[string]*Grid{
		"EmptyGrid":            &EmptyGrid,
		"DoubleBorderGrid":     &DoubleBorderGrid,
		"DoubleHorizontalGrid": &DoubleHorizontalGrid,
		"DoubleVerticalGrid":   &DoubleVerticalGrid,
		"DoubleGrid":           &DoubleGrid,
		"RoundGrid":            &RoundGrid,
		"SimpleGrid":           &SimpleGrid,
		"BoldGrid":             &BoldGrid,
		"InterruptedGrid":      &InterruptedGrid,
		"DashedGrid":           &DashedGrid,
		"DottedGrid":           &DottedGrid,
	}
)
