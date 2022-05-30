package imgutil

import (
	"image"
)

func ModifyImage(baseImg image.Image, modifier PixelModifier) image.Image {
	bounds := baseImg.Bounds()
	maxX := bounds.Dx()
	maxY := bounds.Dy()
	newImg := image.NewRGBA(image.Rect(0, 0, maxX, maxY))

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			newImg.Set(x, y, modifier.ModifyPixel(baseImg, image.Point{X: x, Y: y}))
		}
	}

	return newImg
}
