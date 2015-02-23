package shapes

import . "github.com/tizian/GoTracer/tr"

type Triangle struct {
	v1, v2, v3 Vector3
	n1, n2, n3 Vector3
	material Material
}

func (t *Triangle) Color(v Vector3) Color {
	return t.material.Color()
}

func (t *Triangle) Normal(v Vector3) Vector3 {
	return t.n1	// TODO(tizian): replace with barycentric interpolation
}

func (t *Triangle) Material(v Vector3) Material {
	return t.material
}

func (tri *Triangle) Intersect(ray *Ray) float64 {
	// Möller-Trumbore intersection algorithm
	e1 := tri.v2.Sub(tri.v1)
	e2 := tri.v3.Sub(tri.v1)
	P := ray.Direction.Cross(e2)
	det := e1.Dot(P)

	if det > -EPS && det < EPS {
		return INF	// Ray lies in plane of triangle
	}

	inv_det := 1 / det

	T := ray.Origin.Sub(tri.v1)
	u := T.Dot(P) * inv_det

	if u < 0 || u > 1 {
		return INF
	}

	Q := T.Cross(e1)
	v := ray.Direction.Dot(Q) * inv_det

	if v < 0 || u + v > 1 {
		return INF
	}

	t := e2.Dot(Q) * inv_det

	if t > EPS {
		return t
	}

	return INF
}