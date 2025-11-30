package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Zapisuje tablicę C do pliku w folderze bitmap
func SaveCArrayToFile(carray string, filename string) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("cannot get executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)

	// folder bitmap
	bitmapDir := filepath.Join(exeDir, "bitmap")
	err = os.MkdirAll(bitmapDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("cannot create bitmap folder: %v", err)
	}

	// pełna ścieżka do pliku
	filePath := filepath.Join(bitmapDir, filename)

	err = os.WriteFile(filePath, []byte(carray), 0644)
	if err != nil {
		return fmt.Errorf("cannot write file: %v", err)
	}

	return nil
}
