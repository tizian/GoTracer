package main

import (
	"math"
	// "fmt"
)

type Raytracer struct {
	recursionDepth int
}

func (rt *Raytracer) Trace(ray *Ray, scene *Scene) Color {
	return RecursiveTrace(ray, scene, rt.recursionDepth).Clamp()
}

// Recursive ray tracing function
func RecursiveTrace(ray *Ray, scene *Scene, depth int) Color {

	// Intersect the ray to scene elements and determine the closest hit
	hit, ok := scene.Intersect(ray)

	if !ok {
		return Color{} // Ray doesn't hit any objects -> Return black background color
	}

	ambientColor := GlobalAmbientTerm(&hit)
	diffuseAndSpecularColor := Color{}

	visibleLights := scene.VisibleLights(hit.point)

	for _, light := range visibleLights {
		phong := PhongIllumination(&hit, &light.position, &ray.origin)
		diffuseAndSpecularColor = diffuseAndSpecularColor.Add(phong)	// TODO(tizian): light colors
	}

	objMaterial := hit.object.Material(hit.point)

	if depth <= 0 || (!objMaterial.specular && !objMaterial.transparent) {
		return ambientColor.Add(diffuseAndSpecularColor)
	}

	if objMaterial.specular {
		p := hit.point
		n := hit.object.Normal(p)

		R, T := SchlickApproximation(1.0, objMaterial.index, n, ray.direction)

		reflectedRay := ray.Reflect(p, n)
		reflectedColor := RecursiveTrace(&reflectedRay, scene, depth-1)

		c1 := ambientColor.Add(diffuseAndSpecularColor).Mul(T)
		c2 := reflectedColor.Mul(R)

		return c1.Add(c2)
	}

	if objMaterial.transparent {
		p := hit.point
		d := ray.direction
		n := hit.object.Normal(p)

		var ni, nt float64
		var reflectedRay, refractedRay Ray
		k := objMaterial.color

		if d.Dot(n) < 0 {
			ni = 1.0
			nt = objMaterial.index
		} else {
			nt = objMaterial.index
			ni = 1.0
			n = n.Mul(-1)
			exp := math.Exp(-0.05*hit.t)	// TODO(tizian): Beer's law coefficients?
			k = Color{exp, exp, exp}
		}

		reflectedRay = ray.Reflect(p, n)
		refractedRay, _ = ray.Refract(p, n, ni, nt)

		reflectedColor := RecursiveTrace(&reflectedRay, scene, depth-1)
		refractedColor := RecursiveTrace(&refractedRay, scene, depth-1)

		R, T := SchlickApproximation(ni, nt, n, d)

		return k.MulC(reflectedColor.Mul(R).Add(refractedColor.Mul(T)))
	}


	return ambientColor.Add(diffuseAndSpecularColor)
}

func SchlickApproximation(ni, nt float64, n, i Vector3) (float64, float64) {
	R0 := (nt - ni) / (nt + ni)
	R0 = R0 * R0
	cosX := -n.Dot(i)

	if ni > nt {
		inv_eta := ni / nt
		sinT2 := inv_eta * inv_eta * (1 - cosX * cosX)
		if sinT2 > 1 {
			return 1, 0	// TIR
		}
		cosX = math.Sqrt(1 - sinT2)
	}

	R := R0 + (1 - R0) * math.Pow(1 - cosX, 5)
	T := 1 - R

	return R, T
}

/*
// Recursive ray tracing function
func RecursiveTrace(ray *Ray, scene *Scene, depth int) Color {

	// Intersect the ray to scene elements and determine the closest one
	hit, ok := scene.Intersect(ray)

	if !ok {
		return Color{}	// Ray doesn't hit any object -> Return black background color
	}

	objectMaterial := hit.object.Material(hit.point)

	color := Color{}

	ambient := GlobalAmbientTerm(&hit)
	color = color.Add(ambient)
	visibleLights := scene.VisibleLights(hit.point)

	// fmt.Println(len(visibleLights))

	for _, light := range visibleLights {
		phong := PhongIllumination(&hit, &light.position, &ray.origin)
		color = color.Add(phong)	// TODO(tizian): light colors
	}

	if objectMaterial.specular && depth > 0 {
		p := hit.point
		n := hit.object.Normal(p)

		reflectedRay := ray.Reflect(p, n)
		Ls := RecursiveTrace(&reflectedRay, scene, depth-1)
		color = Ls.MulC(hit.object.Color(p))

	} else if objectMaterial.transparent && depth > 0 {
		p := hit.point
		d := ray.direction
		n := hit.object.Normal(p)

		material := hit.object.Material(p)

		var eta1, eta2 float64
		var reflectedRay, refractedRay Ray

		if d.Dot(n) < 0 {
			eta1 = 1.0
			eta2 = material.index
		} else {
			eta1 = material.index
			eta2 = 1.0
			n = n.Mul(-1)
		}

		reflectedRay = ray.Reflect(p, n)
		refractedRay, ok = ray.Refract(p, n, eta1, eta2)

		if !ok {
			return RecursiveTrace(&reflectedRay, scene, depth-1)
		}

		reflectedColor := RecursiveTrace(&reflectedRay, scene, depth-1)
		refractedColor := RecursiveTrace(&refractedRay, scene, depth-1)

		return reflectedColor.Mul(0.3).Add(refractedColor.Mul(0.7))

		// var cos float64
		if d.Dot(n) < 0 {	// from outside of normal
			reflectedRay = ray.Reflect(p, n)
			refractedRay, _ = ray.Refract(p, n, 1.0, material.index)
			// cos = d.Dot(n)
		} else {	// from inside of normal
			n = n.Mul(-1)
			reflectedRay = ray.Reflect(p, n)
			refractedRay, ok = ray.Refract(p, n, material.index, 1.0)
			if ok {
				// cos = refractedRay.direction.Dot(n)
			} else {
				// return RecursiveTrace(&reflectedRay, scene, depth-1)
			}
		}
		// R0 := (material.index-1)/(material.index+1)
		// R0 = R0*R0
		// R := R0 + (1-R0) * math.Pow(1-cos, 5)

		reflectedColor := RecursiveTrace(&reflectedRay, scene, depth-1)
		refractedColor := RecursiveTrace(&refractedRay, scene, depth-1)

		// return reflectedColor.Mul(R).Add(refractedColor.Mul(1-R))
		return refractedColor
		return reflectedColor.Mul(0.5).Add(refractedColor.Mul(0.5))


	}


	return color
}
*/

func GlobalAmbientTerm(hit *Hit) Color {
	return hit.object.Color(hit.point).Mul(0.4)	// TODO(tizian): Create GlobalAmbientIntensity variable
}

func PhongIllumination(hit *Hit, pLight, pEye *Vector3) Color {
	V := pEye.Sub(hit.point).Normalize()
	L := pLight.Sub(hit.point).Normalize()
	N := hit.normal

	LN := L.Dot(N)
	R := N.Mul(2*LN).Sub(L).Normalize()
	RV := R.Dot(V)

	color := Color{}

	objectColor := hit.object.Color(hit.point)

	// Diffuse
	if LN > 0 {
		diffuse := objectColor.Mul(LN)	// TODO(tizian): Light intensity
		color = color.Add(diffuse)
	}

	// Specular
	if RV > 0 {
		specular := objectColor.Mul(math.Pow(RV, 25))	// TODO(tizian): Light intensity & specular exponent
		color = color.Add(specular)
	}

	return color
}