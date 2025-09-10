package main

import (
	"fmt"
	"github.com/sadghsbaj/go-utils/terminal"
)

func main() {
	// Gemeinsame Testdaten
	text := "Das ist der Terminal Test Text!"
	bold := true
	underline := true

	// Für erfolgreichen Test
	color := "blue"

	// Für fehlschlagenden Test
	colorFail := "sajapa" // Keine gültige Farbe

	passTest(text, color, bold, underline)
	failTest(text, colorFail, bold, underline)
}

// Dieser Test sollte bestehen.
// Der ausgegebene Text sollte Blau, fett und unterstrichen sein.
func passTest(t, c string, b, u bool) {
	textFormatted, _ := terminal.Format(t, c, b, u)

	fmt.Printf("\nTest Erfolgreich!\n%s\n\n", textFormatted)
}

// Dieser Test sollte einen Fehler aufwerfen.
// Es sollte eine Liste aller verfügbaren Farben ausgegeben werden.
func failTest(t, c string, b, u bool) {
	_, e := terminal.Format(t, c, b, u)

	fmt.Printf("Fehler!\n%s\n\n", e)
}
