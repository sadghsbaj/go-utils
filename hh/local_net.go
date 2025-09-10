package hh

import "fmt"
import "net"
import "errors"
import "strings"
import "ws-project/utils/ch"

func GetLocalIp() (string, error) {
	adresses, e := net.InterfaceAddrs()
	if e != nil {
		return "", fmt.Errorf("Netzwerkschnittstellen konnten nicht abgerufen werden: %w", e)
	}

	for _, address := range adresses {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("Keine passende IPv4-Adresse gefunden")
}

func PrintServerUrl(port string) error {
	if port == "" {
		return errors.New("Es wurde kein Port übergeben.")
	}

	// Port normalisieren
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	localIP, e := GetLocalIp()
	if e != nil {
		return fmt.Errorf("Fehler beim abrufen der lokalen IP Adresse: %w", e)
	}

	url := fmt.Sprintf("http://%s%s", localIP, port)
	url = ch.FormatTerminal(url, "blue", false, true)
	fmt.Printf("\nWebserver erreichbar unter: %s\n\n", url)

	return nil
}
