package main

type Lang string

const (
	PL Lang = "pl"
	EN Lang = "en"
)

var CurrentLang = EN

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

func T(key string) string {
	if v, ok := dict[CurrentLang][key]; ok {
		return v
	}
	return key
}

func NextLanguage() Lang {
	if CurrentLang == EN {
		return PL
	}
	return EN
}

// To co będzie na przycisku:
func LanguageButtonLabel() string {
	if CurrentLang == EN {
		return "🇵🇱 PL"
	}
	SaveSettings()
	return "🇬🇧 EN"

}
