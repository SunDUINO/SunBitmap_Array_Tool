package main

import (
	"image"
	"image/color"
)

func ToMonochrome(img image.Image, threshold uint8) image.Image {
	b := img.Bounds()
	out := image.NewGray(b)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {

			r, g, b2, _ := img.At(x, y).RGBA()
			gray := uint8((r + g + b2) / 3 >> 8)

			if gray > threshold {
				out.Set(x, y, color.White)
			} else {
				out.Set(x, y, color.Black)
			}
		}
	}

	return out
}

func ProcessImage(img image.Image, thr uint8) (image.Image, string) {
	p := ToMonochrome(img, thr)
	code := ImageToCArray(p)
	return p, code
}
