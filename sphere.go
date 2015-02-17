package main

import "math"

type Sphere struct {
	center Vector3
	radius float64
	material Material
}

func CreateSphere(center Vector3, radius float64, material Material) Shape {
	return &Sphere{center, radius, material}
}

func (s *Sphere) Intersect(ray *Ray) float64 {
	o := ray.origin
	d := ray.direction
	c := s.center
	r := s.radius

	A := d.x * d.x + d.y * d.y + d.z * d.z
	B := 2 * (o.x * d.x - d.x * c.x + o.y * d.y - d.y * c.y + o.z * d.z - d.z * c.z)
	C := (o.x * o.x - 2 * o.x * c.x + c.x * c.x) + (o.y * o.y - 2 * o.y * c.y + c.y * c.y) + (o.z * o.z - 2 * o.z * c.z + c.z * c.z) - r * r;

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
	return s.material.color
}

func (s *Sphere) Normal(v Vector3) Vector3 {
	return (v.Sub(s.center)).Normalize()
}

func (s *Sphere) Material(v Vector3) Material {
	return s.material
}