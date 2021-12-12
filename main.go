package main

import (
	"fmt"

	stegano "github.com/extinctpotato/put-selected-crypto-problems-lab/stegano_lsb"
)

func main() {
	fmt.Println("Hello world!")

	s := stegano.SteganoLsb{}
	err := s.LoadFromFile("misc/stairs_500_500.png")

	if err != nil {
		fmt.Printf("err: %s", err)
	}

	fmt.Printf("%s\n", s.OriginalImageFmt)

	s.Encode("test")

	rgbMaskTest := stegano.StringToRgbMask("cool")
	fmt.Printf("%d\n", rgbMaskTest)
}
