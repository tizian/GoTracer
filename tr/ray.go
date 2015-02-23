package tr

import (
	"fmt"
	"math"
)

type Ray struct {
	Origin Vector3
	Direction Vector3
}

func (r Ray) Position(t float64) Vector3 {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r Ray) String() string {
	return fmt.Sprintf("%s", r.Direction)
}

func (r Ray) Reflect(p, n Vector3) Ray {
	newD := r.Direction.Sub(n.Mul(2*r.Direction.Dot(n))).Normalize()
	return Ray{p, newD}
}

func (r Ray) Refract(p, n Vector3, eta1, eta2 float64) (Ray, bool) {
	d := r.Direction
	dn := d.Dot(n)
	root := 1 - (eta1*eta1)/(eta2*eta2) * (1 - dn * dn)
	if root < 0 {
		return Ray{}, false
	}
	root = math.Sqrt(root)
	newD := d.Sub(n.Mul(dn)).Mul(eta1/eta2).Sub(n.Mul(root)).Normalize()
	return Ray{p, newD}, true
}