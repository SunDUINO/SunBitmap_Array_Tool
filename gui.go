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

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	btnOpen      *widget.Button
	lblThreshold *widget.Label
	thresholdSld *widget.Slider
	btnLanguage  *widget.Button
	btnTheme     *widget.Button
	btnSave      *widget.Button
)

func RefreshUI() {
	btnOpen.SetText(T("open_image"))
	lblThreshold.SetText(T("threshold"))
	btnLanguage.SetText(LanguageButtonLabel())
	btnSave.SetText(T("save_bitmap"))
}

func StartGUI(a fyne.App) {
	w := a.NewWindow(appname + version)

	var original image.Image

	imgOrig := canvas.NewImageFromImage(nil)
	imgProc := canvas.NewImageFromImage(nil)

	// --------- UI ELEMENTS ---------

	btnOpen = widget.NewButton(T("open_image"), nil)
	lblThreshold = widget.NewLabel(T("threshold"))
	thresholdSld = widget.NewSlider(0, 255)
	thresholdSld.Value = 128
	btnLanguage = widget.NewButton(LanguageButtonLabel(), nil)
	btnTheme = widget.NewButton(ThemeButtonLabel(), nil)

	codeOutput := widget.NewMultiLineEntry()

	// --------- LOGIKA PRZYCISKÓW ---------

	btnOpen.OnTapped = func() {
		dialog.NewFileOpen(func(rc fyne.URIReadCloser, err error) {
			if rc == nil {
				return
			}
			defer rc.Close()

			img, err := png.Decode(rc)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			original = img
			imgOrig.Image = img
			imgOrig.Refresh()

			processed, code := ProcessImage(original, uint8(thresholdSld.Value))
			imgProc.Image = processed
			imgProc.Refresh()
			codeOutput.SetText(code)

		}, w).Show()
	}

	thresholdSld.OnChanged = func(v float64) {
		if original != nil {
			processed, code := ProcessImage(original, uint8(v))
			imgProc.Image = processed
			imgProc.Refresh()
			codeOutput.SetText(code)
		}
	}

	btnLanguage.OnTapped = func() {
		CurrentLang = NextLanguage()
		RefreshUI()
		w.Content().Refresh()
	}

	btnTheme.OnTapped = func() {
		ToggleTheme(a)
		btnTheme.SetText(ThemeButtonLabel())
	}

	btnSave = widget.NewButton(T("save_bitmap"), func() {
		if original == nil {
			dialog.ShowInformation("Info", "No image loaded", w)
			return
		}

		entry := widget.NewEntry()
		entry.SetPlaceHolder("auto.h")

		d := dialog.NewForm("Save Bitmap", "Save", "Cancel",
			[]*widget.FormItem{
				widget.NewFormItem("File name", entry),
			},
			func(ok bool) {
				if !ok || entry.Text == "" {
					return
				}

				filename := entry.Text
				if filepath.Ext(filename) != ".h" {
					filename += ".h"
				}

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

	// równomiernie rozciągnięte przyciski
	langThemeRow := container.New(
		layout.NewGridLayout(3),
		btnLanguage,
		btnTheme,
		btnSave,
	)

	leftSplit := container.NewVSplit(imgOrig, imgProc)
	leftSplit.Offset = 0.5

	mainSplit := container.NewHSplit(leftSplit, codeOutput)
	mainSplit.Offset = 0.6

	topPanel := container.NewVBox(
		btnOpen,
		lblThreshold,
		thresholdSld,
		langThemeRow,
	)

	w.SetContent(container.NewBorder(topPanel, nil, nil, nil, mainSplit))
	w.Resize(fyne.NewSize(1200, 700))
	w.Show()
}
