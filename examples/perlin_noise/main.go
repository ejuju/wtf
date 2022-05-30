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

func main() {
	start := time.Now()

	f, err := os.Open("test.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	baseimg, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	outputImages := []image.Image{}

	numFrames := 5
	for i := 0; i < numFrames; i++ {
		noiseGenerator := random.NewAquilaxPerlinNoiseGenerator(random.AquilaxPerlinNoiseGeneratorConfig{
			Alpha: 1,
			Beta:  1,
			N:     int32(i),
			Seed:  100,
		})

		perlinNoiseModifier := imgutil.NewPerlinNoisePixelModifier(imgutil.PerlinNoisePixelModifierConfig{
			PerlinNoiseGenerator: noiseGenerator,
		})

		outputImages = append(outputImages, imgutil.ModifyImage(baseimg, perlinNoiseModifier))
	}

	outputGif, err := gifutil.ImagesToGIF(500, outputImages...)
	if err != nil {
		panic(err)
	}

	err = gifutil.EncodeAndSaveToFile(outputGif, strconv.Itoa(int(time.Now().Unix()))+".gif")
	if err != nil {
		panic(err)
	}

	fmt.Printf("done in %vms\n", time.Now().Sub(start).Milliseconds())
}
