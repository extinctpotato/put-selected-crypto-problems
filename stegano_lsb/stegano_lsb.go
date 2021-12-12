package stegano_lsb

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
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

func (s SteganoLsb) Encode(msg string) (image.Image, error) {
	img := cloneToRGBA(s.OriginalImage)
	imgBounds := img.Bounds()

	msgMask := StringToRgbMask(msg)

	fmt.Printf("%d\n", msgMask)

	for x := 0; x < imgBounds.Max.X; x++ {
		for y := 0; y < imgBounds.Max.Y; y++ {
			flatIndex := x*imgBounds.Max.X + y

			oldCol := img.At(y, x).(color.RGBA)

			fmt.Printf("[%d][%d] %d: %d, %d, %d\n",
				x, y, flatIndex, oldCol.R, oldCol.G, oldCol.B)

			if flatIndex < len(msgMask) {
				r2, g2, b2 := ChangeLsbUint8(oldCol.R, msgMask[flatIndex][0]),
					ChangeLsbUint8(oldCol.B, msgMask[flatIndex][1]),
					ChangeLsbUint8(oldCol.G, msgMask[flatIndex][2])
				newColor := color.NRGBA{r2, g2, b2, 255}

				img.Set(y, x, newColor)
			}
		}
	}

	return img, nil
}

func (s SteganoLsb) Decode() string {
	img := cloneToRGBA(s.OriginalImage)
	imgBounds := s.OriginalImage.Bounds()

	var tmpBin string
	var tmpCounter int

	var asciiChars []byte

	for x := 0; x < imgBounds.Max.X; x++ {
		for y := 0; y < imgBounds.Max.Y; y++ {
			col := img.At(y, x).(color.RGBA)
			tmpCounter += 3

			// Extract LSB from the RGB component.
			if tmpCounter < 9 {
				tmpBin += strconv.Itoa(int(col.R % 2))
				tmpBin += strconv.Itoa(int(col.G % 2))
				tmpBin += strconv.Itoa(int(col.B % 2))
			} else if tmpCounter == 9 {
				tmpBin += strconv.Itoa(int(col.R % 2))
				tmpBin += strconv.Itoa(int(col.G % 2))

				tmpChar, _ := strconv.ParseUint(tmpBin, 2, 8)

				fmt.Println(tmpBin)

				if tmpChar == 0 {
					return string(asciiChars)
				}

				asciiChars = append(asciiChars, uint8(tmpChar))

				tmpCounter = 0
				tmpBin = ""
			}
		}
	}

	return string(asciiChars)
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

	// Terminate the string with NUL.
	result = append(result, 0)

	return result
}

func cloneToRGBA(src image.Image) *image.RGBA {
	b := src.Bounds()
	dst := image.NewRGBA(b)
	draw.Draw(dst, b, src, b.Min, draw.Src)

	return dst
}

func ChangeLsbUint8(n uint8, zeroOrOne int) uint8 {
	if zeroOrOne == 0 {
		return clearBitUint8(n, 0)
	} else {
		return setBitUint8(n, 0)
	}
}

// Shamelessly stolen from Kevin Burke:
// https://stackoverflow.com/a/23192263
func setBitUint8(n uint8, pos uint) uint8 {
	n |= (uint8(1) << pos)
	return n
}

func clearBitUint8(n uint8, pos uint) uint8 {
	mask := ^(uint8(1) << pos)
	n &= mask
	return n
}
