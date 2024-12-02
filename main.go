package main

import (
	"fmt"
	"os"

	f "github.com/andrwui/gofonts/font"
)

func main() {

	filePath := "/home/andrw/geist/_geist_100.ttf"

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	font, err := f.ParseFont(file)
	if err != nil {
		panic(err)
	}

	family, err := font.GetFontName(f.NameIDPreferredFamily)
	if err != nil {
		panic(err)
	}
	subfamily, err := font.GetFontName(f.NameIDPreferredSubfamily)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s, %s\n", family, subfamily)

}
