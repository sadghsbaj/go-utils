package hh

import "fmt"
import "net/http"
import "ws-project/utils/ch"

func ServeFiles(url string, dir string) {
	// Verzeichnis festlegen
	fs := http.FileServer(http.Dir(dir))

	// Dateien aus dem Verzeichnis nach abschneiden der url zur verfügung stellen
	http.Handle(url, http.StripPrefix(url, fs))

	infoPrefix := ch.FormatTerminal("Info:", "cyan", true, false)
	dir = ch.FormatTerminal(dir, "yellow", true, false)
	fmt.Printf("%s Dateien aus dem Ordner %s werden ausgeliefert.\n", infoPrefix, dir)
}
