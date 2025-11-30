// ==================================================================================
// File:        export.go
// Description: Zawiera funkcje odpowiedzialne za eksportowanie przetworzonych
//              bitmap do tablicy C (format dogodny do wklejenia w .h).
// Author:      SunRiver / Lothar Team
// Website:     https://forum.lothar-team.pl/
// Version:     0.0.02
// Date:        2025-11-30
// ==================================================================================

package main

import (
	"bytes"
	"fmt"
	"image"
)

// ImageToCArray konwertuje obraz (dowolny image.Image) do formatu tablicy C,
// gdzie każdy bajt zawiera 8 pikseli (MSB = lewy/górny piksel).
// Jeśli szerokość nie jest podzielna przez 8 - ostatni bajt w wierszu jest dopisywany
// z bitami ustawionymi od lewej do prawej (reszta bitów = 0).
//
// Zwraca string gotowy do wklejenia do pliku .h
func ImageToCArray(img image.Image, name string) string {
	if img == nil {
		return "// <empty image>\n"
	}

	b := img.Bounds()
	w, h := b.Dx(), b.Dy()

	var buf bytes.Buffer

	// nagłówek komentarza z wymiarami
	buf.WriteString(fmt.Sprintf("/* width: %d, height: %d */\n", w, h))
	buf.WriteString(fmt.Sprintf("const uint8_t %s[] = {\n", sanitizeName(name)))

	// iterujemy wierszami
	for y := 0; y < h; y++ {
		var byteVal uint8 = 0
		bit := 0

		for x := 0; x < w; x++ {
			// wyciągamy jasność piksela (0..255)
			r, g, bcol, _ := img.At(b.Min.X+x, b.Min.Y+y).RGBA()
			// r,g,bcol są w formacie 0..65535 -> konwertujemy do 0..255
			gray := uint8(((r>>8)*30 + (g>>8)*59 + (bcol>>8)*11) / 100)

			// zdecydujemy, kiedy piksel jest "on" — tu >127 (można zmienić)
			if gray > 127 {
				byteVal |= 1 << (7 - bit)
			}

			bit++
			if bit == 8 {
				buf.WriteString(fmt.Sprintf("0x%02X, ", byteVal))
				bit = 0
				byteVal = 0
			}
		}

		// jeśli po zakończeniu wiersza zostały bity (szerokość niepodzielna przez 8)
		if bit != 0 {
			buf.WriteString(fmt.Sprintf("0x%02X, ", byteVal))
		}

		buf.WriteString("\n")
	}

	buf.WriteString("};\n")
	return buf.String()
}

// sanitizeName robi prostą normalizację nazwy (usunięcie spacji i niedozwolonych znaków)
// aby wynikowa nazwa była poprawna jako identyfikator w C.
// Dla prostoty: zamieniamy spacje na '_' i jeśli pusta nazwa - używamy "bitmap".
func sanitizeName(name string) string {
	if name == "" {
		return "bitmap"
	}
	out := make([]rune, 0, len(name))
	for _, r := range name {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_':
			out = append(out, r)
		case r == ' ' || r == '-' || r == '.':
			out = append(out, '_')
			// ignorujemy pozostałe znaki
		}
	}
	if len(out) == 0 {
		return "bitmap"
	}
	return string(out)
}
