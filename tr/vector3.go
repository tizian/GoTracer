package tr

import (
	"fmt"
	"math"
)

type Vector3 struct {
	X, Y, Z float64
}

func (v Vector3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}

func (v Vector3) Add(u Vector3) Vector3 {
	return Vector3{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v Vector3) Sub(u Vector3) Vector3 {
	return Vector3{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v Vector3) Mul(a float64) Vector3 {
	return Vector3{v.X * a, v.Y * a, v.Z * a}
}

func (v Vector3) Div(a float64) Vector3 {
	return Vector3{v.X / a, v.Y / a, v.Z / a}
}

func (v Vector3) MulC(u Vector3) Vector3 {
	return Vector3{v.X * u.X, v.Y * u.Y, v.Z * u.Z}
}

func (v Vector3) DivC(u Vector3) Vector3 {
	return Vector3{v.X / u.X, v.Y / u.Y, v.Z / u.Z}
}

func (v Vector3) Length() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func (v Vector3) Length2() float64 {
	return v.X * v.X + v.Y * v.Y + v.Z * v.Z
}

func (v Vector3) Dot(u Vector3) float64 {
	return v.X * u.X + v.Y * u.Y + v.Z * u.Z
}

func (v Vector3) Cross(u Vector3) Vector3 {
	resx := v.Y * u.Z - v.Z * u.Y
	resy := v.Z * u.X - v.X * u.Z
	resz := v.X * u.Y - v.Y * u.X
	return Vector3{resx, resy, resz}
}

func (v Vector3) Normalize() Vector3 {
	a := v.Length()
	return Vector3{v.X / a, v.Y / a, v.Z / a}
}