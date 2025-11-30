/*
================================================================================
File:        theme.go
Description: Funkcje odpowiedzialne za dynamiczne przełączanie motywu
             Dark/Light w aplikacji.
Author:      SunRiver / Lothar Team
Website:     https://forum.lothar-team.pl/
Version:     0.0.01
Date:        2025-11-30
================================================================================
*/

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
