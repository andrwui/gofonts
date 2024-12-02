package font

import (
	"errors"

	t "github.com/andrwui/gofonts/tables"
)

func (f *Font) Head() (*t.HeadTable, error) {

	offset, exists := f.tableOffsets["head"]
	if !exists {
		return nil, errors.New("Table \"head\" does not exists in the current font file.")
	}

	if f.head == nil {
		head, err := t.ReadHeadTable(f.file, offset)

		if err != nil {
			return nil, err
		}

		f.head = head

	}
	return f.head, nil
}

func (f *Font) Name() (*t.NameTable, error) {

	offset, exists := f.tableOffsets["name"]
	if !exists {
		return nil, errors.New("Table \"name\" does not exists in the current font file.")
	}

	if f.name == nil {
		name, err := t.ReadNameTable(f.file, offset)

		if err != nil {
			return nil, err
		}

		f.name = name

	}
	return f.name, nil
}
