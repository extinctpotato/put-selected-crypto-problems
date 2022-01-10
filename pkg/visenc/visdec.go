package visenc

import (
	"fmt"
	"image"
	"path/filepath"
	"strings"
)

type VisDec struct {
	Shares []image.Image
}

// path is the path to any of the shares
func (v *VisDec) LoadFromFile(path string) error {
	splitShareName := strings.Split(
		strings.TrimSuffix(path, filepath.Ext(path)),
		".",
	)

	fmt.Println(splitShareName)

	return nil
}
