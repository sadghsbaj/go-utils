package ch

import "strings"
import "fmt"

func FormatTerminal(text, color string, bold, underline bool) string {
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

	// Eine Liste, um die anzuwendenden Format-Codes zu sammeln
	var codes []string

	if bold {
		codes = append(codes, "1")
	}
	if underline {
		codes = append(codes, "4")
	}

	// Farbcode hinzufügen, wenn eine gültige Farbe angegeben wurde
	if colorCode, ok := colorMap[strings.ToLower(color)]; ok {
		codes = append(codes, colorCode)
	}

	// Wenn keine Formatierung gewünscht ist, den Originaltext zurückgeben
	if len(codes) == 0 {
		return text
	}

	formatCode := strings.Join(codes, ";")

	return fmt.Sprintf("\x1b[%sm%s\x1b[0m", formatCode, text)
}
