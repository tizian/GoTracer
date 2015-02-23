package shapes

import (
	"math"
	. "github.com/tizian/GoTracer/tr"
)

type Sphere struct {
	center Vector3
	radius float64
	material Material
}

func CreateSphere(center Vector3, radius float64, material Material) Shape {
	return &Sphere{center, radius, material}
}

func (s *Sphere) Intersect(ray *Ray) float64 {
	o := ray.Origin
	d := ray.Direction
	c := s.center
	r := s.radius

	A := d.X * d.X + d.Y * d.Y + d.Z * d.Z
	B := 2 * (o.X * d.X - d.X * c.X + o.Y * d.Y - d.Y * c.Y + o.Z * d.Z - d.Z * c.Z)
	C := (o.X * o.X - 2 * o.X * c.X + c.X * c.X) + (o.Y * o.Y - 2 * o.Y * c.Y + c.Y * c.Y) + (o.Z * o.Z - 2 * o.Z * c.Z + c.Z * c.Z) - r * r;

	discr := B * B - 4 * A * C
	if discr < 0 {
		return INF
	}

	root := math.Sqrt(discr)
	t1 := (-B + root) / (2 * A)
	t2 := (-B - root) / (2 * A)

	t := INF
	if t1 > EPS && t1 < t2 {
		t = t1
	} else if t2 > EPS && t2 < t1 {
		t = t2
	} else if t1 > EPS && t2 <= EPS {
		t = t1
	} else if t2 > EPS && t1 <= EPS {
		t = t2
	}
	return t
}

func (s *Sphere) Color(v Vector3) Color {
	return s.material.Color()
}

func (s *Sphere) Normal(v Vector3) Vector3 {
	return (v.Sub(s.center)).Normalize()
}

func (s *Sphere) Material(v Vector3) Material {
	return s.material
}