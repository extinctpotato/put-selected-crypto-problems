package visenc

import (
	"errors"
	"fmt"
	"image"
	"os"
	"path/filepath"
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
	//  - check if dimensions of both shares are the same
	//  - check if only B/W values present
	//  - check if format is greyscale

	return nil
}

func (v *VisEnc) Decode() error {
	return nil
}
