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
			// // if in padding zone, skip
			// if x <= maxX/4 || x >= 3*maxX/4 || y <= maxY/4 || y >= 3*maxY/4 {
			// 	newImg.Set(x, y, baseImg.At(x, y))
			// 	continue
			// }

			newImg.Set(x, y, modifier.ModifyPixel(baseImg, image.Point{X: x, Y: y}))
		}
	}

	return newImg
}
