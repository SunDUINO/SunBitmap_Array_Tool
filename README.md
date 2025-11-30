# SunBitmap Array Tool

**SunBitmap Array Tool** to aplikacja desktopowa napisana w **Go** z wykorzystaniem **Fyne GUI**, przeznaczona dla elektroników, programistów embedded i twórców firmware.  
Umożliwia przetwarzanie obrazów, konwersję bitmap do tablic C lub Rust oraz eksport gotowych plików nagłówkowych.

---

##⚡ Autor

SunRiver / Lothar Team
https://forum.lothar-team.pl/

---

## 🌟 Funkcje

- Otwieranie i podgląd obrazów **PNG**  
- Regulacja progu (threshold) w czasie rzeczywistym  
- Przetwarzanie bitmap z możliwością **ditheringu i oversamplingu**  
- Generowanie tablic **C / Rust** do użycia w mikrokontrolerach  
- Automatyczne dzielenie obrazów na **tiles** (np. 8x8 lub 16x16)  
- Eksport do pliku `.h` w folderze `bitmap/`  
- **Dynamiczna zmiana języka** (PL / EN)  
- **Dark / Light theme toggle**  
- Ustawienia zapisywane lokalnie w `settings.json`  

---

## 💻 Wymagania

- **Go** >= 1.21  
- System operacyjny: Windows / Linux / macOS  
- Pakiet GUI: [Fyne](https://fyne.io/) v2.7.x  

---

## 🚀 Instalacja

1. Sklonuj repozytorium:

```bash
git clone https://github.com/SunDUINO/SunBitmap_Array_Tool.git
cd SunBitmap_Array_Tool
```

2. Pobierz zależności:

```bash
go mod tidy
```

3. Uruchom aplikację:

```bash
go run main.go
```

4. Lub zbuduj plik wykonywalny:

```bash
go build -ldflags -H=windowsgui -o SunBitmap_Array_Tool.exe main.go
```

## 🖼️ Użycie

- Kliknij Open Image i wybierz plik PNG.
- Ustaw threshold sliderem, aby przetestować przetwarzanie.
- Kliknij 💾 Save Bitmap i wpisz nazwę pliku .h.
- Plik zostanie zapisany w folderze bitmap/ obok programu.
- Możesz zmieniać język i motyw dynamicznie klikając przyciski w górnym wierszu.

## 🌐 Tłumaczenia

PL – polski

EN – angielski

Obsługiwane dynamiczne przełączanie języka w GUI

## 📄 Licencja

Projekt udostępniony na licencji MIT.


---

## Objaśnienia 

Suwak „Threshold” steruje poziomem binarizacji obrazu.

🔍 Co to oznacza?

Binarizacja to proces przekształcania obrazu na czarno–biały (0 lub 1) na podstawie poziomu jasności pikseli.
Suwak ustawia wartość progu od 0 do 255 czyli:

Każdy piksel jaśniejszy niż próg → staje się biały (1) <br>
Każdy piksel ciemniejszy niż próg → staje się czarny (0)


## 🔧 TODO / plan rozwoju

 --- Dodanie różnych metod ditheringu (Floyd–Steinberg, Atkinson)

 --- Możliwość eksportu do innych formatów firmware

 --- Rozbudowany podgląd tiles / zoom

 --- Obsługa większej liczby języków w GUI

