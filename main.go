package main

import "log"
import "plugin"

func main() {

	// Beispieltext
	text := "Hello world!"

	// komilierte Plugin-Datei öffnen
	pluginFile, err := plugin.Open("plugin.bin")
	if err != nil {
		log.Fatalf("Failed to open plugin.bin: %s", err.Error())
	}

	// Funktion Replace() in plugin.bin finden
	replaceFunc, err := pluginFile.Lookup("Replace")
	if err != nil {
		log.Fatalf("Failed to lookup Replace() in plugin.bin: %s", err.Error())
	}

	// Replace()-Funktion von interface{} zu func(string)string casten & ausführen
	text = replaceFunc.(func(string) string)(text)

	// Resultat anzeigen
	log.Printf("Text: '%s'", text)
}
