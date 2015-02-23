package tr

type PixelSampler interface {
	SamplePoints() []Vector2
	SamplesPerPixel() int
}