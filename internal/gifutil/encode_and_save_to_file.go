package gifutil

import (
	"errors"
	"image/gif"
	"os"
)

//
func EncodeAndSaveToFile(g *gif.GIF, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if g == nil {
		return errors.New("No GIF was provided (empty pointer)")
	}
	return gif.EncodeAll(file, g)
}
