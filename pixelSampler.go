package main

import (
	"fmt"
	"math/rand"
)

type PixelSampler interface {
	SamplePoints() []Vector2
	SamplesPerPixel() int
}

type RandomPS struct {
	samplesPerPixel int
}

func (ps RandomPS) String() string {
	return fmt.Sprintf("Random")
}

func (ps RandomPS) SamplePoints() []Vector2 {
	n := ps.samplesPerPixel
	samples := make([]Vector2, n*n, n*n)

	for i := 0; i < n*n; i++ {
		// TODO(tizian): rand is deterministic without setting a seed
		u := rand.Float64()
		v := rand.Float64()
		samples[i] = Vector2{u, v}
	}

	return samples
}

func (ps RandomPS) SamplesPerPixel() int {
	return ps.samplesPerPixel
}

type GridPS struct {
	samplesPerPixel int
}

func (ps GridPS) String() string {
	return fmt.Sprintf("Grid")
}

func (ps GridPS) SamplePoints() []Vector2 {
	n := ps.samplesPerPixel
	samples := make([]Vector2, n*n, n*n)

	dx := 1/float64(n)
	ox := dx/2

	index := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			samples[index] = Vector2{ox + float64(i) * dx, ox + float64(j) * dx}
			index++
		}
	}

	return samples
}

func (ps GridPS) SamplesPerPixel() int {
	return ps.samplesPerPixel
}