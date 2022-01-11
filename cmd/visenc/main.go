package main

import (
	"flag"
	"fmt"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
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

	shares := v.Encode()

	// Get basename, e.g. "misc/small_binary.png" gives "misc/small_binary"
	shareBaseName := strings.TrimSuffix(*inputFilePath,
		filepath.Ext(*inputFilePath),
	)

	// Write shares (share 1 and share 2)
	for i := 1; i < 3; i++ {
		shareFile, err := os.Create(
			fmt.Sprintf("%s.%d%s",
				shareBaseName, i, filepath.Ext(*inputFilePath),
			),
		)

		defer shareFile.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't create %s: %s",
				*shareFile, err,
			)
			os.Exit(1)
		}

		png.Encode(shareFile, shares[i-1])
	}
}
