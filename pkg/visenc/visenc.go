package visenc

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
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

			if colorValue.Y > 0 && colorValue.Y != 255 {
				return fmt.Errorf("value %d not between 0 and 255",
					colorValue.Y,
				)
			}
		}
	}

	v.InputImage = inputImage

	return nil
}

func (v *VisEnc) Print() {
	imgBounds := v.InputImage.Bounds()

	fmt.Println("bounds:", imgBounds)

	for x := 0; x < imgBounds.Max.X; x++ {
		for y := 0; y < imgBounds.Max.Y; y++ {
			colorValue := v.InputImage.At(y, x).(color.Gray)

			fmt.Printf("[%d][%d]: [%d]\n", x, y, colorValue.Y)
		}
	}
}
