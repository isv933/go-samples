package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/isv933/go-samples/url-shortener/cmd/server/settings"
)

func createSettings(fileName string) {
	file, _ := os.Create(fileName)
	defer file.Close()
	file.Write(
		func() []byte {
			data, err := json.MarshalIndent(settings.NewSettings(), "", "   ")
			if err != nil {
				panic(err)
			}

			return append(data, '\n')
		}())
}

func main() {
	settingsFile := flag.String("settings-file-name", "./url-shortener.json", "Имя файла настроек для генерации")
	flag.Parse()
	createSettings(*settingsFile)
}
