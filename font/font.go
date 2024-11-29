package font

import (
	"os"

	h "github.com/andrwui/gofonts/header"
	t "github.com/andrwui/gofonts/tables"
)

type Font struct {
	header *h.FontHeader
	head   *t.HeadTable
	name   *t.NameTable

	file         *os.File
	tableOffsets map[string]uint32
}

func ParseFont(file *os.File) (*Font, error) {
	header, err := h.ParseFontHeader(file)
	if err != nil {
		return nil, err
	}

	font := &Font{
		header:       header,
		file:         file,
		tableOffsets: make(map[string]uint32),
	}

	for _, table := range header.TableDirectory {
		font.tableOffsets[string(table.Tag[:])] = table.Offset
	}

	return font, nil
}
