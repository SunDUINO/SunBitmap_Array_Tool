/*
================================================================================
File:        settings.go
Description: Zarządzanie ustawieniami aplikacji (język, motyw),
             wczytywanie i zapisywanie do settings.json.
Author:      SunRiver / Lothar Team
Website:     https://forum.lothar-team.pl/
Version:     0.0.01
Date:        2025-11-30
================================================================================
*/

package main

import (
	"encoding/json" // do kodowania/odkodowywania JSON

	"os"            // do operacji na plikach
	"path/filepath" // do obsługi ścieżek plików

	"fyne.io/fyne/v2" // framework GUI
	"fyne.io/fyne/v2/theme"
)

// --------- GLOBALNE ZMIENNE ---------

var SettingsPath string // pełna ścieżka do pliku ustawień

// Struktura przechowująca ustawienia aplikacji

type Settings struct {
	Language string `json:"language"`  // język aplikacji: "en" lub "pl"
	DarkMode bool   `json:"dark_mode"` // tryb ciemny: true/false
}

// Domyślne ustawienia aplikacji

var AppSettings = Settings{
	Language: "en",
	DarkMode: false,
}

const SettingsFile = "settings.json" // nazwa pliku ustawień

// --------- INICJALIZACJA ŚCIEŻKI ---------

func InitSettingsPath() {
	exePath, err := os.Executable() // pobranie ścieżki do uruchomionej aplikacji
	if err != nil {
		println("Cannot get executable path:", err.Error())
		SettingsPath = "settings.json" // fallback – zapis w bieżącym katalogu
		return
	}

	exeDir := filepath.Dir(exePath)                    // katalog aplikacji
	SettingsPath = filepath.Join(exeDir, SettingsFile) // pełna ścieżka do settings.json
}

// --------- WŁĄCZENIE USTAWIEŃ Z PLIKU ---------

func LoadSettings() {
	data, err := os.ReadFile(SettingsPath)
	if err != nil {
		println("Settings file not found, using defaults.")
		return
	}

	err = json.Unmarshal(data, &AppSettings)
	if err != nil {
		println("Error decoding settings JSON:", err.Error())
		return
	}

	if AppSettings.Language == "pl" {
		CurrentLang = PL
	} else {
		CurrentLang = EN
	}

	IsDark = AppSettings.DarkMode
}

// --------- ZAPIS USTAWIEŃ DO PLIKU ---------

func SaveSettings() {
	AppSettings.Language = string(CurrentLang) // aktualny język
	AppSettings.DarkMode = IsDark              // aktualny tryb ciemny

	// Serializacja do JSON z wcięciami (czytelne dla człowieka)
	data, err := json.MarshalIndent(AppSettings, "", "  ")
	if err != nil {
		println("JSON marshal error:", err.Error())
		return
	}

	// Zapis do pliku JSON
	err = os.WriteFile(SettingsPath, data, 0644)
	if err != nil {
		println("Write settings error:", err.Error())
		return
	}

	println("Settings saved:", SettingsPath)
}

// --------- ZASTOSOWANIE MOTYWU W APLIKACJI ---------

func ApplyTheme(a fyne.App) {
	if IsDark {
		a.Settings().SetTheme(theme.DarkTheme()) // tryb ciemny
	} else {
		a.Settings().SetTheme(theme.LightTheme()) // tryb jasny
	}
}
