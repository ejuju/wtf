package imgutil

import (
	"image"
	"image/color"
)

func NewMonochromeImage(clr color.Color, width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dx(); y++ {
			img.Set(x, y, clr)
		}
	}
	return img
}
