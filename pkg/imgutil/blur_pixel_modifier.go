package imgutil

import (
	"image"
	"image/color"
	"math/rand"
)

type BlurPixelModifier struct {
	config BlurPixelModifierConfig
}

type BlurPixelModifierConfig struct {
	MaxAmplitude int // recommended: 20
}

func NewBlurPixelModifier(config BlurPixelModifierConfig) *BlurPixelModifier {
	return &BlurPixelModifier{config: config}
}

func (p *BlurPixelModifier) ModifyPixel(img image.Image, point image.Point) color.Color {
	newX := point.X + (rand.Intn(2*p.config.MaxAmplitude) - p.config.MaxAmplitude)
	newY := point.Y + (rand.Intn(2*p.config.MaxAmplitude) - p.config.MaxAmplitude)
	return img.At(newX, newY)
}
