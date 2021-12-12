package stegano_lsb

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
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

			rLSB := setBitUint32(r, 0)

			rB, gB, bB, rLSBB := strconv.FormatUint(uint64(r), 2),
				strconv.FormatUint(uint64(g), 2),
				strconv.FormatUint(uint64(b), 2),
				strconv.FormatUint(uint64(rLSB), 2)

			fmt.Printf("\t%s, %s, %s, %s\n", rB, rLSBB, gB, bB)
		}
	}
}

func (s SteganoLsb) Encode(msg string) (image.Image, error) {
	img := cloneToRGBA(s.OriginalImage)
	imgBounds := img.Bounds()

	msgMask := StringToRgbMask(msg)

	fmt.Printf("%d\n", msgMask)

	for x := 0; x < imgBounds.Max.X; x++ {
		for y := 0; y < imgBounds.Max.Y; y++ {
			flatIndex := x*imgBounds.Max.X + y

			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()
			fmt.Printf("[%d][%d] %d: %d, %d, %d\n",
				x, y, flatIndex, r, g, b)

			r2, g2, b2 := ChangeLsbUint32(r, msgMask[flatIndex][0]),
				ChangeLsbUint32(r, msgMask[flatIndex][1]),
				ChangeLsbUint32(r, msgMask[flatIndex][2])

			newColor := color.RGBA{r2, g2, b2, 1}

			img.Set(x, y, newColor)
		}
	}

	return img, nil
}

func StringToRgbMask(s string) [][]int {
	charArray := stringToCharArray(s)
	result := make([][]int, 3*len(charArray))

	for charIdx, char := range charArray {
		binRepr := fmt.Sprintf("%08b0", char)

		// Consme three elements in each iteration.
		for i := 0; i < 9; i += 3 {
			// Divide to get the iteration index.
			resultIdx := i/3 + charIdx*3
			result[resultIdx] = make([]int, 3)

			result[resultIdx][0], _ = strconv.Atoi(string(binRepr[i]))
			result[resultIdx][1], _ = strconv.Atoi(string(binRepr[i+1]))
			result[resultIdx][2], _ = strconv.Atoi(string(binRepr[i+2]))
		}
	}

	return result
}

func stringToCharArray(s string) []int {
	bytes := []byte(s)
	var result []int

	for i := 0; i < len(bytes); i++ {
		result = append(result, int(bytes[i]))
	}

	return result
}

func cloneToRGBA(src image.Image) *image.RGBA {
	b := src.Bounds()
	dst := image.NewRGBA(b)
	draw.Draw(dst, b, src, b.Min, draw.Src)

	return dst
}

func ChangeLsbUint32(n uint32, zeroOrOne int) uint32 {
	if zeroOrOne == 0 {
		return clearBitUint32(n, 0)
	} else {
		return setBitUint32(n, 0)
	}
}

// Shamelessly stolen from Kevin Burke:
// https://stackoverflow.com/a/23192263
func setBitUint32(n uint32, pos uint) uint32 {
	n |= (uint32(1) << pos)
	return n
}

func clearBitUint32(n uint32, pos uint) uint32 {
	mask := ^(uint32(1) << pos)
	n &= mask
	return n
}
