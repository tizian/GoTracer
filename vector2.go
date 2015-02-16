package main

import (
	"fmt"
	"math"
)

type Vector2 struct {
	x, y float64
}

func (v Vector2) String() string {
	return fmt.Sprintf("(%f, %f)", v.x, v.y)
}

func (v Vector2) Add(u Vector2) Vector2 {
	return Vector2{v.x + u.x, v.y + u.y}
}

func (v Vector2) Sub(u Vector2) Vector2 {
	return Vector2{v.x - u.x, v.y - u.y}
}

func (v Vector2) Mul(a float64) Vector2 {
	return Vector2{v.x * a, v.y * a}
}

func (v Vector2) Div(a float64) Vector2 {
	return Vector2{v.x / a, v.y / a}
}

func (v Vector2) MulC(u Vector2) Vector2 {
	return Vector2{v.x * u.x, v.y * u.y}
}

func (v Vector2) DivC(u Vector2) Vector2 {
	return Vector2{v.x / u.x, v.y / u.y}
}

func (v Vector2) Length() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func (v Vector2) Length2() float64 {
	return v.x * v.x + v.y * v.y
}

func (v Vector2) Dot(u Vector2) float64 {
	return v.x * u.x + v.y * u.y
}

func (v Vector2) Normalize() Vector2 {
	a := v.Length()
	return Vector2{v.x / a, v.y / a}
}