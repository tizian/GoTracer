package tr

import (
	"fmt"
	"math"
)

type Color struct {
	R, G, B float64
}

func ColorFromHex(x int) Color {
	r := float64((x >> 16) & 0xff) / 255
	g := float64((x >> 8) & 0xff) / 255
	b := float64((x >> 0) & 0xff) / 255
	return Color{r, g, b}
}

func ColorFromInts(r, g, b uint8) Color {
	rf := float64(r) / 255
	gf := float64(g) / 255
	bf := float64(b) / 255
	return Color{rf, gf, bf}
}

func (c *Color) ToInts() (r, g, b uint8) {
	r = uint8(c.R * 255)
	g = uint8(c.G * 255)
	b = uint8(c.B * 255)
	return
}

func (c Color) String() string {
	return fmt.Sprintf("(%f, %f, %f)", c.R, c.G, c.B)
}

func (c Color) Clamp() Color {
	r := math.Max(0.0, math.Min(1.0, c.R))
	g := math.Max(0.0, math.Min(1.0, c.G))
	b := math.Max(0.0, math.Min(1.0, c.B))
	return Color{r, g, b}
}

func (c Color) Add(d Color) Color {
	return Color{c.R + d.R, c.G + d.G, c.B + d.B}
}

func (c Color) Sub(d Color) Color {
	return Color{c.R - d.R, c.G - d.G, c.B - d.B}
}

func (c Color) Mul(a float64) Color {
	return Color{c.R * a, c.G * a, c.B * a}
}

func (c Color) Div(a float64) Color {
	return Color{c.R / a, c.G / a, c.B / a}
}

func (c Color) MulC(d Color) Color {
	return Color{c.R * d.R, c.G * d.G, c.B * d.B}
}

func (c Color) DivC(d Color) Color {
	return Color{c.R / d.R, c.G / d.G, c.B / d.B}
}