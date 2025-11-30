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

var appname = "SunBitmap_Array_Tool v."
var version = "0.0.01"

func main() {
	a := app.NewWithID("com.lothar-TEAM.bitmaptool")

	InitSettingsPath()
	LoadSettings()
	ApplyTheme(a)

	StartGUI(a)
	a.Run()
}
