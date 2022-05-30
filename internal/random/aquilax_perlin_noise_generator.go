package random

import "github.com/aquilax/go-perlin"

type AquilaxPerlinNoiseGenerator struct {
	p *perlin.Perlin
}

// - Alpha: is the weight when the sum is formed,
// usually around 2 (as this approaches 1 the function is noisier)
//
// - Beta: is the harmonic scaling/spacing,
// usually around 2
//
// - N is the number of iterations
//
// - Seed is the math.rand seed value to use
type AquilaxPerlinNoiseGeneratorConfig struct {
	Alpha float64
	Beta  float64
	N     int32
	Seed  int64
}

func NewAquilaxPerlinNoiseGenerator(config AquilaxPerlinNoiseGeneratorConfig) *AquilaxPerlinNoiseGenerator {
	return &AquilaxPerlinNoiseGenerator{
		p: perlin.NewPerlin(config.Alpha, config.Beta, config.N, config.Seed),
	}
}

func (a *AquilaxPerlinNoiseGenerator) Get1D(x float64) float64 {
	return a.p.Noise1D(x)
}

func (a *AquilaxPerlinNoiseGenerator) Get2D(x, y float64) float64 {
	return a.p.Noise2D(x, y)
}

func (a *AquilaxPerlinNoiseGenerator) Get3D(x, y, z float64) float64 {
	return a.p.Noise3D(x, y, z)
}
