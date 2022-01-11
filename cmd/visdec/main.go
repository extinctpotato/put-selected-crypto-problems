package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/extinctpotato/put-selected-crypto-problems-lab/pkg/visenc"
)

func main() {
	v := visenc.VisDec{}

	// Parse command-line parameters.
	inputFilePath := flag.String("in", "", "input image file")
	outputFilePath := flag.String("out", "", "output image file")
	flag.Parse()

	if *inputFilePath == "" {
		fmt.Fprintf(os.Stderr, "provide input file path!")
		os.Exit(1)
	}

	if *outputFilePath == "" {
		fmt.Fprintf(os.Stderr, "provide output file path!")
		os.Exit(1)
	}

	err := v.LoadFromFile(*inputFilePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open %s: %s", *inputFilePath, err)
		os.Exit(1)
	}

	outputImage, err := v.Decode()

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't decode image: %s", err)
		os.Exit(1)
	}

	outputImageFile, err := os.Create(*outputFilePath)

	defer outputImageFile.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't create %s: %s", *outputFilePath, err)
		os.Exit(1)
	}

	png.Encode(outputImageFile, outputImage)
}
