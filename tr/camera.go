package tr

import (
	"fmt"
	"math"
)

type Camera interface {
	CastRay(x int, y int, w int, h int, dx, dy float64) Ray
}

type PerspectiveCamera struct {
	position, forward, up, right Vector3
	cFov float64
}

func (c PerspectiveCamera) String() string {
	return fmt.Sprintf("Perspective")
}

func LookAt(pos, look, up Vector3, fov float64) PerspectiveCamera {
	p := pos
	f := look.Sub(pos).Normalize()
	u := up.Normalize()
	r := f.Cross(u).Normalize()
	c := math.Tan(fov*math.Pi/360)
	return PerspectiveCamera{p, f, u, r, c}
}

func (c *PerspectiveCamera) CastRay(x, y, w, h int, dx, dy float64) Ray {
	height := 2*c.cFov
	width := float64(w) * height / float64(h)
	pixelDelta := width / float64(w)
	offsetX := c.right.Mul(pixelDelta * (dx + float64(x)))
	offsetY := c.up.Mul(-pixelDelta * (dy + float64(y)))
	o := c.position
	d := c.position.Add(c.forward).Sub(c.right.Mul(width/2)).Add(c.up.Mul(height/2)).Add(offsetX).Add(offsetY)
	d = d.Sub(o)
	d = d.Normalize()
	return Ray{o, d}
}