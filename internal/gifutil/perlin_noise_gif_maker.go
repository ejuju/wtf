package gifutil

import (
	"image"
	"image/color"
	"math"
	"time"

	"github.com/ejuju/wtf/internal/imgutil"
	"github.com/ejuju/wtf/internal/random"
)

type PerlinNoiseGIFMaker struct {
	config PerlinNoiseGIFMakerConfig
}

type PerlinNoiseGIFMakerConfig struct {
	OutOfFrameFallbackColor color.RGBA
	MaxAmplitude            float64
	Generator               random.PerlinNoiseGenerator
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
			perlinNoiseModifier := imgutil.NewPerlinNoisePixelModifier(imgutil.PerlinNoisePixelModifierConfig{
				OutOfFrameFallbackColor: pngm.config.OutOfFrameFallbackColor,
				Amplitude:               math.Abs(-pngm.config.MaxAmplitude + (2*pngm.config.MaxAmplitude)*(float64(i)/float64(numFrames))),
				PerlinNoiseGenerator:    pngm.config.Generator,
			})
			resultChan <- Result{number: i, img: imgutil.ModifyImage(img, perlinNoiseModifier)}
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
	outputGif, err := ImagesToGIF(1, outputImages...)
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
			NumPixels:               numPixels,
			NumFrames:               numFrames,
			TotalPixels:             totalPixels,
			FrameGenerationDuration: endFrameGeneration.Sub(startFrameGeneration),
			FrameEncodingDuration:   endFrameEncoding.Sub(endFrameGeneration),
			PixelsPerMillisecond:    totalPixels / int(endFrameEncoding.Sub(startFrameGeneration).Milliseconds()),
		},
	}

	return makeFuncResult
}
