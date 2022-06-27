package noise

import ojracnoise "github.com/ojrac/opensimplex-go"

type OjracGenerator struct {
	config OjracGeneratorConfig
}

type OjracGeneratorConfig struct {
	Seed int64
}

func NewOjracGenerator(config OjracGeneratorConfig) *OjracGenerator {
	return &OjracGenerator{config: config}
}

func (o *OjracGenerator) Get1D(x float64) float64 {
	noise := ojracnoise.New(o.config.Seed)
	return noise.Eval2(0, x)
}

func (o *OjracGenerator) Get2D(x, y float64) float64 {
	noise := ojracnoise.New(o.config.Seed)
	return noise.Eval2(x, y)
}

func (o *OjracGenerator) Get3D(x, y, z float64) float64 {
	noise := ojracnoise.New(o.config.Seed)
	return noise.Eval3(x, y, z)
}
