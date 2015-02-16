package main

import (
	"fmt"
	"math"
)

type Ray struct {
	origin Vector3
	direction Vector3
}

func (r Ray) Position(t float64) Vector3 {
	return r.origin.Add(r.direction.Mul(t))
}

func (r Ray) String() string {
	return fmt.Sprintf("%s", r.direction)
}

func (r Ray) Reflect(p, n Vector3) Ray {
	newD := r.direction.Sub(n.Mul(2*r.direction.Dot(n))).Normalize()
	return Ray{p, newD}
}

func (r Ray) Refract(p, n Vector3, eta1, eta2 float64) (Ray, bool) {
	d := r.direction
	dn := d.Dot(n)
	root := 1 - (eta1*eta1)/(eta2*eta2) * (1 - dn * dn)
	if root < 0 {
		return Ray{}, false
	}
	root = math.Sqrt(root)
	newD := d.Sub(n.Mul(dn)).Mul(eta1/eta2).Sub(n.Mul(root)).Normalize()
	return Ray{p, newD}, true
}