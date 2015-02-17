package main

import (
	"math"
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

		R, T := FresnelConductor(objMaterial.index, objMaterial.absorption, n, ray.direction)

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

		var etai, etat float64
		var reflectedRay, refractedRay Ray
		k := objMaterial.color

		if d.Dot(n) < 0 {
			etai = 1.0
			etat = objMaterial.index
		} else {
			etat = objMaterial.index
			etai = 1.0
			n = n.Mul(-1)
			exp := math.Exp(-0.05 * hit.t)	// TODO(tizian): Beer's law coefficients?
			k = Color{exp, exp, exp}
		}

		reflectedRay = ray.Reflect(p, n)
		refractedRay, _ = ray.Refract(p, n, etai, etat)

		reflectedColor := RecursiveTrace(&reflectedRay, scene, depth-1)
		refractedColor := RecursiveTrace(&refractedRay, scene, depth-1)

		R, T := FresnelDielectric(etai, etat, n, d)

		return k.MulC(reflectedColor.Mul(R).Add(refractedColor.Mul(T)))
	}


	return ambientColor.Add(diffuseAndSpecularColor)
}

func FresnelDielectric(etai, etat float64, n, i Vector3) (R, T float64) {
	eta := etai / etat
	cosi := -n.Dot(i)
	sint2 := eta * eta * (1 - cosi * cosi)
	if sint2 > 1 {
		return 1, 0 	// TIR
	}

	cost := math.Sqrt(1 - sint2)
	Rparl := (etat * cosi - etai * cost) / (etat * cosi + etai * cost)
	Rperp := (etai * cosi - etat * cost) / (etai * cosi + etat * cost)

	R = (Rparl * Rparl + Rperp * Rperp) / 2
	T = 1 - R

	return
}

func FresnelConductor(eta, k float64, n, i Vector3) (R, T float64) {
	cosi := -n.Dot(i)
	tmp := (eta * eta + k * k)
	Rparl2 := (tmp - (2 * eta * cosi) + 1) / (tmp + (2 * eta * cosi) + 1)
	tmp_f := eta * eta + k * k
	Rperp2 := (tmp_f - (2 * eta * cosi) + cosi * cosi) / (tmp_f + (2 * eta * cosi) + cosi * cosi)

	R = (Rparl2 + Rperp2) / 2
	T = 1 - R

	return
}

func FresnelNoOp() (R, T float64) {
	R = 1
	T = 0

	return
}

func SchlickApproximation(etai, etat float64, n, i Vector3) (R, T float64) {
	R0 := (etat - etai) / (etat + etai)
	R0 = R0 * R0
	cosi := -n.Dot(i)

	if etai > etat {
		eta := etai / etat
		sint2 := eta * eta * (1 - cosi * cosi)
		if sint2 > 1 {
			return 1, 0	// TIR
		}
		cosi = math.Sqrt(1 - sint2)
	}

	R = R0 + (1 - R0) * math.Pow(1 - cosi, 5)
	T = 1 - R

	return
}

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