package noise

type Generator interface {
	Get3D(x, y, z float64) float64
}
