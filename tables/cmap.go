package tables

import (
	"encoding/binary"
	"errors"
	"os"
)

type CmapTable struct {
	Version   uint16
	NumTables uint16
	Subtables []CmapSubtableHeader
}

type CmapSubtableHeader struct {
	PlatformID uint16
	EncodingID uint16
	Offset     uint32
	Data       []byte
}

func ReadCmapTable(file *os.File, offset uint32) (*CmapTable, error) {
	cmap := &CmapTable{}
	_, err := file.Seek(int64(offset), 0)
	if err != nil {
		return nil, err
	}

	// Read header
	err = binary.Read(file, binary.BigEndian, &cmap.Version)
	if err != nil {
		return nil, errors.New("Cannot read cmap table version")
	}
	err = binary.Read(file, binary.BigEndian, &cmap.NumTables)
	if err != nil {
		return nil, errors.New("Cannot read number of cmap subtables")
	}

	cmap.Subtables = make([]CmapSubtableHeader, cmap.NumTables)
	for _, subtable := range cmap.Subtables {
		err = binary.Read(file, binary.BigEndian, &subtable.PlatformID)
		if err != nil {
			return nil, errors.New("Cannot read one of the cmap subtables's platform ID")
		}
		err = binary.Read(file, binary.BigEndian, &subtable.EncodingID)
		if err != nil {
			return nil, errors.New("Cannot read one of the cmap subtables's encoding ID")
		}
		err = binary.Read(file, binary.BigEndian, &subtable.Offset)
		if err != nil {
			return nil, errors.New("Cannot read one of the cmap subtables's offset")
		}
	}

	for _, subtable := range cmap.Subtables {
		subtStartingPos, err := file.Seek(0, 1)
		if err != nil {
			return nil, errors.New("Error seeking into one of cmap's subtable")
		}

		_, err = file.Seek(int64(offset)+int64(subtable.Offset), 0)
		if err != nil {
			return nil, errors.New("Error seeking into one of cmap subtables's offset")
		}

		var format uint16
		err = binary.Read(file, binary.BigEndian, &format)
		if err != nil {
			return nil, errors.New("Cannot read cmap's subtable format")
		}

		var length uint16
		err = binary.Read(file, binary.BigEndian, &length)
		if err != nil {
			return nil, errors.New("Cannot read cmap's subtable length")
		}

		_, err = file.Seek(int64(offset)+int64(subtable.Offset), 0)
		if err != nil {
			return nil, errors.New("Error seeking into one of cmap's subtable's data")
		}

		subtable.Data = make([]byte, length)
		_, err = file.Read(subtable.Data)
		if err != nil {
			return nil, errors.New("Cannot read one of cmap's subtable's data")
		}

		_, err = file.Seek(subtStartingPos, 0)
		if err != nil {
			return nil, errors.New("Error seeking to starting position of cmap subtable")
		}
	}

	return cmap, nil
}
