package web

import "fmt"
import "net/http"
import "github.com/sadghsbaj/go-utils/terminal"

func ServeFiles(url string, dir string) error {
	// Verzeichnis festlegen
	fs := http.FileServer(http.Dir(dir))

	// Dateien aus dem Verzeichnis nach abschneiden der url zur verfügung stellen
	http.Handle(url, http.StripPrefix(url, fs))

	infoPrefix, e := terminal.Format("Info:", "cyan", true, false)
	if e != nil {
		return fmt.Errorf("Fehler beim formatieren des infoPrefix: %w", e)
	}

	dir, e = terminal.Format(dir, "yellow", true, false)
	if e != nil {
		return fmt.Errorf("Fehler beim formatieren des directory: %w", e)
	}

	fmt.Printf("%s Dateien aus dem Ordner %s werden ausgeliefert.\n", infoPrefix, dir)
	return nil
}
