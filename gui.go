/*
================================================================================
File:        gui.go
Description: Definiuje wszystkie elementy GUI aplikacji przy użyciu Fyne,
             obsługę przycisków, sliderów i podziałów okien.
Author:      SunRiver / Lothar Team
Website:     https://forum.lothar-team.pl/
Version:     0.0.01
Date:        2025-11-30
================================================================================
*/

package main

import (
	"image"
	"image/png"
	"path/filepath"

	// Fyne – framework do tworzenia GUI w Go
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// --------- GLOBALNE ZMIENNE UI ---------
// Deklaracja przycisków, suwaka i etykiety, aby były dostępne w całej funkcji RefreshUI
var (
	btnOpen      *widget.Button
	lblThreshold *widget.Label
	thresholdSld *widget.Slider
	btnLanguage  *widget.Button
	btnTheme     *widget.Button
	btnSave      *widget.Button
)

// --------- FUNKCJA ODŚWIEŻANIA UI ---------
// Aktualizuje teksty przycisków i etykiet w zależności od języka lub stanu aplikacji

func RefreshUI() {
	btnOpen.SetText(T("open_image"))           // "Open Image" w aktualnym języku
	lblThreshold.SetText(T("threshold"))       // "Threshold" w aktualnym języku
	btnLanguage.SetText(LanguageButtonLabel()) // Tekst przycisku zmiany języka
	btnSave.SetText(T("save_bitmap"))          // "Save Bitmap" w aktualnym języku
}

// --------- FUNKCJA URUCHAMIAJĄCA GUI ---------

func StartGUI(a fyne.App) {
	// Tworzymy główne okno aplikacji z nazwą i wersją
	w := a.NewWindow(appname + version)

	var original image.Image // przechowuje oryginalny obraz wczytany przez użytkownika

	// Obiekty canvas do wyświetlania obrazów w GUI
	imgOrig := canvas.NewImageFromImage(nil) // oryginalny obraz
	imgProc := canvas.NewImageFromImage(nil) // przetworzony obraz (po threshold)

	// --------- ELEMENTY UI ---------
	btnOpen = widget.NewButton(T("open_image"), nil)           // przycisk do otwierania obrazów
	lblThreshold = widget.NewLabel(T("threshold"))             // etykieta suwaka progowego
	thresholdSld = widget.NewSlider(0, 255)                    // suwak progowy (0-255)
	thresholdSld.Value = 128                                   // wartość domyślna
	btnLanguage = widget.NewButton(LanguageButtonLabel(), nil) // przycisk zmiany języka
	btnTheme = widget.NewButton(ThemeButtonLabel(), nil)       // przycisk zmiany motywu

	codeOutput := widget.NewMultiLineEntry() // pole tekstowe do wyświetlania wygenerowanego kodu C

	// --------- LOGIKA PRZYCISKÓW ---------

	// Przycisk "Open Image"
	btnOpen.OnTapped = func() {
		dialog.NewFileOpen(func(rc fyne.URIReadCloser, err error) {
			if rc == nil { // jeśli użytkownik anulował
				return
			}
			defer func() {
				if err := rc.Close(); err != nil {
					dialog.ShowError(err, w)
				}
			}()

			// Dekodowanie obrazu PNG
			img, err := png.Decode(rc)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			// Aktualizacja oryginalnego obrazu
			original = img
			imgOrig.Image = img
			imgOrig.Refresh()

			// Przetwarzanie obrazu wg progu i generowanie kodu C
			processed, code := ProcessImage(original, uint8(thresholdSld.Value))
			imgProc.Image = processed
			imgProc.Refresh()
			codeOutput.SetText(code)

		}, w).Show()
	}

	// Suwak progu – automatycznie przetwarza obraz przy zmianie wartości
	thresholdSld.OnChanged = func(v float64) {
		if original != nil {
			processed, code := ProcessImage(original, uint8(v))
			imgProc.Image = processed
			imgProc.Refresh()
			codeOutput.SetText(code)
		}
	}

	// Przycisk zmiany języka
	btnLanguage.OnTapped = func() {
		CurrentLang = NextLanguage() // ustawienie kolejnego języka
		RefreshUI()                  // odświeżenie wszystkich tekstów w GUI
		w.Content().Refresh()        // wymuszenie odświeżenia okna
	}

	// Przycisk zmiany motywu
	btnTheme.OnTapped = func() {
		ToggleTheme(a)                       // zmiana motywu aplikacji
		btnTheme.SetText(ThemeButtonLabel()) // aktualizacja tekstu przycisku
	}

	// Przycisk "Save Bitmap"
	btnSave = widget.NewButton(T("save_bitmap"), func() {
		if original == nil { // jeśli nie ma wczytanego obrazu
			dialog.ShowInformation("Info", "No image loaded", w)
			return
		}

		entry := widget.NewEntry() // pole tekstowe na nazwę pliku
		entry.SetPlaceHolder("auto.h")

		// Dialog do zapisania pliku
		d := dialog.NewForm("Save Bitmap", "Save", "Cancel",
			[]*widget.FormItem{
				widget.NewFormItem("File name", entry),
			},
			func(ok bool) {
				if !ok || entry.Text == "" {
					return
				}

				filename := entry.Text
				if filepath.Ext(filename) != ".h" { // dodaj .h jeśli brak
					filename += ".h"
				}

				// Przetwarzanie obrazu i zapis do pliku
				_, code := ProcessImage(original, uint8(thresholdSld.Value))
				err := SaveCArrayToFile(code, filename)
				if err != nil {
					dialog.ShowError(err, w)
				} else {
					dialog.ShowInformation("Saved", "File saved to ./bitmap/"+filename, w)
				}
			}, w)
		d.Show()
	})

	// --------- UKŁAD GUI ---------

	// Równomiernie rozciągnięte przyciski w wierszu (język, motyw, zapis)
	langThemeRow := container.New(
		layout.NewGridLayout(3),
		btnLanguage,
		btnTheme,
		btnSave,
	)

	// Podział pionowy dla obrazów (oryginalny vs przetworzony)
	leftSplit := container.NewVSplit(imgOrig, imgProc)
	leftSplit.Offset = 0.5 // 50% wysokości dla obu obrazów

	// Podział poziomy: lewe panele (obrazy) + prawe (kod C)
	mainSplit := container.NewHSplit(leftSplit, codeOutput)
	mainSplit.Offset = 0.6 // 60% szerokości dla obrazów, reszta dla kodu

	// Górny panel z przyciskami i suwakiem
	topPanel := container.NewVBox(
		btnOpen,
		lblThreshold,
		thresholdSld,
		langThemeRow,
	)

	// Ustawienie głównej zawartości okna: topPanel na górze, reszta w mainSplit
	w.SetContent(container.NewBorder(topPanel, nil, nil, nil, mainSplit))
	w.Resize(fyne.NewSize(1200, 700)) // domyślny rozmiar okna
	w.Show()                          // pokazanie okna
}
