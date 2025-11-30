/*
================================================================================
File:        logic.go
Description: Logika przetwarzania obrazów, np. progowanie, dithering,
             oversampling, detekcja konturów oraz generowanie tablic C/Rust.
Author:      SunRiver / Lothar Team
Website:     https://forum.lothar-team.pl/
Version:     0.0.01
Date:        2025-11-30
================================================================================
*/

package main

import (
	"image"
	"image/color"
)

// --------- KONWERSJA NA MONOCHROM ---------

// ToMonochrome – konwertuje dowolny obraz na obraz czarno-biały
// threshold – próg jasności (0-255); powyżej -> biały, poniżej -> czarny
func ToMonochrome(img image.Image, threshold uint8) image.Image {
	b := img.Bounds()       // Pobranie granic obrazu (prostokąt z Min i Max)
	out := image.NewGray(b) // Tworzymy nowy obraz w skali szarości

	// Iteracja po wszystkich pikselach obrazu
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {

			// Pobranie wartości RGBA pikseli (każda składowa 16-bitowa: 0..65535)
			r, g, b2, _ := img.At(x, y).RGBA()

			// Konwersja na 8-bitową wartość szarości (średnia z RGB)
			// >>8 konwertuje 16-bitowy kanał (0..65535) na 8-bitowy (0..255)
			gray := uint8((r + g + b2) / 3 >> 8)

			// Porównanie z progiem – ustalenie czerni lub bieli
			if gray > threshold {
				out.Set(x, y, color.White)
			} else {
				out.Set(x, y, color.Black)
			}
		}
	}

	return out // zwracamy przetworzony obraz
}

// --------- PRZETWARZANIE OBRAZU I GENEROWANIE TABLICY C ---------

// ProcessImage – łączy konwersję monochromatyczną i generowanie tablicy C
func ProcessImage(img image.Image, thr uint8) (image.Image, string) {
	p := ToMonochrome(img, thr) // konwersja na czarno-biały
	code := ImageToCArray(p)    // generowanie tablicy C z obrazu
	return p, code              // zwracamy zarówno obraz, jak i kod
}
