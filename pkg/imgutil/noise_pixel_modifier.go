package imgutil

import (
	"image"
	"image/color"

	"github.com/ejuju/wtf/pkg/noise"
)

type NoisePixelModifier struct {
	config NoisePixelModifierConfig
}

type NoisePixelModifierConfig struct {
	NoiseGenerator          noise.Generator
	Amplitude               float64     // recommended: 100
	OutOfFrameFallbackColor color.Color // recommended: black or white
	PositionGapDivider      float64     // recommended: 50
}

func NewNoisePixelModifier(config NoisePixelModifierConfig) *NoisePixelModifier {
	return &NoisePixelModifier{config: config}
}

func (p *NoisePixelModifier) ModifyPixel(img image.Image, point image.Point) color.Color {
	x := float64(point.X)
	y := float64(point.Y)
	noiseval := (p.config.NoiseGenerator.Get3D(x/p.config.PositionGapDivider, y/p.config.PositionGapDivider, 0)) * p.config.Amplitude

	newX := int(x + noiseval)
	newY := int(y + noiseval)

	// use fallback color if pixel is out of bounds
	bounds := img.Bounds()
	if newX < bounds.Min.X || newX >= bounds.Max.X || newY < bounds.Min.Y || newY >= bounds.Max.Y {
		return p.config.OutOfFrameFallbackColor
	}

	return img.At(newX, newY)
}
