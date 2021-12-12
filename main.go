package main

import (
	"fmt"
	"image/png"
	"os"

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

	enc, _ := s.Encode("test")

	encodedFile, err := os.Create("misc/stairs_encoded.png")
	defer encodedFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	png.Encode(encodedFile, enc)
}
