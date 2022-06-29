package imgutil

import (
	"image/color"
	"reflect"
	"testing"
)

func TestNewMonochromeImage(t *testing.T) {
	width := 2
	height := width
	clr := color.RGBA{R: 123}
	img := NewMonochromeImage(clr, width, height)
	bounds := img.Bounds()

	t.Run("should have the right dimensions", func(t *testing.T) {
		if bounds.Dx() != width {
			t.Fatalf("unexpected width, want %v, but got %v", width, bounds.Dx())
		}
		if bounds.Dy() != height {
			t.Fatalf("unexpected height, want %v, but got %v", height, bounds.Dy())
		}
	})

	t.Run("should have the right color on each pixel", func(t *testing.T) {
		for x := 0; x < bounds.Dx(); x++ {
			for y := 0; y < bounds.Dy(); y++ {
				pxColor, ok := img.At(x, y).(color.RGBA)
				if !ok {
					t.Fatal("failed to assert color to RGBA")
				}

				if !reflect.DeepEqual(clr, pxColor) {
					t.Fatalf("unexpected color at pixel [%d;%d], want %#v, but got %#v", x, y, clr, pxColor)
				}
			}
		}
	})
}
