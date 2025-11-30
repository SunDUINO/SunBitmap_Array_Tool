/*
================================================================================
File:        i18n.go
Description: Zawiera definicje słowników językowych, funkcje do dynamicznej
             zmiany języka i pobierania tłumaczeń.
Author:      SunRiver / Lothar Team
Website:     https://forum.lothar-team.pl/
Version:     0.0.01
Date:        2025-11-30
================================================================================
*/

package main

// --------- TYPY I STAŁE ---------

// Lang to typ reprezentujący język w aplikacji
type Lang string

// Dostępne języki: polski i angielski
const (
	PL Lang = "pl"
	EN Lang = "en"
)

// Aktualnie wybrany język (domyślnie angielski)

var CurrentLang = EN

// --------- SŁOWNIK TRANSLACJI ---------

// dict to mapa map, która przechowuje tłumaczenia tekstów w różnych językach
// Klucz zewnętrzny: język (Lang), klucz wewnętrzny: identyfikator tekstu
var dict = map[Lang]map[string]string{
	EN: {
		"open_image":  "Open Image",
		"threshold":   "Threshold",
		"language":    "Language",
		"save_bitmap": "💾 Save Bitmap",
	},
	PL: {
		"open_image":  "Otwórz obraz",
		"threshold":   "Próg",
		"language":    "Język",
		"save_bitmap": "💾 Zapisz bitmapę",
	},
}

// --------- FUNKCJE POMOCNICZE ---------

// T – tłumaczy dany klucz na aktualnie wybrany język
func T(key string) string {
	if v, ok := dict[CurrentLang][key]; ok { // sprawdzenie, czy klucz istnieje w słowniku
		return v // zwróć tłumaczenie
	}
	return key // jeśli brak tłumaczenia, zwróć sam klucz
}

// NextLanguage – zwraca kolejny język (przełączenie między EN i PL)
func NextLanguage() Lang {
	if CurrentLang == EN {
		return PL
	}
	return EN
}

// LanguageButtonLabel – tekst, który pojawi się na przycisku zmiany języka
func LanguageButtonLabel() string {
	if CurrentLang == EN {
		return "🇵🇱 PL" // jeśli aktualnie EN, przycisk pokazuje opcję PL
	}
	SaveSettings() // jeśli aktualnie PL, zapis ustawień (np. języka)
	return "🇬🇧 EN" // przycisk pokazuje opcję EN
}
