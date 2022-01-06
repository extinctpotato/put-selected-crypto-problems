package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/extinctpotato/put-selected-crypto-problems-lab/pkg/visenc"
)

func main() {
	v := visenc.VisEnc{}
	fmt.Println("vim-go")

	// Parse command-line parameters.
	inputFilePath := flag.String("in", "", "input image file")
	flag.Parse()

	if *inputFilePath == "" {
		fmt.Fprintf(os.Stderr, "provide input file path!")
		os.Exit(1)
	}

	err := v.LoadFromFile(*inputFilePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open %s: %s", *inputFilePath, err)
		os.Exit(1)
	}
}
