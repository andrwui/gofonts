package main

import (
	"fmt"
	"os"

	h "github.com/andrwui/gofonts/header"
	t "github.com/andrwui/gofonts/tables"
)

func main() {

	filePath := "/home/andrw/geist/_geist_regular.ttf"

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rawFont, err := h.ParseFontHeader(file)
	if err != nil {
		panic(err)
	}

	for _, table := range rawFont.TableDirectory {
		tableTag := string(table.Tag[:])
		fmt.Printf("Table tag: %v\n", tableTag)

		switch tableTag {
		case "head":
			_, err := t.ReadHeadTable(file, table.Offset)
			if err != nil {
				panic(err)
			}

			break

		case "name":

			_, err := t.ReadNameTable(file, table.Offset)
			if err != nil {
				panic(err)
			}

		}

	}

}
