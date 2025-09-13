package errorutils

import (
	"fmt"
	"os"
	"strings"
	"github.com/sadghsbaj/go-utils/terminal"
)

type AppError struct {
	Message      string
	Err          error
	Level        string
	FunctionName string
}

func Handler(e error) bool {
	// Falls kein Fehler false zurückgeben
	if e == nil {
		return false
	}


}
