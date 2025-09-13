package network

import (
	"fmt"
	"net"
	"errors"
)

// GetLocalIP durchsucht die Netzwerkschnittstellen des Systems und gibt
// die erste gefundene, nicht-loopback IPv4-Adresse als String zurück.
func GetLocalIP() (string, error) {
	// Alle konfigurierten Netzwerkadressen des Computers
	adresses, e := net.InterfaceAddrs()
	if e != nil {
		return "", fmt.Errorf("Netzwerkschnittstellen konnten nicht abgerufen werden: %w", e)
	}

	// Jede Adresse durchgehen
	for _, address := range adresses {
		// Prüfen ob Adresse vom Typ *net.IPNet ist
		// und ob die Adresse KEINE "Loopback-Adresse" ist, d.h. KEINE 'lokale' Computer Adresse
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// Versuchen die Adresse in IPv4 umzwuandeln, Bedingung wird nur Wahr wenn es sich um eine IPv4 Adresse handelt
			if ipnet.IP.To4() != nil {
				// Adresse in lesbaren Text umwandeln und übergeben
				return ipnet.IP.String(), nil
			}
		}
	}
	// Fehler zurückgeben falls keine IPv4 Adresse gefunden wird
	return "", errors.New("Keine passende IPv4-Adresse gefunden")
}
