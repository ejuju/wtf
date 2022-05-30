package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"time"

	"github.com/ejuju/wtf/internal/gifutil"
	"github.com/ejuju/wtf/internal/imgutil"
	"github.com/ejuju/wtf/internal/random"
)

type Result struct {
	number int
	img    image.Image
}

func main() {
	// start time for timer
	start := time.Now()

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

	// generate images concurrently
	numFrames := 50
	outputImages := make([]image.Image, numFrames)
	resultChan := make(chan Result, numFrames)
	defer close(resultChan)
	for i := 0; i < numFrames; i++ {
		go func(i int) {
			noiseGenerator := random.NewAquilaxPerlinNoiseGenerator(random.AquilaxPerlinNoiseGeneratorConfig{
				Alpha: 2,
				Beta:  2,
				N:     1,
				Seed:  100,
			})

			perlinNoiseModifier := imgutil.NewPerlinNoisePixelModifier(imgutil.PerlinNoisePixelModifierConfig{
				PerlinNoiseGenerator: noiseGenerator,
				Amplitude:            300 * (float64(i) / float64(numFrames)),
			})

			resultChan <- Result{number: i, img: imgutil.ModifyImage(baseimg, perlinNoiseModifier)}
		}(i)
	}

	results := map[int]image.Image{}
	for i := 0; i < numFrames; i++ {
		result := <-resultChan
		results[result.number] = result.img
	}

	// sort images from results
	for i, result := range results {
		outputImages[i] = result
	}

	// encode images to GIF
	outputGif, err := gifutil.ImagesToGIF(1, outputImages...)
	if err != nil {
		panic(err)
	}

	// save GIF to file on local disk
	err = gifutil.EncodeAndSaveToFile(outputGif, strconv.Itoa(int(time.Now().Unix()))+".gif")
	if err != nil {
		panic(err)
	}

	// print total execution time
	fmt.Printf("done in %vs\n", time.Now().Sub(start).Seconds())
}
