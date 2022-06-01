package imgutil

import (
	"image"
)

type ImageModificationOptions struct {
	Padding float64 // between 0 and 1
}

func ModifyImage(baseImg image.Image, modifier PixelModifier, opts ImageModificationOptions) image.Image {
	bounds := baseImg.Bounds()
	maxX := bounds.Dx()
	maxY := bounds.Dy()
	newImg := image.NewRGBA(image.Rect(0, 0, maxX, maxY))

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			// if in padding zone, skip
			if inPaddingZone(x, y, maxX, maxY, opts.Padding) {
				newImg.Set(x, y, baseImg.At(x, y))
				continue
			}

			newImg.Set(x, y, modifier.ModifyPixel(baseImg, image.Point{X: x, Y: y}))
		}
	}

	return newImg
}

func inPaddingZone(x, y, maxX, maxY int, padding float64) bool {
	return float64(x) < float64(maxX)*padding || float64(x) >= (1-padding)*float64(maxX) ||
		float64(y) < float64(maxY)*padding || float64(y) >= (1-padding)*float64(maxY)
}
