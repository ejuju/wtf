package gifutil

import (
	"image"
	"image/color"
	"math"
	"time"

	"github.com/ejuju/wtf/pkg/imgutil"
	"github.com/ejuju/wtf/pkg/random"
)

type PerlinNoiseGIFMaker struct {
	config PerlinNoiseGIFMakerConfig
}

type PerlinNoiseGIFMakerConfig struct {
	FrameDelay               int
	OutOfFrameFallbackColor  color.RGBA
	MaxAmplitude             float64
	Generator                random.PerlinNoiseGenerator
	ImageModificationOptions imgutil.ImageModificationOptions
	PositionGapDivider       float64
}

func NewPerlinNoiseGIFMaker(config PerlinNoiseGIFMakerConfig) *PerlinNoiseGIFMaker {
	return &PerlinNoiseGIFMaker{config: config}
}

func (pngm *PerlinNoiseGIFMaker) Generate(img image.Image, numFrames int) MakeFuncResult {
	type Result struct {
		number int
		img    image.Image
	}

	// init variables for generating frames
	startFrameGeneration := time.Now()
	outputImages := make([]image.Image, numFrames)
	resultChan := make(chan Result, numFrames)
	defer close(resultChan)

	// generate each image frame concurrently
	for i := 0; i < numFrames; i++ {
		go func(i int) {
			step := float64(i) / float64(numFrames)
			perlinNoiseModifier := imgutil.NewPerlinNoisePixelModifier(imgutil.PerlinNoisePixelModifierConfig{
				OutOfFrameFallbackColor: pngm.config.OutOfFrameFallbackColor,
				Amplitude:               pngm.config.MaxAmplitude * (math.Cos(2*math.Pi*step)/2 + 0.5),
				PerlinNoiseGenerator:    pngm.config.Generator,
				PositionGapDivider:      pngm.config.PositionGapDivider,
			})
			resultChan <- Result{number: i, img: imgutil.ModifyImage(img, perlinNoiseModifier, pngm.config.ImageModificationOptions)}
		}(i)
	}

	// collect results from channel
	results := map[int]image.Image{}
	for i := 0; i < numFrames; i++ {
		result := <-resultChan
		results[result.number] = result.img
	}

	// order images from results chan
	for i, result := range results {
		outputImages[i] = result
	}

	// store timestamp for end of frame generation
	endFrameGeneration := time.Now()

	// encode images to GIF
	outputGif, err := ImagesToGIF(pngm.config.FrameDelay, outputImages...)
	if err != nil {
		panic(err)
	}

	// log duration to concatenate GIF frames
	endFrameEncoding := time.Now()

	// init make result
	bounds := img.Bounds()
	numPixels := bounds.Dx() * bounds.Dy()
	totalPixels := numPixels * numFrames
	makeFuncResult := MakeFuncResult{
		GIF: outputGif,
		PerformanceReport: MakePerformanceReport{
			FrameGenerationDuration: endFrameGeneration.Sub(startFrameGeneration),
			FrameEncodingDuration:   endFrameEncoding.Sub(endFrameGeneration),
			PixelsPerMillisecond:    totalPixels / int(endFrameEncoding.Sub(startFrameGeneration).Milliseconds()),
		},
	}

	return makeFuncResult
}
