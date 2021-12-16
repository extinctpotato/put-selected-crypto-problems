package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/extinctpotato/put-selected-crypto-problems-lab/pkg/stegano_lsb"
)

func main() {
	// Declare structs used for processing.
	s := stegano_lsb.SteganoLsb{}

	// Parse command-line parameters.
	inputFilePath := flag.String("in", "", "input image file")
	outputFilePath := flag.String("out", "./out.png", "output image file")
	message := flag.String("text", "/dev/stdin", "text to be encoded")
	flag.Parse()

	if *inputFilePath == "" {
		fmt.Fprintf(os.Stderr, "provide input file path!")
		os.Exit(1)
	}

	err := s.LoadFromFile(*inputFilePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open %s: %s", *inputFilePath, err)
		os.Exit(1)
	}

	alteredImg, err := s.Encode(*message)

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't encode the message: %s", err)
		os.Exit(1)
	}

	alteredImgFile, err := os.Create(*outputFilePath)
	defer alteredImgFile.Close()

	err = png.Encode(alteredImgFile, alteredImg)

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't encode PNG: %s", err)
		os.Exit(1)
	}
}
