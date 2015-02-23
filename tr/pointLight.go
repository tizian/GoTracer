package tr

type PointLight struct {
	position Vector3
}

func CreatePointLight(position Vector3) PointLight {
	return PointLight{position}
}