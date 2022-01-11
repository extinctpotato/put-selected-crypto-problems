package visenc

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

type VisDec struct {
	Shares         []image.Image
	ShareImageDimX int
	ShareImageDimY int
}

// path is the path to any of the shares
func (v *VisDec) LoadFromFile(path string) error {
	shareFileExt := filepath.Ext(path)
	splitShareName := strings.Split(
		strings.TrimSuffix(path, shareFileExt),
		".",
	)

	if _, err := strconv.Atoi(splitShareName[len(splitShareName)-1]); err != nil {
		return errors.New("input file does not end in <int>.<ext>")
	}

	v.Shares = make([]image.Image, 2)

	for i := 1; i < 3; i++ {
		imageFile, err := os.Open(
			fmt.Sprintf("%s.%d%s",
				splitShareName[0], i, shareFileExt,
			),
		)

		defer imageFile.Close()

		if err != nil {
			return err
		}

		shareImage, _, err := image.Decode(imageFile)

		if err != nil {
			return err
		}

		bounds := shareImage.Bounds()

		if v.ShareImageDimX == 0 {
			v.ShareImageDimX = bounds.Max.X
		} else if v.ShareImageDimX != 0 && v.ShareImageDimX != bounds.Max.X {
			return fmt.Errorf("x dim %s should be %d for %d",
				v.ShareImageDimX, bounds.Max.X, i,
			)
		}

		if v.ShareImageDimY == 0 {
			v.ShareImageDimY = bounds.Max.Y
		} else if v.ShareImageDimY != 0 && v.ShareImageDimY != bounds.Max.Y {
			return fmt.Errorf("x dim %s should be %d for %d",
				v.ShareImageDimX, bounds.Max.Y, i,
			)
		}

		v.Shares[i-1] = shareImage
	}

	// TODO: add the following checks:
	//  - check if only B/W values present
	//  - check if format is greyscale

	return nil
}

func (v *VisDec) Decode() (*image.Gray, error) {
	inputImageDimX, inputImageDimY := v.ShareImageDimX/2, v.ShareImageDimY

	inputImage := image.NewGray(image.Rectangle{
		image.Point{0, 0},
		image.Point{inputImageDimX, inputImageDimY},
	})

	whiteComb := [][]uint8{
		{BLACK, WHITE, BLACK, WHITE},
		{WHITE, BLACK, WHITE, BLACK},
	}

	blackComb := [][]uint8{
		{BLACK, WHITE, WHITE, BLACK},
		{WHITE, BLACK, BLACK, WHITE},
	}

	for x := 0; x < inputImageDimX; x++ {
		for y := 0; y < inputImageDimY; y++ {
			shareColors := []uint8{
				v.Shares[0].At(x*2, y).(color.Gray).Y,
				v.Shares[0].At(x*2+1, y).(color.Gray).Y,
				v.Shares[1].At(x*2, y).(color.Gray).Y,
				v.Shares[1].At(x*2+1, y).(color.Gray).Y,
			}

			for _, combVal := range whiteComb {
				if reflect.DeepEqual(combVal, shareColors) {
					inputImage.Set(x, y, color.Gray{WHITE})
				}
			}

			for _, combVal := range blackComb {
				if reflect.DeepEqual(combVal, shareColors) {
					inputImage.Set(x, y, color.Gray{BLACK})
				}
			}
		}
	}

	return inputImage, nil
}
