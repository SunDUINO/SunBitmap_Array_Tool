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
