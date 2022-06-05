package main

import (
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"time"

	"github.com/ejuju/wtf/pkg/imgutil"
)

func main() {
	f, err := os.Open("./test.jpg")
	if err != nil {
		panic(err)
	}
	baseimg, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	result := imgutil.ModifyImage(
		baseimg,
		imgutil.NewBlurPixelModifier(imgutil.BlurPixelModifierConfig{MaxAmplitude: 40}),
		imgutil.ImageModificationOptions{Padding: 0.2},
	)

	f, err = os.Create(strconv.Itoa(int(time.Now().Unix())) + ".png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, result)
	if err != nil {
		panic(err)
	}
}
