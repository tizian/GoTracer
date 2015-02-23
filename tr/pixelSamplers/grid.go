package pixelSamplers

import (
	"fmt"
	. "github.com/tizian/tracer/tr"
)

type GridPS struct {
	samplesPerPixel int
}

func CreateGridPS(spp int) GridPS {
	return GridPS{spp}
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