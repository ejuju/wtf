package random

import "github.com/ojrac/opensimplex-go"

type OpenSimplexNoiseGenerator struct {
	config OpenSimplexNoiseGeneratorConfig
}

type OpenSimplexNoiseGeneratorConfig struct {
	Seed int64
}

func NewOpenSimplexNoiseGenerator(config OpenSimplexNoiseGeneratorConfig) *OpenSimplexNoiseGenerator {
	return &OpenSimplexNoiseGenerator{config: config}
}

func (o *OpenSimplexNoiseGenerator) Get1D(x float64) float64 {
	noise := opensimplex.New(o.config.Seed)
	return noise.Eval2(0, x)
}

func (o *OpenSimplexNoiseGenerator) Get2D(x, y float64) float64 {
	noise := opensimplex.New(o.config.Seed)
	return noise.Eval2(x, y)
}

func (o *OpenSimplexNoiseGenerator) Get3D(x, y, z float64) float64 {
	noise := opensimplex.New(o.config.Seed)
	return noise.Eval3(x, y, z)
}
