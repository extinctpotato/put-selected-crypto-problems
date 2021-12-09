package stegano_lsb

import (
	"fmt"
	"image"
	_ "image/png"
	"math/bits"
	"os"
	"strconv"
)

type SteganoLsb struct {
	OriginalImage    image.Image
	OriginalImageFmt string
}

func (s *SteganoLsb) LoadFromFile(path string) error {
	imageFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	originalImage, originalImageFmt, err := image.Decode(imageFile)
	if err != nil {
		return err
	}

	s.OriginalImage = originalImage
	s.OriginalImageFmt = originalImageFmt

	return nil
}

func (s SteganoLsb) PrintColors() {
	bounds := s.OriginalImage.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			color := s.OriginalImage.At(x, y)
			r, g, b, _ := color.RGBA()
			fmt.Printf("[%d][%d]: %d[%d], %d[%d], %d[%d]\n",
				x, y, r, bits.Len32(r), g, bits.Len32(g), b, bits.Len32(b))

			rLSB := r
			rLSB |= (uint32(1) << 0)

			rB, gB, bB, rLSBB := strconv.FormatUint(uint64(r), 2),
				strconv.FormatUint(uint64(g), 2),
				strconv.FormatUint(uint64(b), 2),
				strconv.FormatUint(uint64(rLSB), 2)

			fmt.Printf("\t%s, %s, %s, %s\n", rB, rLSBB, gB, bB)
		}
	}
}