package tr

type Shape interface {
	Intersect(*Ray) float64
	Normal(Vector3) Vector3
	Color(Vector3) Color
	Material(Vector3) Material
}