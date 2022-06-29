package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"strconv"
	"time"

	"github.com/ejuju/wtf/pkg/colors"
	"github.com/ejuju/wtf/pkg/gifutil"
	"github.com/ejuju/wtf/pkg/imgutil"
	"github.com/ejuju/wtf/pkg/noise"
)

func main() {
	start := time.Now()
	var err error

	// from monochrome image
	// baseimg := imgutil.NewMonochromeImage(colors.White, 1024, 1024)

	// from local image
	f, err := os.Open("./input/alexandre-brondino.jpg")
	if err != nil {
		panic(err)
	}
	baseimg, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	// calculate stats
	numFrames := 20
	imgbounds := baseimg.Bounds()
	width := imgbounds.Dx()
	height := imgbounds.Dy()
	numPixelsPerFrame := width * height
	totalPixels := numFrames * numPixelsPerFrame
	estimatedPixelsPerMillisecond := 1500
	estimatedTotalDuration := time.Duration(totalPixels/estimatedPixelsPerMillisecond) * 1000 * 1000

	fmt.Printf("\n"+
		"Input image width: %v\n"+
		"Input image height: %v\n"+
		"Number of pixels per frame: %v\n"+
		"Number of frames: %v\n"+
		"Total number of pixels: %v\n"+
		"Estimated pixels per millisecond: %v\n"+
		"Estimated total duration: %s\n",
		width,
		height,
		numPixelsPerFrame,
		numFrames,
		totalPixels,
		estimatedPixelsPerMillisecond,
		estimatedTotalDuration,
	)

	// use gif generator
	result := gifutil.NewPerlinNoiseGIFMaker(gifutil.PerlinNoiseGIFMakerConfig{
		ImageModificationOptions: imgutil.ImageModificationOptions{Padding: 0.0},
		MaxAmplitude:             1000,
		PositionGapDivider:       1000.65,
		OutOfFrameFallbackColor:  colors.Black,
		// Generator: noise.NewAquilaxGenerator(noise.AquilaxGeneratorConfig{Alpha: 2, Beta: 2, N: 4, Seed: 0}),
		Generator: noise.NewOjracGenerator(0),
	}).Generate(baseimg, numFrames)

	// log performance report
	fmt.Printf("\n"+
		"Pixels per millisecond: %v\n"+
		"Frame encoding duration: %v\n"+
		"Frame encoding duration: %v\n",
		result.PerformanceReport.PixelsPerMillisecond,
		result.PerformanceReport.FrameGenerationDuration,
		result.PerformanceReport.FrameEncodingDuration,
	)

	// save GIF to file on local disk
	err = gifutil.EncodeAndSaveToFile(result.GIF, "./output/"+strconv.Itoa(int(time.Now().Unix()))+".gif")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total duration: %s\n\n", time.Now().Sub(start))
}
