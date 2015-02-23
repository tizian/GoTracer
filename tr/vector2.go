package tr

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X, Y float64
}

func (v Vector2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}

func (v Vector2) Add(u Vector2) Vector2 {
	return Vector2{v.X + u.X, v.Y + u.Y}
}

func (v Vector2) Sub(u Vector2) Vector2 {
	return Vector2{v.X - u.X, v.Y - u.Y}
}

func (v Vector2) Mul(a float64) Vector2 {
	return Vector2{v.X * a, v.Y * a}
}

func (v Vector2) Div(a float64) Vector2 {
	return Vector2{v.X / a, v.Y / a}
}

func (v Vector2) MulC(u Vector2) Vector2 {
	return Vector2{v.X * u.X, v.Y * u.Y}
}

func (v Vector2) DivC(u Vector2) Vector2 {
	return Vector2{v.X / u.X, v.Y / u.Y}
}

func (v Vector2) Length() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v Vector2) Length2() float64 {
	return v.X * v.X + v.Y * v.Y
}

func (v Vector2) Dot(u Vector2) float64 {
	return v.X * u.X + v.Y * u.Y
}

func (v Vector2) Normalize() Vector2 {
	a := v.Length()
	return Vector2{v.X / a, v.Y / a}
}