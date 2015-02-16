package  main

import (
	"fmt"
	"math"
)

type Color struct {
	r, g, b float64
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
	r = uint8(c.r * 255)
	g = uint8(c.g * 255)
	b = uint8(c.b * 255)
	return
}

func (c Color) String() string {
	return fmt.Sprintf("(%f, %f, %f)", c.r, c.g, c.b)
}

func (c Color) Clamp() Color {
	r := math.Max(0.0, math.Min(1.0, c.r))
	g := math.Max(0.0, math.Min(1.0, c.g))
	b := math.Max(0.0, math.Min(1.0, c.b))
	return Color{r, g, b}
}

func (c Color) Add(d Color) Color {
	return Color{c.r + d.r, c.g + d.g, c.b + d.b}
}

func (c Color) Sub(d Color) Color {
	return Color{c.r - d.r, c.g - d.g, c.b - d.b}
}

func (c Color) Mul(a float64) Color {
	return Color{c.r * a, c.g * a, c.b * a}
}

func (c Color) Div(a float64) Color {
	return Color{c.r / a, c.g / a, c.b / a}
}

func (c Color) MulC(d Color) Color {
	return Color{c.r * d.r, c.g * d.g, c.b * d.b}
}

func (c Color) DivC(d Color) Color {
	return Color{c.r / d.r, c.g / d.g, c.b / d.b}
}