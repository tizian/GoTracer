package main

import (
	"fmt"
	"image/png"
	"os"
	"log"
	"runtime"
	"github.com/tizian/tracer/tr"
	"github.com/tizian/tracer/tr/shapes"
	"github.com/tizian/tracer/tr/pixelSamplers"
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

	gold := tr.SpecularMaterial(tr.ColorFromInts(185, 148, 69), 2, 3)
	// glass := TransparentMaterial(Color{1, 1, 1}, 1.5)
	blueGlass := tr.TransparentMaterial(tr.Color{0.87, 0.87, 1}, 1.5)
	redGlass := tr.TransparentMaterial(tr.Color{1, 0.87, 0.87}, 1.5)
	greenGlass := tr.TransparentMaterial(tr.Color{0.87, 1, 0.87}, 1.5)
	yellowGlass := tr.TransparentMaterial(tr.Color{1, 1, 0.87}, 1.5)
	purpleGlass := tr.TransparentMaterial(tr.Color{1, 0.87, 1}, 1.5)

	white := tr.DiffuseMaterial(tr.Color{0.740, 0.742, 0.734})
	red := tr.DiffuseMaterial(tr.Color{0.366, 0.037, 0.042})
	green := tr.DiffuseMaterial(tr.Color{0.163, 0.409, 0.083})
	// yellow := DiffuseMaterial(Color{0.5, 0.5, 0})

	// triangle := Triangle{}
	// triangle.v1 = Vector3{-10, -10, -9.99}
	// triangle.v2 = Vector3{0, 10, -9.99}
	// triangle.v3 = Vector3{10, -10, -9.99}
	// triangle.n1 = Vector3{0, 0, 1}
	// triangle.material = yellow

	scene := tr.Scene{}
	scene.AddObject(shapes.CreatePlane(tr.Vector3{0, -10, 0}, tr.Vector3{0, 1, 0}, white))
	scene.AddObject(shapes.CreatePlane(tr.Vector3{0, 10, 0}, tr.Vector3{0, -1, 0}, white))
	scene.AddObject(shapes.CreatePlane(tr.Vector3{0, 0, 10}, tr.Vector3{0, 0, -1}, white))
	scene.AddObject(shapes.CreatePlane(tr.Vector3{0, 0, -10}, tr.Vector3{0, 0, 1}, white))
	scene.AddObject(shapes.CreatePlane(tr.Vector3{-10, 0, 0}, tr.Vector3{1, 0, 0}, green))
	scene.AddObject(shapes.CreatePlane(tr.Vector3{10, 0, 0}, tr.Vector3{-1, 0, 0}, red))

	// scene.AddObject(&triangle)

	/*
	scene.AddObject(CreateSphere(Vector3{0, -6, 0}, 4, glass))
	scene.AddObject(CreateSphere(Vector3{0, 0, 0}, 2, glass))
	scene.AddObject(CreateSphere(Vector3{0, 3, 0}, 1, glass))
	scene.AddObject(CreateSphere(Vector3{0, 4.5, 0}, 0.5, glass))
	scene.AddObject(CreateSphere(Vector3{0, 5.25, 0}, 0.25, glass))
	*/

	scene.AddObject(shapes.CreateSphere(tr.Vector3{-4, -8, 0}, 2, yellowGlass))
	scene.AddObject(shapes.CreateSphere(tr.Vector3{-4, -4, 0}, 2, greenGlass))
	scene.AddObject(shapes.CreateSphere(tr.Vector3{-4, 0, 0}, 2, blueGlass))
	scene.AddObject(shapes.CreateSphere(tr.Vector3{-4, 4, 0}, 2, purpleGlass))
	scene.AddObject(shapes.CreateSphere(tr.Vector3{-4, 8, 0}, 2, redGlass))

	scene.AddObject(shapes.CreateSphere(tr.Vector3{4, -8, 0}, 2, gold))
	scene.AddObject(shapes.CreateSphere(tr.Vector3{4, -4, 0}, 2, gold))
	scene.AddObject(shapes.CreateSphere(tr.Vector3{4, 0, 0}, 2, gold))
	scene.AddObject(shapes.CreateSphere(tr.Vector3{4, 4, 0}, 2, gold))
	scene.AddObject(shapes.CreateSphere(tr.Vector3{4, 8, 0}, 2, gold))

	scene.AddLight(tr.CreatePointLight(tr.Vector3{0, 6, 0}))

	// Camera
	// camera := LookAt(Vector3{0, 0, -20}, Vector3{0, 0, 1}, Vector3{0, 1, 0}, 65)
	camera := tr.LookAt(tr.Vector3{0, 0, 20}, tr.Vector3{0, 0, -1}, tr.Vector3{0, 1, 0}, 65)

	// PixelSampler
	ps := pixelSamplers.CreateGridPS(4)

	// Raytracer
	rt := tr.Raytracer{}
	rt.RecursionDepth = 5

	// Renderer
	renderer := tr.CreateRenderer(512, 512, runtime.GOMAXPROCS(0), &camera, ps)

	fmt.Printf("Tracer\n\n")

	fmt.Printf("Scene: Test Scene\n")
	fmt.Printf("%s", renderer)

	img := renderer.Render(&rt, &scene)

	file, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	png.Encode(file, img)
}