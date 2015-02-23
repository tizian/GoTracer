package tr

type Material struct {
	color Color
	index float64
	absorption float64
	transparent bool
	specular bool
}

func DiffuseMaterial(c Color) Material {
	return Material{c, -1, -1, false, false}
}

func SpecularMaterial(c Color, index, absorption float64) Material {
	return Material{c, index, absorption, false, true}
}

func TransparentMaterial(c Color, index float64) Material {
	return Material{c, index, -1, true, false}
}

func (m *Material) Color() Color {
	return m.color
}