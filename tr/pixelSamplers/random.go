package pixelSamplers

import (
	"fmt"
	"math/rand"
	. "github.com/tizian/tracer/tr"
)

type RandomPS struct {
	samplesPerPixel int
}

func CreateRandomPS(spp int) RandomPS {
	return RandomPS{spp}
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