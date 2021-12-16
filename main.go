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
	err := s.LoadFromFile("misc/stairs_32_32.png")

	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fmt.Printf("%s\n", s.OriginalImageFmt)

	enc, _ := s.Encode("this is a very sophisticated string which is rather long and has a lot of characters")

	encodedFile, err := os.Create("misc/stairs_encoded.png")
	defer encodedFile.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	png.Encode(encodedFile, enc)

	s2 := stegano.SteganoLsb{}
	err = s2.LoadFromFile("misc/stairs_encoded.png")

	if err != nil {
		fmt.Printf("err: %s", err)
	}

	dec := s2.Decode()

	fmt.Printf("|%s|\n", dec)
}
