// This package contains the structure to parse the header of the font file.
package header

import (
	"encoding/binary"
	"errors"
	"os"
)

// The representation of the structure of a TrueType (TTF) or OpenType (OTF) font file.
//
// This implementation falls under Apple's TTF specification, and Microsoft's OTF specification.
//
// This low level API allows for direct reading/manipulation of font data.
type FontHeader struct {
	// Represents the structure of the font offset subtable.
	//
	// It serves as the entry point for parsing the font.
	OffsetSubtable *OffsetSubtable

	// Holds a slice of table entries, each one describes the
	// length, offset, validity, and name of the table.
	//
	// The order of tables in this slice matches their appearance
	// in the font file.
	//
	// Number of table directories have to be the same as the NumTables field in the OffsetSubtable
	TableDirectory []*TableEntry
}

// Represents the structure of the font offset subtable.
//
// It serves as the entry point for parsing the font.
type OffsetSubtable struct {
	// Identifies the font file type. Values can be:
	//   0x00010000: TrueType outlines or OpenType with TrueType outlines.
	//   0x4F54544F: OpenType with CFF outlines.
	//   0x74727565: Apple TrueType font.
	//   0x74797031: Apple PostScript Type 1 font.
	Format uint32

	// Number of tables on the font file.
	NumTables uint16

	// (Maximum power of 2 <= numTables)*16.
	SearchRange uint16

	// log2(Maximum power of 2 <= numTables).
	EntrySelector uint16

	// NumTables*16-SearchRange.
	RangeShift uint16
}

// Table entry of the font
type TableEntry struct {
	// 4-byte identifier of the font tag.
	//
	// The table tag represents it's name, like "head" or "glyf".
	Tag [4]byte

	// Checksum of the table.
	Checksum uint32

	// Offset from beginning of the file to the data of the table.
	Offset uint32

	// Length of the table.
	Length uint32
}

// Parses the font file's offset subtable.
func readOffsetSubtable(file *os.File) (*OffsetSubtable, error) {

	header := &OffsetSubtable{}

	err := binary.Read(file, binary.BigEndian, header)
	if err != nil {
		panic(err)
	}
	return header, nil

}

// Parses all the font file's table entries.
func readTableDirectory(file *os.File, numTables uint16) ([]*TableEntry, error) {
	tables := make([]*TableEntry, numTables)
	for i := range tables {
		tables[i] = &TableEntry{}
		err := binary.Read(file, binary.BigEndian, tables[i])
		if err != nil {
			return nil, err
		}
	}
	return tables, nil
}

// Extracts font header structure from the TTF/OTF file
func ParseFontHeader(file *os.File) (*FontHeader, error) {

	fh, err := readOffsetSubtable(file)
	if err != nil {
		err = errors.New("Could not read font header. The font file may be corrupt.")
		return nil, err
	}

	ft, err := readTableDirectory(file, fh.NumTables)
	if err != nil {
		err = errors.New("Could not read font table directory. The font file may be corrupt.")
		return nil, err
	}

	return &FontHeader{
		OffsetSubtable: fh,
		TableDirectory: ft,
	}, nil

}
