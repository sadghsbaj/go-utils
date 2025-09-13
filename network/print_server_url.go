package network

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
	"github.com/sadghsbaj/go-utils/terminal"
)

// PrintServerUrl validiert einen Port-String, kombiniert ihn mit der lokalen
// Netzwerk-IP und gibt die resultierende, klickbare Server-URL formatiert
// auf der Konsole aus.
func PrintServerUrl(port string) error {
	// Port validieren
	port, e := validatePort(port)
	if e != nil {
		return fmt.Errorf("Fehler beim validieren des Ports: %w", e)
	}

	// Lokale IP deklarieren
	localIP, e := GetLocalIP()
	if e != nil {
		return fmt.Errorf("Fehler beim abrufen der lokalen IP Adresse: %w", e)
	}

	// Url formatieren
	url, e := formatServerUrl(localIP, port)
	if e != nil {
		return fmt.Errorf("Fehler beim formatieren der Server Url: %w", e)
	}

	// Server Url ausgeben
	fmt.Printf("\nWebserver erreichbar unter: %s\n\n", url)

	return nil
}

func validatePort(port string) (string, error) {
	// Mögliche Leerzeichen im übergebenen Port entfernen
	port = strings.ReplaceAll(port, " ", "")

	// Prüfen, ob Port ein leerer String ist
	if port == "" {
		return "", errors.New("Es wurde kein Port übergeben.")
	}

	// Portnummer extrahieren (ohne ':')
	portNumStr := strings.TrimPrefix(port, ":")

	// Prüfen, ob der übergebene Port eine Zahl ist
	portNum, e := strconv.Atoi(portNumStr)
	if e != nil {
		return "", errors.New("Der Port muss eine Zahl sein.")
	}

	// Prüfen, ob die Zahl im gültigen Bereich liegt (gültig sind 1 inklusive bis 65535 inklusive)
	if portNum < 1 || portNum > 65535 {
		return "", fmt.Errorf("Die Portnummer %d liegt außerhalb des gültigen Bereiches (1-65535).", portNum)
	}

	return port, nil
}

func formatServerUrl(localIP, port string) (string, error) {
	// ':' zum Port hinzufügen falls nicht mit übergeben wurde
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	// Server Url zusammensetzen
	url := fmt.Sprintf("http://%s%s", localIP, port)

	// Server Url färben und unterstreichen
	url, e := terminal.Format(url, "blue", false, true)
	if e != nil {
		return "", fmt.Errorf("Fehler bei der Formatierung der Server Url: %w", e)
	}

	return url, nil
}
