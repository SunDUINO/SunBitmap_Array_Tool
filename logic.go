// ==================================================================================
// File:        logic.go
// Description: Logika przetwarzania obrazów: wczytywanie (multi-format),
//              konwersja do monochromu, skalowanie (resize) oraz generowanie
//              tablicy C (wykorzystuje export.ImageToCArray).
// Author:      SunRiver / Lothar Team
// Website:     https://forum.lothar-team.pl/
// Version:     0.0.02
// Date:        2025-11-30
// ==================================================================================

package main

import (
	"image"
	"image/color"
	_ "image/gif"  // rejestracja formatu
	_ "image/jpeg" // rejestracja formatu
	_ "image/png"  // rejestracja formatu

	// dodatkowe dekodery (bmp, webp)
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/webp"

	"golang.org/x/image/draw"
)

// ToGray converts any image.Image to *image.Gray using standard luminance formula.
func ToGray(src image.Image) *image.Gray {
	if src == nil {
		return nil
	}
	b := src.Bounds()
	dst := image.NewGray(image.Rect(0, 0, b.Dx(), b.Dy()))
	// używamy draw.Draw do zachowania wydajności + poprawnej konwersji
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}

// ToMonochrome – konwertuje obraz do czarno-białego (image.Gray) na podstawie progu (0-255).
// Piksele z jasnością > threshold ustawione są na 255 (white), w przeciwnym wypadku 0 (black).
func ToMonochrome(img image.Image, threshold uint8) *image.Gray {
	if img == nil {
		return nil
	}
	gray := ToGray(img)
	b := gray.Bounds()
	out := image.NewGray(image.Rect(0, 0, b.Dx(), b.Dy()))

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			Y := gray.GrayAt(x, y).Y
			if Y > threshold {
				out.SetGray(x-b.Min.X, y-b.Min.Y, color.Gray{Y: 255})
			} else {
				out.SetGray(x-b.Min.X, y-b.Min.Y, color.Gray{Y: 0})
			}
		}
	}
	return out
}

// ProcessImage – główna funkcja używana w GUI: bierze oryginalny obraz i threshold,
// zwraca przetworzony obraz do wyświetlenia oraz tekstową tablicę C.
func ProcessImage(img image.Image, thr uint8) (image.Image, string) {
	if img == nil {
		return nil, "// empty image\n"
	}

	mono := ToMonochrome(img, thr)
	// domyślna nazwa "img" — GUI może nadpisać przy zapisie do pliku
	code := ImageToCArray(mono, "img")
	return mono, code
}

// ---------- FUNKCJE NIE UZYWANE W TEJ WERSJI  -------------------------------------------------------------------------
/*
// ProcessImageResize – dodatkowa funkcja: najpierw skalujemy obraz do (w,h),
// zachowując proporcje opcjonalnie (tu skalujemy bez zachowania proporcji — pixel perfect)
// następnie konwertujemy do monochromu i generujemy kod.
func ProcessImageResize(src image.Image, thr uint8, targetW, targetH int) (image.Image, string) {
	if src == nil {
		return nil, "// empty image\n"
	}
	if targetW <= 0 || targetH <= 0 {
		return ProcessImage(src, thr)
	}

	// przygotuj docelowy obraz o zadanych wymiarach
	dst := image.NewRGBA(image.Rect(0, 0, targetW, targetH))
	// możemy użyć interpoalcji CatmullRom lub NearestNeighbor dla pixel artu
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)

	// konwertujemy do monochromu
	mono := ToMonochrome(dst, thr)
	code := ImageToCArray(mono, fmt.Sprintf("img_%dx%d", targetW, targetH))
	return mono, code
}

// helper: ImageToPNGBytes – pomocnicza funkcja jeśli chcesz zapisać przetworzony obraz jako PNG (np. preview)
// zwraca []byte zawierające PNG - przydatne do zapisu pliku lub podglądu

func ImageToPNGBytes(img image.Image) ([]byte, error) {
	if img == nil {
		return nil, fmt.Errorf("nil image")
	}
	var buf bytes.Buffer
	if err := pngEncode(&buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}



// pngEncode - wydzielona enkapsulacja aby nie importować bezpośrednio image/png w nagłówku,
// ale wygodniej jest z niego skorzystać. Implementacja prostego wrappera:
func pngEncode(buf *bytes.Buffer, img image.Image) error {
	// import image/png lokalnie poprzez alias (żeby nie dodawać globalnego importu)
	// ale prostsze: użyj image/png bez aliasu - tu użyjemy bezpośrednio:
	return (&pngEncoder{}).Encode(buf, img)
}

//--------------------------------------------------------------------------------------------------------------------------


// minimalistyczny wrapper do kodowania PNG (bez eksportowania image/png globalnie)
type pngEncoder struct{}

func (p *pngEncoder) Encode(buf *bytes.Buffer, img image.Image) error {
	// importujemy standard image/png tutaj (funkcja wewnętrzna)
	// aby uniknąć błędów importu w topie pliku – po prostu użyjemy standardowej biblioteki
	// NOTE: poniżej korzystamy z image/png bezpośrednio
	return encodePNG(buf, img)
}

func encodePNG(buf *bytes.Buffer, img image.Image) error {
	return png.Encode(buf, img)
}
*/
