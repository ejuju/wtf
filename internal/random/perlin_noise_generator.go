package random

type PerlinNoiseGenerator interface {
	Get1D(x float64) float64
	Get2D(x, y float64) float64
	Get3D(x, y, z float64) float64
}
