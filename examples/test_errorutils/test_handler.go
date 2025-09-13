package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sadghsbaj/go-utils/terminal"
)

// Dein Struct, genau wie du es wolltest.
type AppError struct {
	Message      string
	Err          error
	Level        string
	FunctionName string
}

// Eine Methode, die `Error()` heißt und einen `string` zurückgibt.
func (e *AppError) Error() string {
	errMsg := e.Message
	err := e.Err
	errLevel := e.Level
	errFunction := e.FunctionName

	// Formatieren
	test, _ := terminal.FormatAlert(errMsg, errLevel, err)
	output := fmt.Sprintf("%s %s", test, errFunction)


	return output
}

func handler(err error) bool {
	if err == nil {
		return false
	}

	fmt.Println(err)
	return true
}

func main() {
	err := fail()
	if handler(err) {fmt.Println(err)}
}

// Deine Funktion, die den AppError erstellt und zurückgibt.
func fail() error {
	e := os.ErrNotExist // Irgendein Beispiel-Fehler

	// Du erstellst deinen struct und gibst ihn zurück.
	// Weil er die Error() Methode hat, darfst du das.
	return &AppError{
		Message:      "Das ist meine simple Fehlermeldung",
		Err:          e,
		Level:        "Warning",
		FunctionName: getFunctionName(),
	}
}

func getFunctionName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "unbekannt"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unbekannt"
	}
	return fn.Name()
}
