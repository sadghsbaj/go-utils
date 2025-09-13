package errorutils

import (
	"log"
	"os"
	"errors"
	"fmt"
	"path/filepath"
	"github.com/pelletier/go-toml"
)

type AppConfig struct {
	Mode string `toml:"mode"`
	Log LogConfig `toml:"log"`
}

type LogConfig struct {
	FilePath string `toml:"file_path"`
}

func LoadConfig() *AppConfig {
	tomlFilePath := "./app.toml"

	// Prüfen der Toml Config Datei
	e := valTomlFile(tomlFilePath)
	if e != nil {
		log.Fatalf("[FATAL-ERROR] %v", e)
	}

	// Toml Config einlesen
	data, e := os.ReadFile(tomlFilePath)
	if e != nil {
		log.Fatalf("[FATAL-ERROR] Einlesen der App Config nicht möglich. %v", e)
	}

	// Config Struct mit Toml Werten befüllen
	var conf AppConfig
	e = toml.Unmarshal(data, &conf)
	if e != nil {
		log.Fatalf("[FATAL-ERROR] Verwenden der App Config Werte nicht möglich. %v", e)
	}

	// Prüfen ob Werte gültig sind
	e = valAppConfigValues(conf)
	if e != nil {
		log.Fatalf("%v", e)
	}

	return &conf
}

func valTomlFile(filePath string) error {
	_, e := os.Stat(filePath)
	if e != nil {
		// Prüfen ob Datei nicht gefunden und ggf. behandeln
		if errors.Is(e, os.ErrNotExist) {
			fmt.Println("[ERROR] Die App Config Datei wurde nicht gefunden.")
			fmt.Println("[INFO] Versuche Die App Config Datei auf Root Ebene anzulegen.")
			e := createAppConfig(filePath)
			if e != nil {
				return fmt.Errorf("Die App Config Datei konnte nicht angelegt werden. %w", e)
			}

			return nil
		}

		// Sonstige Fehler im Umgang mit der Datei
		return e
	}

	return nil
}

func valAppConfigValues(conf AppConfig) error {
	if conf.Mode != "development" && conf.Mode != "production" {
		return fmt.Errorf("[FATAL-ERROR] Ungültiger Mode in der App Config. Mode: %s. Erlaubt: 'development' oder 'production'\n", conf.Mode)
	}
	if filepath.Ext(conf.Log.FilePath) != ".jsonl" {
		return fmt.Errorf("[FATAL-ERROR] Ungültiger Dateityp in der App Config. Dateityp: %s. Erlaubt: '.jsonl'\n", conf.Log.FilePath)
	}

	return nil
}

func createAppConfig(filePath string) error {
	// Toml Standard Werte festlegen
	conf := AppConfig{
		Mode: "development",
		Log: LogConfig{
			FilePath: "./logs.jsonl",
		},
	}

	// Content in Toml Format umwandeln
	content, e := toml.Marshal(conf)
	if e != nil {
		return fmt.Errorf("Die App Config konnte nicht umgewandelt werden. %w", e)
	}

	// Toml Datei erstellen
	e = os.WriteFile(filePath, content, 0644)
	if e != nil {
		return fmt.Errorf("Die App Config konnte nicht erstellt werden. %w", e)
	}

	fmt.Println("[INFO] Die App Config Datei wurde erfolgreich mit Standardwerten erstellt.")

	return nil
}
