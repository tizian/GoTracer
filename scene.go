package main

// import "fmt"

type Scene struct {
	objects []Shape
	lights []PointLight
}

func (s *Scene) AddObject(obj Shape) {
	s.objects = append(s.objects, obj)
}

func (s *Scene) AddLight(light PointLight) {
	s.lights = append(s.lights, light)
}

func (s *Scene) Intersect(r *Ray) (Hit, bool) {
	hit := Hit{}
	nearest := INF
	for _, obj := range s.objects {
		t := obj.Intersect(r)
		if t < nearest && t > EPS {
			nearest = t
			p := r.Position(t)
			n := obj.Normal(p)
			hit = Hit{obj, p, n, t}
		}
	}
	ok := nearest < INF
	return hit, ok
}

func (s *Scene) IntersectAny(r *Ray, maxDist float64) bool {
	for _, obj := range s.objects {
		t := obj.Intersect(r)
		if t < maxDist && t > EPS {
			// fmt.Println("Intersection!")
			return true
		}
	}

	return false
}

func (s *Scene) VisibleLights(p Vector3) []PointLight {
	lights := []PointLight{}

	for _, light := range s.lights {
		shadowRay := Ray{p, light.position.Sub(p).Normalize()}

		lightOccluded := s.IntersectAny(&shadowRay, light.position.Sub(p).Length())

		// fmt.Println(lightOccluded)

		if !lightOccluded {
			lights = append(lights, light)
		}
	}

	if !(len(lights) == 1) {
		// fmt.Println("hmm")
	}

	return lights
}