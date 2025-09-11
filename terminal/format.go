package terminal

import "strings"
import "fmt"
import "sort"
import "strconv"

// Wird benötigt um aus der Map ein sortierbares Slice zu erstellen
type KeyValue struct {
	Key string
	Value int
}

// Format formatiert einen gegebenen Text für die Terminal-Ausgabe
// mit Farbe, Fett- und Unterstrichen-Option.
func Format(text, color string, bold, underline bool) (string, error) {
	colorMap := map[string]string{
		// Standardfarben
		"black":   "30",
		"red":     "31",
		"green":   "32",
		"yellow":  "33",
		"blue":    "34",
		"magenta": "35",
		"cyan":    "36",
		"white":   "37",
		// Helle/Intensive Farben
		"bright-black":   "90", // oft als Grau dargestellt
		"bright-red":     "91",
		"bright-green":   "92",
		"bright-yellow":  "93",
		"bright-blue":    "94",
		"bright-magenta": "95",
		"bright-cyan":    "96",
		"bright-white":   "97",
	}

	// Konstanten für die Ansi-Code Darstellung
	const (
		escapeCode = "\x1b["
		resetCode  = "\x1b[0m"
	)

	// Eine Liste, um die anzuwendenden Format-Codes zu sammeln
	var codes []string

	if bold {
		codes = append(codes, "1")
	}
	if underline {
		codes = append(codes, "4")
	}

	// Farbcode hinzufügen, wenn eine gültige Farbe angegeben wurde, ansonsten Fehlermeldung mit verfügbaren Farben ausgeben
	if colorCode, ok := colorMap[strings.ToLower(color)]; ok {
		codes = append(codes, colorCode)
	} else {
		availableColors, e := sortMap(colorMap)
		if e != nil {
			return "", e
		}

		availableColorsString := strings.Join(availableColors, ", ")
		return "", fmt.Errorf("Ungültige Farbe! Folgende Farben sind verfügbar:\n%s", availableColorsString)
	}

	// Wenn keine Formatierung gewünscht ist, den Originaltext zurückgeben
	if len(codes) == 0 {
		return text, nil
	}

	formatCode := strings.Join(codes, ";")

	return fmt.Sprintf("%s%sm%s%s", escapeCode, formatCode, text, resetCode), nil
}

// Interne Hilfsfunktion für Format Terminal.
// Erstellt eine sortierte Liste aller möglichen Farben, um diese für die Fehlermeldung ausgeben zu können.
func sortMap(colorMap map[string]string) ([]string, error) {
	colorSlice := []KeyValue{}

	// Für Jeden Map Eintrag Value in Int konvertieren und Slice aus Structs vom Type KeyValue erstellen
	for k, v := range colorMap {
		valueToInt, e := strconv.Atoi(v)
		if e != nil {
			return nil, fmt.Errorf("interner fehler in colormap bei schlüssel '%s': %w", k, e)
		}

		colorSlice = append(colorSlice, KeyValue{Key: k, Value: valueToInt})
	}

	// Slice nach Value Wert sortieren
	sort.Slice(colorSlice, func(x, y int) bool {
		return colorSlice[x].Value < colorSlice[y].Value
	})

	// Neuer Slice nur für die Keys - Value Werte werden nicht benötigt
	keySlice := []string{}
	for _, v := range colorSlice {
		keySlice = append(keySlice, v.Key)
	}

	return keySlice, nil
}
