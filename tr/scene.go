package tr

// import "fmt"

type Scene struct {
	objects []Shape		// TODO(tizian): Replace with abstract Acceleration Structure (Possible implementations: List, Grid, Octree, K-d Tree, BSP Tree)
	lights []PointLight		// TODO(tizian): Eventually replace with abstract Light interface
}

func (s *Scene) AddObject(obj Shape) {
	s.objects = append(s.objects, obj)	// TODO(tizian): If obj is triangle mesh: add individual triangles
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

		if !lightOccluded {
			lights = append(lights, light)
		}
	}

	return lights
}