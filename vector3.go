package main

import (
	"fmt"
	"math"
)

type Vector3 struct {
	x, y, z float64
}

func (v Vector3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.x, v.y, v.z)
}

func (v Vector3) Add(u Vector3) Vector3 {
	return Vector3{v.x + u.x, v.y + u.y, v.z + u.z}
}

func (v Vector3) Sub(u Vector3) Vector3 {
	return Vector3{v.x - u.x, v.y - u.y, v.z - u.z}
}

func (v Vector3) Mul(a float64) Vector3 {
	return Vector3{v.x * a, v.y * a, v.z * a}
}

func (v Vector3) Div(a float64) Vector3 {
	return Vector3{v.x / a, v.y / a, v.z / a}
}

func (v Vector3) MulC(u Vector3) Vector3 {
	return Vector3{v.x * u.x, v.y * u.y, v.z * u.z}
}

func (v Vector3) DivC(u Vector3) Vector3 {
	return Vector3{v.x / u.x, v.y / u.y, v.z / u.z}
}

func (v Vector3) Length() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y + v.z * v.z)
}

func (v Vector3) Length2() float64 {
	return v.x * v.x + v.y * v.y + v.z * v.z
}

func (v Vector3) Dot(u Vector3) float64 {
	return v.x * u.x + v.y * u.y + v.z * u.z
}

func (v Vector3) Cross(u Vector3) Vector3 {
	resx := v.y * u.z - v.z * u.y
	resy := v.z * u.x - v.x * u.z
	resz := v.x * u.y - v.y * u.x
	return Vector3{resx, resy, resz}
}

func (v Vector3) Normalize() Vector3 {
	a := v.Length()
	return Vector3{v.x / a, v.y / a, v.z / a}
}