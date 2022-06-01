package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ejuju/wtf/pkg/colors"
	"github.com/ejuju/wtf/pkg/gifutil"
	"github.com/ejuju/wtf/pkg/imgutil"
	"github.com/ejuju/wtf/pkg/random"
)

func main() {
	start := time.Now()

	// from monochrome image
	width := 1024
	height := width
	baseimg := imgutil.NewMonochromeImage(colors.White, width, height)

	// // from local image
	// f, err := os.Open("./test.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// baseimg, err := jpeg.Decode(f)
	// if err != nil {
	// 	panic(err)
	// }

	// calculate stats
	numFrames := 100
	// imgbounds := baseimg.Bounds()
	// width := imgbounds.Dx()
	// height := imgbounds.Dy()
	numPixelsPerFrame := width * height
	totalPixels := numFrames * numPixelsPerFrame
	estimatedPixelsPerMillisecond := 1000
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
		ImageModificationOptions: imgutil.ImageModificationOptions{Padding: 0.069},
		FrameDelay:               0,
		MaxAmplitude:             1024,
		PositionGapDivider:       20.67,
		OutOfFrameFallbackColor:  colors.Black,
		Generator:                random.NewOpenSimplexNoiseGenerator(random.OpenSimplexNoiseGeneratorConfig{Seed: 0}),
		// Generator: random.NewAquilaxPerlinNoiseGenerator(random.AquilaxPerlinNoiseGeneratorConfig{
		// 	Alpha: 2,
		// 	Beta:  2,
		// 	N:     4,
		// 	Seed:  0,
		// }),
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
	err := gifutil.EncodeAndSaveToFile(result.GIF, strconv.Itoa(int(time.Now().Unix()))+".gif")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total duration: %s\n\n", time.Now().Sub(start))
}
