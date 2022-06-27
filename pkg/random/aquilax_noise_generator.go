package random

import "github.com/aquilax/go-perlin"

type AquilaxNoiseGenerator struct {
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
type AquilaxNoiseGeneratorConfig struct {
	Alpha float64
	Beta  float64
	N     int32
	Seed  int64
}

func NewAquilaxNoiseGenerator(config AquilaxNoiseGeneratorConfig) *AquilaxNoiseGenerator {
	return &AquilaxNoiseGenerator{
		p: perlin.NewPerlin(config.Alpha, config.Beta, config.N, config.Seed),
	}
}

func (a *AquilaxNoiseGenerator) Get1D(x float64) float64 {
	return a.p.Noise1D(x)
}

func (a *AquilaxNoiseGenerator) Get2D(x, y float64) float64 {
	return a.p.Noise2D(x, y)
}

func (a *AquilaxNoiseGenerator) Get3D(x, y, z float64) float64 {
	return a.p.Noise3D(x, y, z)
}
