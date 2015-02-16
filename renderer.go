package main

import (
	"image"
	"image/color"
	"sync"
)

type Renderer struct {
	width, height int
	numberCPUs int
	camera Camera
	pixelSampler PixelSampler
}

func (r *Renderer) Render(rt *Raytracer, scene *Scene) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, r.width, r.height))

	var wg sync.WaitGroup
	wg.Add(r.numberCPUs)

	for cpu := 0; cpu < r.numberCPUs; cpu++ {
		go func(cpu int) {

			defer wg.Done()
			for y := cpu; y < r.height; y += r.numberCPUs {
				for x := 0; x < r.width; x++ {

					samples := r.pixelSampler.SamplePoints()

					c := Color{0, 0, 0}

					for _, offset := range samples {
						ray := r.camera.CastRay(x, y, r.width, r.height, offset.x, offset.y)
						c = c.Add(rt.Trace(&ray, scene))
					}

					numSamples := r.pixelSampler.SamplesPerPixel() * r.pixelSampler.SamplesPerPixel()
					c = c.Div(float64(numSamples))

					r, g, b := c.ToInts()

					img.Set(x, y, color.RGBA{r, g, b, 0xff})
				}
			}

		}(cpu)
	}

	wg.Wait()

	return img
}