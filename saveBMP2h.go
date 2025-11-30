/*
================================================================================
File:        saveBMP2h.go
Description: Funkcje zapisujące przetworzone bitmapy jako pliki .h
             w lokalnym folderze bitmap.
Author:      SunRiver / Lothar Team
Website:     https://forum.lothar-team.pl/
Version:     0.0.01
Date:        2025-11-30
================================================================================
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// --------- ZAPIS TABLICY C DO PLIKU ---------

// SaveCArrayToFile – zapisuje tekstową reprezentację tablicy C do pliku w katalogu "bitmap"
// carray  – ciąg znaków zawierający kod C (np. wygenerowany przez ImageToCArray)
// filename – nazwa pliku, np. "auto.h"
// Zwraca błąd w przypadku problemów z utworzeniem folderu lub zapisem pliku
func SaveCArrayToFile(carray string, filename string) error {
	// Pobranie ścieżki do uruchamianego pliku wykonywalnego
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("cannot get executable path: %v", err)
	}

	// Wyciągnięcie katalogu z pełnej ścieżki pliku wykonywalnego
	exeDir := filepath.Dir(exePath)

	// Utworzenie katalogu "bitmap" w katalogu aplikacji
	bitmapDir := filepath.Join(exeDir, "bitmap")
	err = os.MkdirAll(bitmapDir, os.ModePerm) // tworzy wszystkie brakujące katalogi
	if err != nil {
		return fmt.Errorf("cannot create bitmap folder: %v", err)
	}

	// Pełna ścieżka do pliku docelowego
	filePath := filepath.Join(bitmapDir, filename)

	// Zapisanie zawartości tablicy C do pliku (0644 = prawa dostępu do pliku)
	err = os.WriteFile(filePath, []byte(carray), 0644)
	if err != nil {
		return fmt.Errorf("cannot write file: %v", err)
	}

	return nil // brak błędów – zapis zakończony sukcesem
}
