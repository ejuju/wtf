package gifutil

import (
	"bytes"
	"errors"
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

//
func encodeImage(img image.Image) (*image.Paletted, error) {
	// encode image in gif format in bytes buffer
	buf := bytes.Buffer{}
	err := gif.Encode(&buf, img, nil)
	if err != nil {
		return nil, err
	}

	// decode buffer to gif type
	decodedImg, err := gif.Decode(&buf)
	if err != nil {
		return nil, err
	}

	// assert to paletted image
	palettedImg, ok := decodedImg.(*image.Paletted)
	if ok == false {
		return nil, errors.New("unable to assert type image.Image to image.Paletted")
	}

	return palettedImg, nil
}
