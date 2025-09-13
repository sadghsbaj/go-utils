package web

import "fmt"
import "net/http"
import "strings"
import "os"

func RenderHtml(w http.ResponseWriter, filename string) error {
	// Übergebenen Dateinamen an allen "." aufteilen um das angebene Format zu erlangen
	parts := strings.Split(filename, ".")
	format := parts[len(parts) - 1]

	// Falls kein HTML, Fehlermeldung und abbrechen
	if format != "html" {
		return fmt.Errorf("Die Funktion 'RenderHtml' erwartet den Dateityp '.html'. Übergebener Dateityp: '.%s'\n", format)
	}

	// Html Datei einlesen
	html, e := os.ReadFile(filename)
	if e != nil {
		return fmt.Errorf("Konnte die Datei '%s' nicht lesen. %w", filename, e)
	}

	// Content Type setzen
	w.Header().Set("Content-Type", "Text/Html; charset=utf-8")

	// Datei servieren
	w.Write(html)

	return nil
}
