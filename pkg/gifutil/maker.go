package gifutil

import (
	"image"
	"image/gif"
	"time"
)

type Maker interface {
	Make(image.Image, int) MakeFuncResult
}

type MakeFuncResult struct {
	GIF               *gif.GIF
	PerformanceReport MakePerformanceReport
}

type MakePerformanceReport struct {
	PixelsPerMillisecond    int
	FrameGenerationDuration time.Duration
	FrameEncodingDuration   time.Duration
}
