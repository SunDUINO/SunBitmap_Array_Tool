package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var IsDark = false

func ThemeButtonLabel() string {
	if IsDark {
		return "☀️"
	}
	return "🌙"
}

func ToggleTheme(a fyne.App) {
	IsDark = !IsDark
	if IsDark {
		a.Settings().SetTheme(theme.DarkTheme())
	} else {
		a.Settings().SetTheme(theme.LightTheme())
	}
	SaveSettings()
}
