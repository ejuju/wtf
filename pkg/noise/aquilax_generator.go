package noise

import aquilaxnoise "github.com/aquilax/go-perlin"

type AquilaxGenerator struct {
	p *aquilaxnoise.Perlin
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
type AquilaxGeneratorConfig struct {
	Alpha float64
	Beta  float64
	N     int32
	Seed  int64
}

func NewAquilaxGenerator(config AquilaxGeneratorConfig) *AquilaxGenerator {
	if config.Alpha <= 0 {
		config.Alpha = 2
	}
	if config.Beta <= 0 {
		config.Beta = 2
	}
	if config.N <= 0 {
		config.N = 1
	}

	return &AquilaxGenerator{
		p: aquilaxnoise.NewPerlin(config.Alpha, config.Beta, config.N, config.Seed),
	}
}

func (a *AquilaxGenerator) Get3D(x, y, z float64) float64 {
	return a.p.Noise3D(x, y, z)
}
