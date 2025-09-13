package terminal

import "strings"
import "fmt"
import "os"
import "sort"
import "strconv"

// Konstanten für Fehlerlevel
const (
	Info 	= "info"
	Warning = "warning"
	Error 	= "error"
	Fatal 	= "fatal"
)

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
		// Verfügbare Farben sortieren und in Slice anlegen
		availableColors, e := sortMap(colorMap)
		if e != nil {
			return "", e
		}

		// Color Labels der verfügbaren Farben im Color Slice einfärben
		formattedAvailableColors := formatAvailableColorLabels(colorMap, availableColors, escapeCode, resetCode)

		// Slice zu String getrennt mit ',' umwandeln
		formattedAvailableColorsString := strings.Join(formattedAvailableColors, ", ")
		return "", fmt.Errorf("Ungültige Farbe! Folgende Farben sind verfügbar:\n%s", formattedAvailableColorsString)
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

// Interne Hilfsfunktion für Format Terminal.
// Nimmt die sortierte Liste der sortMap und färbt die einzelnen Color Labels.
func formatAvailableColorLabels(colorMap map[string]string, avColorSlice []string, escape string, reset string) ([]string) {
	// Neuen Slice für eingefärbte Color Label initalisieren
	var formattedColorsSlice []string

	// Jedes Color Label einfärben
	for _, v := range avColorSlice {
		formatCode := colorMap[v]
		colorLabel := fmt.Sprintf("%s%sm%s%s", escape, formatCode, v, reset)
		formattedColorsSlice = append(formattedColorsSlice, colorLabel)
	}

	return formattedColorsSlice
}

// Verwwendet die Format() Funktion um vordefinierte Alerts zu formatieren.
// Formatiert für die Level Info Warning Error Fatal-Error
func FormatAlert(msg string, level string, e error) (string, error) {
	// Zur Sicherheit prüfen ob wirklich ein Fehler vorhanden ist, außer Level ist Info
	if e == nil && level != "info" {
		return "", nil
	}

	// Je nach Level formatieren
	var label string
	var color string

	switch strings.ToLower(level) {
		case "info":
			label = "[INFO]"
			color = "blue"

		case "warning":
			label = "[WARNING]"
			color = "yellow"

		case "error":
			label = "[ERROR]"
			color = "red"

		case "fatal":
			label = "[FATAL-ERROR]"
			color = "bright-red"

		default:
			validLevels := []string{Info, Warning, Error, Fatal}

			// Fehler ausgeben dass ungültiger Level verwendet wurde und verfügbare Level auflisten (Magenta um aufzufallen)
			errorPrefix, err := Format("[FORMAT-ALERT-ERROR]", "magenta", true, false)
			if err != nil {
				fmt.Fprintf(
					os.Stderr,
					"[ERROR] 'FORMAT-ALERT-ERROR' konnte nicht formatiert werden.",
				)
			}

			fmt.Fprintf(
				os.Stderr,
				"%s Ungültiger Level '%s'. Bitte einen der folgenden Level verwenden: %s\n",
				errorPrefix,
				level,
				strings.Join(validLevels, ", "),
			)

			// Ursprünglicher Fehler soll trz. nicht verloren geben, deshalb mit label "Undefined" und weißer Schrift ausgeben
			label = fmt.Sprintf("[%s-UNDEFINED]", strings.ToUpper(level))
			color = "magenta"
	}

	// Prefix formatieren
	prefix, err := Format(label, color, true, false)
	if err != nil {
		return "", err
	}

	// Nachricht zusammensetzen, falls Level Info keinen Fehler 'e' ausgeben hat
	if level == "info" {
		msg = fmt.Sprintf("%s %s\n", prefix, msg)
		return msg, nil
	}

	msg = fmt.Sprintf("%s %s %v\n", prefix, msg, e)

	return msg, nil
}
