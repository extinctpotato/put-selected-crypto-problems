package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/extinctpotato/put-selected-crypto-problems-lab/stegano_lsb"
)

func main() {
	// Declare structs used for processing.
	s := stegano_lsb.SteganoLsb{}

	// Parse command-line parameters.
	inputFilePath := flag.String("in", "", "input image file")
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

	decodedMsg := s.Decode()

	fmt.Println(decodedMsg)
}
