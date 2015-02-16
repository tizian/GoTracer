package main

type Plane struct {
	point Vector3
	normal Vector3
	material Material
}

func CreatePlane(point, normal Vector3, material Material) SceneObject {
	return &Plane{point, normal.Normalize(), material}
}

func (p *Plane) Intersect(r *Ray) float64 {
	denom := r.direction.Dot(p.normal)

	if denom > EPS {
		return INF
	}
	t := p.point.Sub(r.origin).Dot(p.normal) / denom
	if t < EPS {
		return INF
	}
	return t
}

func (p *Plane) Color(v Vector3) Color {
	return p.material.color
}

func (p *Plane) Normal(v Vector3) Vector3 {
	return p.normal
}

func (p *Plane) Material(v Vector3) Material {
	return p.material
}