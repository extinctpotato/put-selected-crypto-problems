package visenc

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
)

const (
	WHITE = 255
	BLACK = 0
)

type VisEnc struct {
	InputImage image.Image
}

func (v *VisEnc) LoadFromFile(path string) error {
	imageFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	inputImage, _, err := image.Decode(imageFile)
	if err != nil {
		return err
	}

	// Check if image is in greyscale.
	switch inputImage.ColorModel() {
	case color.GrayModel:
	default:
		return errors.New("image is not in grayscale format")
	}

	// Check if only 0 and 255 values are present.
	imgBounds := inputImage.Bounds()

	for x := 0; x < imgBounds.Max.X; x++ {
		for y := 0; y < imgBounds.Max.Y; y++ {
			colorValue := inputImage.At(y, x).(color.Gray)

			if colorValue.Y > BLACK && colorValue.Y != WHITE {
				return fmt.Errorf("value %d not between %d and %d",
					colorValue.Y, BLACK, WHITE,
				)
			}
		}
	}

	v.InputImage = inputImage

	return nil
}

//func (v *VisEnc) Encode() {
//	inputImageBoundaries := v.InputImage.Bounds()
//}

func (v *VisEnc) Print() {
	imgBounds := v.InputImage.Bounds()

	fmt.Println("bounds:", imgBounds)

	var blacks, whites int

	for x := 0; x < imgBounds.Max.X; x++ {
		for y := 0; y < imgBounds.Max.Y; y++ {
			colorValue := v.InputImage.At(x, y).(color.Gray)

			var colorStr string

			if colorValue.Y == WHITE {
				colorStr = "white"
				whites++
			} else {
				colorStr = "black"
				blacks++
			}

			fmt.Printf("x: %d, y: %d: %s\n", x, y, colorStr)
		}
	}

	fmt.Println("blacks:", blacks)
	fmt.Println("whites:", whites)
}
