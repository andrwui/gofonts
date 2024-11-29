package tables

import (
	"encoding/binary"
	"errors"
	"os"
)

type HeadTable struct {
	MajorVersion       uint16 `json:"major_version"`
	MinorVersion       uint16 `json:"minor_version"`
	FontRevision       int32  `json:"font_revision"`
	CheckSumAdjustment uint32 `json:"check_sum_adjustment"`
	MagicNumber        uint32 `json:"magic_number"`
	Flags              uint16 `json:"flags"`
	UnitsPerEm         uint16 `json:"units_per_em"`
	Created            int64  `json:"created"`
	Modified           int64  `json:"modified"`
	XMin               int16  `json:"x_min"`
	YMin               int16  `json:"y_min"`
	XMax               int16  `json:"x_max"`
	YMax               int16  `json:"y_max"`
	MacStyle           uint16 `json:"mac_style"`
	LowestRecPpem      uint16 `json:"lowest_rec_ppem"`
	FontDirectionHint  int16  `json:"font_direction_hint"`
	IndexToLocFormat   int16  `json:"index_to_loc_format"`
	GlyphDataFormat    int16  `json:"glyph_data_format"`
}

func ReadHeadTable(file *os.File, offset uint32) (*HeadTable, error) {

	head := &HeadTable{}

	_, err := file.Seek(int64(offset), 0)
	err = binary.Read(file, binary.BigEndian, head)
	if err != nil {
		return nil, errors.New("This table does not exists in the current font file.")
	}

	return head, nil

}
