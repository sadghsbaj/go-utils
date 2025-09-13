package errorutils

import (
	"fmt"
	"strings"
	"os"
	"github.com/sadghsbaj/go-utils/terminal"
)

// Konstanten für Fehlerlevel
const (
	LevelWarning 	= "warning"
	LevelError 	= "error"
	LevelFatal 	= "fatal"
)

// Handler prüft, ob ein Fehler vorliegt, und gibt ihn formatiert auf Stderr aus.
// Er nimmt einen Fehler, einen Level ("warning", "error", "fatal")
// und gibt 'true' zurück, wenn ein Fehler behandelt wurde.
func HandlerOld(e error, level string) bool {
	// Falls kein Fehler false zurückgeben
	if e == nil {
		return false
	}

	// Je nach FehlerLevel an die printError Funktion übergeben
	var label string
	var color string

	switch strings.ToLower(level) {
		case LevelWarning:
			label = "[WARNING]"
			color = "yellow"

		case LevelError:
			label = "[ERROR]"
			color = "red"

		case LevelFatal:
			label = "[FATAL-ERROR]"
			color = "bright-red"

		// Wird nur erreicht wenn ein ungültiger Level angegeben wurde
		default:
			validLevels := []string{LevelWarning, LevelError, LevelFatal}

			// Fehler ausgeben dass ungültiger Level verwendet wurde und verfügbare Level auflisten
			fmt.Fprintf(
				os.Stderr,
				"[ERROR-HANDLER-FEHLER] Ungültiger Level '%s'. Bitte einen der folgenden Level verwenden: %s\n",
				level,
				strings.Join(validLevels, ", "),
			)

			// Ursprünglicher Fehler soll trz. nicht verloren geben, deshalb mit label "Undefined" und weißer Schrift ausgeben
			label = fmt.Sprintf("%s-UNDEFINED", strings.ToUpper(level))
			color = "white"
	}

	// Funktion zur Ausgabe im Terminal aufrufen
	printErrorOld(e, label, color)

	// Wird nur erreicht wenn es einen Fehler gab daher true zurückgeben
	return true
}

func printErrorOld(e error, label, color string) {
	// Level Anzeige formatieren
	label, err := terminal.Format(label, color, true, false)
	if err != nil {
		// Wenn Formatierung fehlschlägt direkten Fehler ausgeben, Label wird unformatiert ausgegeben
		fmt.Fprintf(
			os.Stderr,
			"[ERROR-HANDLER-WARN] Konnte Fehler-Label nicht formatieren: %v}n",
			err,
		)
	}

	// Ausgabe zusammensetzen
	output := fmt.Sprintf("%s %s\n", label, e)

	// Ausgabe
	fmt.Fprintln(os.Stderr, output)
}
