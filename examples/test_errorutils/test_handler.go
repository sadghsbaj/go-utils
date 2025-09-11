package main

import (
	"fmt"
	"errors"
	"github.com/sadghsbaj/go-utils/errorutils"
)

func main() {
	// Kein Fehler
	e := noError()
	if errorutils.Handler(e, errorutils.LevelError) {}

	// Warnung
	e = isWarning()
	if errorutils.Handler(e, errorutils.LevelWarning) {}

	// Error
	e = isError()
	if errorutils.Handler(e, errorutils.LevelError) {}

	// Fataler Fehler
	e = isFatalError()
	if errorutils.Handler(e, errorutils.LevelFatal) {}
}

func noError() error {
	fmt.Println(" -> Funktion 'noError' wird ausgeführt...")
	return nil
}

func isWarning() error {
	fmt.Println(" -> Funktion 'isWarning' wird ausgeführt...")
	return errors.New("Das ist eine Warnung.")
}

func isError() error {
	fmt.Println(" -> Funktion 'isError' wird ausgeführt...")
	return errors.New("Das ist ein Fehler.")
}

func isFatalError() error {
	fmt.Println(" -> Funktion 'isFatalError' wird ausgeführt...")
	return errors.New("Das ist ein Fataler Fehler.")
}
