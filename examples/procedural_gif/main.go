package main

import (
	"fmt"
	"image"
	"strconv"
	"time"

	"github.com/ejuju/wtf/internal/colors"
	"github.com/ejuju/wtf/internal/gifutil"
	"github.com/ejuju/wtf/internal/imgutil"
	"github.com/ejuju/wtf/internal/random"
)

type Result struct {
	number int
	img    image.Image
}

func main() {
	width := 1024
	height := width
	baseimg := imgutil.NewMonochromeImage(colors.WhiteRGBA, width, height)

	// use gif generator
	result := gifutil.NewPerlinNoiseGIFMaker(gifutil.PerlinNoiseGIFMakerConfig{
		OutOfFrameFallbackColor: colors.BlackRGBA,
		MaxAmplitude:            float64(width * 2),
		Generator: random.NewAquilaxPerlinNoiseGenerator(random.AquilaxPerlinNoiseGeneratorConfig{
			Alpha: 2,
			Beta:  2,
			N:     10,
			Seed:  100,
		}),
	}).Generate(baseimg, 10)

	// log performance report
	fmt.Printf("Performance report\n%+v\n", result.PerformanceReport)

	// save GIF to file on local disk
	err := gifutil.EncodeAndSaveToFile(result.GIF, strconv.Itoa(int(time.Now().Unix()))+".gif")
	if err != nil {
		panic(err)
	}
}
