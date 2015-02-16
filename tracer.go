package main

import (
	"fmt"
	"image/png"
	"os"
	"log"
	"runtime"
)

var width int = 640
var height int = 480

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	/*
	blue := DiffuseMaterial(Color{0.123, 0.24, 0.87})
	// mirror := SpecularMaterial(Color{1, 1, 1})
	// gold := SpecularMaterial(ColorFromInts(185, 148, 69))
	glass := TransparentMaterial(Color{1, 1, 1}, 1.5)

	// Cornell box scene
	scene := Scene{}

	scene.AddObject(CreatePlane(Vector3{0, -10, 0}, Vector3{0, 1, 0}, white))
	//scene.AddObject(CreatePlane(Vector3{0, 10, 0}, Vector3{0, -1, 0}, white))
	//scene.AddObject(CreatePlane(Vector3{0, 0, 10}, Vector3{0, 0, -1}, white))
	//scene.AddObject(CreatePlane(Vector3{0, 0, -10}, Vector3{0, 0, 1}, white))
	//scene.AddObject(CreatePlane(Vector3{-10, 0, 0}, Vector3{1, 0, 0}, green))
	//scene.AddObject(CreatePlane(Vector3{10, 0, 0}, Vector3{-1, 0, 0}, redWall))

	scene.AddObject(CreateSphere(Vector3{0, -6, 8}, 2.0, blue))
	scene.AddObject(CreateSphere(Vector3{0, -6, 0}, 4.0, glass))

	scene.AddLight(PointLight{Vector3{0, 8, 0}})
	*/

	// gold := SpecularMaterial(ColorFromInts(185, 148, 69), 5)
	// glass := TransparentMaterial(Color{1, 1, 1}, 1.5)
	blueGlass := TransparentMaterial(Color{0.87, 0.87, 1}, 1.5)
	redGlass := TransparentMaterial(Color{1, 0.87, 0.87}, 1.5)
	greenGlass := TransparentMaterial(Color{0.87, 1, 0.87}, 1.5)
	yellowGlass := TransparentMaterial(Color{1, 1, 0.87}, 1.5)
	purpleGlass := TransparentMaterial(Color{1, 0.87, 1}, 1.5)

	white := DiffuseMaterial(Color{0.740, 0.742, 0.734})
	red := DiffuseMaterial(Color{0.366, 0.037, 0.042})
	green := DiffuseMaterial(Color{0.163, 0.409, 0.083})

	scene := Scene{}
	scene.AddObject(CreatePlane(Vector3{0, -10, 0}, Vector3{0, 1, 0}, white))
	scene.AddObject(CreatePlane(Vector3{0, 10, 0}, Vector3{0, -1, 0}, white))
	scene.AddObject(CreatePlane(Vector3{0, 0, 10}, Vector3{0, 0, -1}, white))
	scene.AddObject(CreatePlane(Vector3{0, 0, -10}, Vector3{0, 0, 1}, white))
	scene.AddObject(CreatePlane(Vector3{-10, 0, 0}, Vector3{1, 0, 0}, green))
	scene.AddObject(CreatePlane(Vector3{10, 0, 0}, Vector3{-1, 0, 0}, red))

	/*
	scene.AddObject(CreateSphere(Vector3{0, -6, 0}, 4, glass))
	scene.AddObject(CreateSphere(Vector3{0, 0, 0}, 2, glass))
	scene.AddObject(CreateSphere(Vector3{0, 3, 0}, 1, glass))
	scene.AddObject(CreateSphere(Vector3{0, 4.5, 0}, 0.5, glass))
	scene.AddObject(CreateSphere(Vector3{0, 5.25, 0}, 0.25, glass))
	*/

	scene.AddObject(CreateSphere(Vector3{0, -8, 0}, 2, yellowGlass))
	scene.AddObject(CreateSphere(Vector3{0, -4, 0}, 2, greenGlass))
	scene.AddObject(CreateSphere(Vector3{0, 0, 0}, 2, blueGlass))
	scene.AddObject(CreateSphere(Vector3{0, 4, 0}, 2, purpleGlass))
	scene.AddObject(CreateSphere(Vector3{0, 8, 0}, 2, redGlass))

	/*
	scene.AddObject(CreateSphere(Vector3{0, -8, 0}, 2, gold))
	scene.AddObject(CreateSphere(Vector3{0, -4, 0}, 2, gold))
	scene.AddObject(CreateSphere(Vector3{0, 0, 0}, 2, gold))
	scene.AddObject(CreateSphere(Vector3{0, 4, 0}, 2, gold))
	scene.AddObject(CreateSphere(Vector3{0, 8, 0}, 2, gold))
	*/

	scene.AddLight(PointLight{Vector3{8, 8, 8}})

	// Camera
	// camera := LookAt(Vector3{0, 0, -20}, Vector3{0, 0, 1}, Vector3{0, 1, 0}, 65)
	camera := LookAt(Vector3{0, 0, 20}, Vector3{0, 0, -1}, Vector3{0, 1, 0}, 65)

	// PixelSampler
	ps := GridPS{}
	ps.samplesPerPixel = 4

	// Raytracer
	rt := Raytracer{}
	rt.recursionDepth = 5

	// Renderer
	renderer := Renderer{}
	renderer.width = 512
	renderer.height = 512
	renderer.numberCPUs = runtime.GOMAXPROCS(0)
	renderer.pixelSampler = ps
	renderer.camera = &camera

	fmt.Printf("Tracer\n\n")

	fmt.Printf("Scene: Test Scene\n")
	fmt.Printf("Camera: %s\n", renderer.camera)
	fmt.Printf("Resolution: %d x %d\n", renderer.width, renderer.height)
	fmt.Printf("Pixel sampler: %s\n", renderer.pixelSampler)
	fmt.Printf("%d x %d = %d samples per pixel\n", renderer.pixelSampler.SamplesPerPixel(), renderer.pixelSampler.SamplesPerPixel(), renderer.pixelSampler.SamplesPerPixel() * renderer.pixelSampler.SamplesPerPixel())
	fmt.Printf("Number of CPUs: %d\n", renderer.numberCPUs)

	img := renderer.Render(&rt, &scene)

	file, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	png.Encode(file, img)
}