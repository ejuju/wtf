package imgutil

import (
	"image"
	"image/color"

	"github.com/ejuju/wtf/pkg/random"
)

type PerlinNoisePixelModifier struct {
	config PerlinNoisePixelModifierConfig
}

type PerlinNoisePixelModifierConfig struct {
	PerlinNoiseGenerator    random.NoiseGenerator
	Amplitude               float64     // recommended: 100
	OutOfFrameFallbackColor color.Color // recommended: black or white
	PositionGapDivider      float64     // recommended: 50
}

func NewPerlinNoisePixelModifier(config PerlinNoisePixelModifierConfig) *PerlinNoisePixelModifier {
	return &PerlinNoisePixelModifier{config: config}
}

func (p *PerlinNoisePixelModifier) ModifyPixel(img image.Image, point image.Point) color.Color {
	x := float64(point.X)
	y := float64(point.Y)
	noiseval := (p.config.PerlinNoiseGenerator.Get2D(x/p.config.PositionGapDivider, y/p.config.PositionGapDivider)) * p.config.Amplitude

	newX := int(x + noiseval)
	newY := int(y + noiseval)

	// use fallback color if pixel is out of bounds
	bounds := img.Bounds()
	if newX < bounds.Min.X || newX >= bounds.Max.X || newY < bounds.Min.Y || newY >= bounds.Max.Y {
		return p.config.OutOfFrameFallbackColor
	}

	return img.At(newX, newY)
}
