package main

import (
	"flag"
	"fmt"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/extinctpotato/put-selected-crypto-problems-lab/pkg/visenc"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	v := visenc.VisEnc{}

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

	v.Print()
	shares := v.Encode()

	share1File, _ := os.Create("/tmp/1.png")
	share2File, _ := os.Create("/tmp/2.png")

	defer share1File.Close()
	defer share2File.Close()

	png.Encode(share1File, shares[0])
	png.Encode(share2File, shares[1])
}
