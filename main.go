/*
================================================================================
File:        main.go
Description: Punkt wejścia aplikacji; inicjalizacja ustawień, motywu, GUI
             i uruchomienie głównej pętli aplikacji.
Author:      SunRiver / Lothar Team
Website:     https://forum.lothar-team.pl/
Version:     0.0.01
Date:        2025-11-30
================================================================================
*/

package main

import "fyne.io/fyne/v2/app"

// --------- STAŁE APLIKACJI ---------
var appname = "SunBitmap_Array_Tool v." // nazwa aplikacji wyświetlana w GUI
var version = "0.0.3"                   // wersja aplikacji

// --------- FUNKCJA GŁÓWNA ---------
func main() {
	// Tworzymy nową aplikację Fyne z unikalnym identyfikatorem (ID) – przydatne np. do ustawień
	a := app.NewWithID("com.lothar-TEAM.bitmaptool")

	// --------- INICJALIZACJA ---------
	InitSettingsPath() // ustawienie ścieżki do pliku konfiguracyjnego
	LoadSettings()     // wczytanie zapisanych ustawień użytkownika
	ApplyTheme(a)      // zastosowanie motywu (ciemny/jasny) do aplikacji

	// --------- URUCHOMIENIE INTERFEJSU ---------
	StartGUI(a) // stworzenie okna, layoutu i wszystkich elementów GUI

	a.Run() // uruchomienie głównej pętli aplikacji
}
