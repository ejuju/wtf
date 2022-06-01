package imgutil

import (
	"image"
	"image/color"
)

type PixelModifier interface {
	ModifyPixel(image.Image, image.Point) color.Color
}
