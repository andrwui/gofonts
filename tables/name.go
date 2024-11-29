package tables

import (
	"encoding/binary"
	"os"
)

var (
	CopyrightNoticeID  = 0
	FontFamilyID       = 1
	FontSubfamilyID    = 2
	UniqueSubfamilyID  = 3
	FullNameID         = 4
	VersionNameTableID = 5
	PostScriptNameID   = 6
	TrademarkNoticeID  = 7
)

type NameTable struct {
	Format       uint16
	Count        uint16
	StringOffset uint16
	NameRecords  []NameRecord
	Name         []byte
}

type NameRecord struct {
	PlatformID         uint16
	PlatformSpecificID uint16
	LanguageID         uint16
	NameID             uint16
	Length             uint16
	Offset             uint16
}

func ReadNameTable(file *os.File, offset uint32) (*NameTable, error) {

	name := &NameTable{}

	_, err := file.Seek(int64(offset), 0)
	if err != nil {
		panic(err)
	}

	err = binary.Read(file, binary.BigEndian, &name.Format)
	err = binary.Read(file, binary.BigEndian, &name.Count)
	err = binary.Read(file, binary.BigEndian, &name.StringOffset)
	if err != nil {
		panic(err)
	}

	name.NameRecords = make([]NameRecord, name.Count)
	err = binary.Read(file, binary.BigEndian, &name.NameRecords)
	if err != nil {
		panic(err)
	}

	var stringSize uint16
	for _, record := range name.NameRecords {
		endOffset := record.Offset + record.Length
		if endOffset > stringSize {
			stringSize = endOffset
		}
	}

	_, err = file.Seek(int64(offset)+int64(name.StringOffset), 0)
	if err != nil {
		panic(err)
	}

	name.Name = make([]byte, stringSize)
	_, err = file.Read(name.Name)
	if err != nil {
		panic(err)
	}

	return name, nil

}

// This function returns the given name as string
//
// This is commented because this ain't gotta be here but i want to keep the implementation
func (n NameTable) getName(nameID int) string {

	for _, record := range n.NameRecords {
		if int64(record.NameID) == int64(nameID) {
			return string(n.Name[record.Offset : record.Offset+record.Length])
		}
	}
	return ""
}
