package main

import (
	"fmt"
	"os"
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
	failTest(colorFail, bold, underline)
	testAlertFormat()
}


// Dieser Test sollte bestehen.
// Der ausgegebene Text sollte Blau, fett und unterstrichen sein.
func passTest(t, c string, b, u bool) {
	textFormatted, _ := terminal.Format(t, c, b, u)

	fmt.Printf("\nTest Erfolgreich!\n%s\n\n", textFormatted)
}

// Dieser Test sollte einen Fehler aufwerfen.
// Es sollte eine Liste aller verfügbaren Farben ausgegeben werden.
func failTest(c string, b, u bool) {
	// 1. Falsche Farbe
	t := "Dieser Test wird fehlschlagen, da eine ungültige Farbe verwendet wird."
	_, e := terminal.Format(t, c, b, u)
	if e != nil {
		formattedMsg, _ := terminal.FormatAlert(t, "error", e)
		fmt.Println(formattedMsg)
	}

	// 2. Falsches Level
	t = "Dieser Test wird fehlschlagen, da ein ungültiges Level verwendet wird. Der Fehler wird trz. ausgegeben."
	level := "Hinweis"

	// Fehler simulieren
	_, e = os.ReadFile("/does/not/exist.txt")
	if e != nil {
		formattedMsg, _ := terminal.FormatAlert(t, level, e)
		fmt.Println(formattedMsg)
	}

}

func testAlertFormat() {
	infoText := "Das ist eine Info."
	warningText := "Das ist eine Warnung."
	errorText := "Das ist eine Fehlermeldung."
	fatalText := "Das ist eine kritische Fehlermeldung."

	// Fehler simulieren
	_, e := os.ReadFile("/does/not/exist.txt")

	// Formatieren
	formattedInfo, _ := terminal.FormatAlert(infoText, terminal.Info, nil)
	formattedWarning, _ := terminal.FormatAlert(warningText, terminal.Warning, e)
	formattedError, _ := terminal.FormatAlert(errorText, terminal.Error, e)
	formattedFatal, _ := terminal.FormatAlert(fatalText, terminal.Fatal, e)

	fmt.Println(formattedInfo)
	fmt.Println(formattedWarning)
	fmt.Println(formattedError)
	fmt.Println(formattedFatal)
}
