package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
)

func ImageToCArray(img image.Image) string {
	b := img.Bounds()
	w, h := b.Dx(), b.Dy()

	var buf bytes.Buffer
	buf.WriteString("const uint8_t img[] = {\n")

	for y := 0; y < h; y++ {
		var byteVal uint8
		bit := 0

		for x := 0; x < w; x++ {
			isWhite := img.At(x, y) == color.White
			if isWhite {
				byteVal |= (1 << (7 - bit))
			}
			bit++

			if bit == 8 {
				buf.WriteString(fmt.Sprintf("0x%02X, ", byteVal))
				bit = 0
				byteVal = 0
			}
		}
		buf.WriteString("\n")
	}

	buf.WriteString("};\n")
	return buf.String()
}
