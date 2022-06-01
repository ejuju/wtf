package gifutil

import (
	"errors"
	"image/gif"
	"os"
)

//
func EncodeAndSaveToFile(g *gif.GIF, path string) error {
	if g == nil {
		return errors.New("No GIF was provided (empty pointer)")
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return gif.EncodeAll(file, g)
}
