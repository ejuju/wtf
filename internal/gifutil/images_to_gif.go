package gifutil

import (
	"image"
	"image/gif"
)

// Images must be of the same size
func ImagesToGIF(delay int, imgs ...image.Image) (*gif.GIF, error) {
	result := &gif.GIF{}
	for _, img := range imgs {
		// add gif image to result
		var err error
		result, err = addImageToGIF(result, img, delay)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

//
func addImageToGIF(g *gif.GIF, img image.Image, delay int) (*gif.GIF, error) {
	paletted, err := encodeImage(img) // convert regular image to paletted image for gif
	if err != nil {
		return nil, err
	}
	g.Delay = append(g.Delay, delay)
	g.Image = append(g.Image, paletted)
	return g, nil
}
