package ch

import (
	"fmt"
	"strings"
	"os"
	"path/filepath"
	"runtime"
)

var DebugMode = true

// Fehler Count für die Sitzung
var ErrorCount = 1

func HandleError(e error, level string, description string) bool {
	if e == nil {
		return false // Kein Fehler
	}

	// Die Fehlerbeschreibung formatieren
	locationKey := FormatTerminal("Fehlerort:", "bright-black", false, false)
	errorKey := FormatTerminal("Fehlermeldung:", "bright-black", false, false)
	descriptionKey := FormatTerminal("Beschreibung:", "bright-black", false, false)

	// Die Fehlermeldung.
	finalMessage := fmt.Sprintf("%s %s", errorKey, e.Error())

	// Falls Beschreibung vorhanden, dann hinzufügen
	if description != "" {
		finalMessage = fmt.Sprintf("%s\n%s %s",finalMessage, descriptionKey, description)
	}

	// Kontext des Aufrufs herausfinden - nur verwenden wenn DebugMode true
	if DebugMode {
		pc, _, line, ok := runtime.Caller(2) // Zum Aufrufer müssen 2 Schirtte zuürckgegangen werden
		context := ""
		if ok {
			funcName := runtime.FuncForPC(pc).Name() // Funktionsnamen ermitteln
			context = fmt.Sprintf("%s Zeile %d in %s", locationKey, line, filepath.Base(funcName)) // Kontext Zusammensetzen
			finalMessage = fmt.Sprintf("%s\n%s\n", context, finalMessage)
		}
	}

	hyphen := "=========================================================================================================="

	switch strings.ToLower(level) {
		case "warning":
			hyphen = FormatTerminal(hyphen, "bright-yellow", true, false)
			prefix := FormatTerminal(fmt.Sprintf("%d. Warnung:", ErrorCount), "bright-yellow", true, false)
			fmt.Printf("\n%s\n%s\n\n", prefix, finalMessage)

		case "fatal":
			hyphen = FormatTerminal(hyphen, "bright-red", true, false)
			prefix := FormatTerminal(fmt.Sprintf("%d. Kritischer Fehler:", ErrorCount), "bright-red", true, false)
			fmt.Printf("\n%s\n%s\n\n", prefix, finalMessage)
			os.Exit(1) // Programm mit Fehlercode 1 beenden

		case "error", "default":
			hyphen = FormatTerminal(hyphen, "red", true, false)
			// Fallback für unbekannte Level
			prefix := FormatTerminal(fmt.Sprintf("%d. Fehler:", ErrorCount), "red", false, false)
			fmt.Printf("\n%s\n%s\n%s%s\n\n", hyphen, prefix, finalMessage, hyphen)
	}

	ErrorCount++
	return true
}
