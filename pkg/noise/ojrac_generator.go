package noise

import ojracnoise "github.com/ojrac/opensimplex-go"

type OjracGenerator struct {
	generator ojracnoise.Noise
}

func NewOjracGenerator(seed int64) *OjracGenerator {
	return &OjracGenerator{generator: ojracnoise.New(seed)}
}

func (g *OjracGenerator) Get3D(x, y, z float64) float64 {
	return g.generator.Eval3(x, y, z)
}
