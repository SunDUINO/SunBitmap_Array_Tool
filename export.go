/*
================================================================================
File:        export.go
Description: Zawiera funkcje odpowiedzialne za eksportowanie przetworzonych
             bitmap do tablicy C lub Rust.
Author:      SunRiver / Lothar Team
Website:     https://forum.lothar-team.pl/
Version:     0.0.01
Date:        2025-11-30
================================================================================
*/

package main

import (
	"bytes"       // do budowania ciągów znaków w sposób wydajny
	"fmt"         // do formatowania ciągów (np. sprintf)
	"image"       // podstawowe struktury dla obrazów
	"image/color" // dostęp do predefiniowanych kolorów, np. color.White
)

// ImageToCArray Funkcja konwertuje obraz (image.Image) do tablicy bajtów w formacie C
func ImageToCArray(img image.Image) string {
	b := img.Bounds()      // Pobranie prostokątnych granic obrazu
	w, h := b.Dx(), b.Dy() // Szerokość i wysokość obrazu

	var buf bytes.Buffer
	buf.WriteString("const uint8_t img[] = {\n") // Rozpoczynamy tworzenie tablicy C

	// Iteracja po każdym wierszu obrazu
	for y := 0; y < h; y++ {
		var byteVal uint8 // Będzie przechowywać 8 pikseli w jednym bajcie
		bit := 0          // Licznik bitów w bieżącym bajcie

		// Iteracja po każdym pikselu w wierszu
		for x := 0; x < w; x++ {
			isWhite := img.At(x, y) == color.White // Sprawdzenie, czy piksel jest biały
			if isWhite {
				// Ustawienie odpowiedniego bitu na 1, jeśli piksel jest biały
				byteVal |= 1 << (7 - bit) // likwidacja parantezy wokół (1 << (7 - bit)) niepotrzebne, bo operator bitowy |= ma niższy priorytet niż <<
			}
			bit++ // Przechodzimy do następnego bitu

			// Jeśli mamy pełny bajt (8 bitów), zapisujemy go do buf
			if bit == 8 {
				buf.WriteString(fmt.Sprintf("0x%02X, ", byteVal))
				bit = 0     // reset bitu
				byteVal = 0 // reset wartości bajtu
			}
		}
		// Przejście do nowej linii po każdym wierszu obrazu
		buf.WriteString("\n")
	}

	buf.WriteString("};\n") // Zakończenie deklaracji tablicy C
	return buf.String()     // Zwracamy gotowy ciąg znaków
}
