package tables

import (
	"encoding/binary"
	"errors"
	"os"
)

type NameTable struct {
	Format       uint16       `json:"format"`
	Count        uint16       `json:"count"`
	StringOffset uint16       `json:"string_offset"`
	NameRecords  []NameRecord `json:"name_records"`
	Name         []byte       `json:"name"`
}

type NameRecord struct {
	PlatformId         uint16 `json:"platform_id"`
	PlatformSpecificId uint16 `json:"platform_specific_id"`
	LanguageId         uint16 `json:"language_id"`
	NameId             uint16 `json:"name_id"`
	Length             uint16 `json:"length"`
	Offset             uint16
}

func ReadNameTable(file *os.File, offset uint32) (*NameTable, error) {

	name := &NameTable{}

	_, err := file.Seek(int64(offset), 0)
	if err != nil {
		return nil, errors.New("Cannot seek to name table offset")
	}

	err = binary.Read(file, binary.BigEndian, &name.Format)
	if err != nil {
		return nil, errors.New("Error seeking to name table offset")
	}
	err = binary.Read(file, binary.BigEndian, &name.Count)
	if err != nil {
		return nil, errors.New("Cannot read name table subtable count")
	}
	err = binary.Read(file, binary.BigEndian, &name.StringOffset)
	if err != nil {
		return nil, errors.New("Cannot read name table string offset")
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
		return nil, errors.New("Error seeking into one of the name table records's offset")
	}

	name.Name = make([]byte, stringSize)
	_, err = file.Read(name.Name)
	if err != nil {
		return nil, errors.New("Cannot read one of the name table records")
	}

	return name, nil
}
