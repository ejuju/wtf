package gifutil

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
)

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
