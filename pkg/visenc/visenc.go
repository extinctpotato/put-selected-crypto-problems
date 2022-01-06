package visenc

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"math/rand"
	"os"
)

const (
	WHITE = 255
	BLACK = 0
)

type VisEnc struct {
	InputImage     image.Image
	InputImageDimX int
	InputImageDimY int
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
	v.InputImageDimX, v.InputImageDimY = imgBounds.Max.X, imgBounds.Max.Y

	for x := 0; x < v.InputImageDimX; x++ {
		for y := 0; y < v.InputImageDimY; y++ {
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

func (v *VisEnc) Encode() []*image.Gray {
	shareImageBoundaries := image.Rectangle{
		image.Point{0, 0},
		image.Point{v.InputImageDimX * 2, v.InputImageDimY},
	}

	shares := []*image.Gray{
		image.NewGray(shareImageBoundaries),
		image.NewGray(shareImageBoundaries),
	}

	_ = shares

	fmt.Println("shareImageBoundaries:", shareImageBoundaries)

	for x := 0; x < v.InputImageDimX; x++ {
		for y := 0; y < v.InputImageDimY; y++ {
			colorValue := v.InputImage.At(x, y).(color.Gray)
			_ = colorValue

			randValue := rand.Float64()

			fmt.Printf("randValue: %f\n", randValue)

			switch colorValue.Y {
			case WHITE:
				fmt.Println("white!")
				if randValue > 0.5 {
					shares[0].Set(x*2, y, color.Gray{WHITE})
					shares[0].Set(x*2+1, y, color.Gray{BLACK})
					shares[1].Set(x*2, y, color.Gray{WHITE})
					shares[1].Set(x*2+1, y, color.Gray{BLACK})
				} else {
					shares[0].Set(x*2, y, color.Gray{BLACK})
					shares[0].Set(x*2+1, y, color.Gray{WHITE})
					shares[1].Set(x*2, y, color.Gray{BLACK})
					shares[1].Set(x*2+1, y, color.Gray{WHITE})
				}
			case BLACK:
				fmt.Println("black!")
				if randValue > 0.5 {
					shares[0].Set(x*2, y, color.Gray{WHITE})
					shares[0].Set(x*2+1, y, color.Gray{BLACK})
					shares[1].Set(x*2, y, color.Gray{BLACK})
					shares[1].Set(x*2+1, y, color.Gray{WHITE})
				} else {
					shares[0].Set(x*2, y, color.Gray{BLACK})
					shares[0].Set(x*2+1, y, color.Gray{WHITE})
					shares[1].Set(x*2, y, color.Gray{WHITE})
					shares[1].Set(x*2+1, y, color.Gray{BLACK})
				}
			}
		}
	}

	fmt.Println(shares[0])
	fmt.Println(shares[1])

	return shares
}

func (v *VisEnc) Print() {
	fmt.Printf("w,h (x,y): %d, %d\n", v.InputImageDimX, v.InputImageDimY)

	var blacks, whites int

	for x := 0; x < v.InputImageDimX; x++ {
		for y := 0; y < v.InputImageDimY; y++ {
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
