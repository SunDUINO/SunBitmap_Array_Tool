package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var SettingsPath string

type Settings struct {
	Language string `json:"language"`
	DarkMode bool   `json:"dark_mode"`
}

var AppSettings = Settings{
	Language: "en",
	DarkMode: false,
}

const SettingsFile = "settings.json"

func InitSettingsPath() {
	exePath, err := os.Executable()
	if err != nil {
		println("Cannot get executable path:", err.Error())
		SettingsPath = "settings.json" // fallback
		return
	}

	exeDir := filepath.Dir(exePath)
	SettingsPath = filepath.Join(exeDir, "settings.json")
}

// ------ LOAD SETTINGS ------
func LoadSettings() {
	data, err := os.ReadFile(SettingsPath)
	if err != nil {
		println("Settings file not found, using defaults.")
		return
	}

	json.Unmarshal(data, &AppSettings)

	if AppSettings.Language == "pl" {
		CurrentLang = PL
	} else {
		CurrentLang = EN
	}

	IsDark = AppSettings.DarkMode
}

// ------ SAVE SETTINGS ------
func SaveSettings() {
	AppSettings.Language = string(CurrentLang)
	AppSettings.DarkMode = IsDark

	data, err := json.MarshalIndent(AppSettings, "", "  ")
	if err != nil {
		println("JSON marshal error:", err.Error())
		return
	}

	err = os.WriteFile(SettingsPath, data, 0644)
	if err != nil {
		println("Write settings error:", err.Error())
		return
	}

	println("Settings saved:", SettingsPath)
}

// ------ APPLY THEME ------
func ApplyTheme(a fyne.App) {
	if IsDark {
		a.Settings().SetTheme(theme.DarkTheme())
	} else {
		a.Settings().SetTheme(theme.LightTheme())
	}
}
