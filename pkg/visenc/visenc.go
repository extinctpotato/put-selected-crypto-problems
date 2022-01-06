package visenc

import (
	"image"
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

	v.InputImage = inputImage

	return nil
}
