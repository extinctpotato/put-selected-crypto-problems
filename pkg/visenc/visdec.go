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
	Shares []image.Image
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

		v.Shares[i-1] = shareImage
	}

	fmt.Println(splitShareName)

	return nil
}
