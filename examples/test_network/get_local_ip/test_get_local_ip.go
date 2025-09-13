package main

import (
	"fmt"
	"github.com/sadghsbaj/go-utils/errorutils"
	"github.com/sadghsbaj/go-utils/network"
)

func main() {
	// Erwartetes Ergebnis
	expected := "192.168.178.97"

	// GetLocalIP Funktion zum Test aufrufen
	localIP, e := network.GetLocalIP()
	if errorutils.Handler(e, "error") {return}

	// Prüfen, ob erwartetes Ergebnis gleich tatsächlichem Ergebnis
	result := expected == localIP
	if result {
		fmt.Printf("IP Adresse: %s - Test Erfolgreich.\n", localIP)
		return
	}
	fmt.Printf("IP Adresse: %s - Test fehlgeschlagen.\n", localIP)
}
