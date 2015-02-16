package main

type Material struct {
	color Color
	index float64
	transparent bool
	specular bool
}

func DiffuseMaterial(c Color) Material {
	return Material{c, -1, false, false}
}

func SpecularMaterial(c Color, index float64) Material {
	return Material{c, index, false, true}
}

func TransparentMaterial(c Color, index float64) Material {
	return Material{c, index, true, false}
}