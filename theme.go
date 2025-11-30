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
	"fyne.io/fyne/v2" // framework GUI
	"fyne.io/fyne/v2/theme"
)

// --------- GLOBALNE ZMIENNE ---------

// IsDark przechowuje aktualny stan motywu: true = ciemny, false = jasny
var IsDark = false

// --------- ETYKIETA PRZYCISKU MOTYWU ---------

// ThemeButtonLabel – zwraca symbol, który pojawia się na przycisku zmiany motywu
// 🌙 = przełącz na ciemny, ☀️ = przełącz na jasny
func ThemeButtonLabel() string {
	if IsDark {
		return "☀️" // jeśli aktualnie ciemny, przycisk pokazuje ikonę słoneczka (zmiana na jasny)
	}
	return "🌙" // jeśli aktualnie jasny, przycisk pokazuje ikonę księżyca (zmiana na ciemny)
}

// --------- PRZEŁĄCZANIE MOTYWU ---------

// ToggleTheme – zmienia motyw aplikacji na przeciwny (ciemny ↔ jasny)
func ToggleTheme(a fyne.App) {
	IsDark = !IsDark // zmiana stanu motywu

	// Ustawienie motywu w Fyne
	if IsDark {
		a.Settings().SetTheme(theme.DarkTheme()) // tryb ciemny (deprecated w nowych wersjach Fyne)
	} else {
		a.Settings().SetTheme(theme.LightTheme()) // tryb jasny
	}

	SaveSettings() // zapis nowego stanu motywu do pliku ustawień
}
