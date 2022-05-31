package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"strconv"
	"time"

	"github.com/ejuju/wtf/internal/gifutil"
	"github.com/ejuju/wtf/internal/random"
)

func main() {
	// open input image file
	f, err := os.Open("test.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// decode input image
	baseimg, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	// use gif generator
	result := gifutil.NewPerlinNoiseGIFMaker(gifutil.PerlinNoiseGIFMakerConfig{
		MaxAmplitude: 1000,
		Generator: random.NewAquilaxPerlinNoiseGenerator(random.AquilaxPerlinNoiseGeneratorConfig{
			Alpha: 2,
			Beta:  2,
			N:     2,
			Seed:  100,
		}),
	}).Generate(baseimg, 50)

	// log performance report
	fmt.Printf("performance report:\n%+v\n", result.PerformanceReport)

	// save GIF to file on local disk
	err = gifutil.EncodeAndSaveToFile(result.GIF, strconv.Itoa(int(time.Now().Unix()))+".gif")
	if err != nil {
		panic(err)
	}
}
