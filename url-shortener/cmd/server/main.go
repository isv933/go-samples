package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/isv933/go-samples/url-shortener/cmd/server/settings"
	"github.com/isv933/go-samples/url-shortener/gen/api"
)

func readConfigSettings(fileName string) settings.Settings {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var settings settings.Settings

	if err := json.Unmarshal(data, &settings); err != nil {
		panic(err)
	}

	return settings
}

func readSettings(configFile string) settings.Settings {
	if len(configFile) > 0 {
		return readConfigSettings(configFile)
	} else {
		return settings.NewSettings()
	}
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	configFile := flag.String("config-file", "", "Конфигурационный файл с настройками")
	flag.Parse()

	settings := readSettings(*configFile)

	shortenerHandler, err := api.NewServer(shortenerApi{})
	if err != nil {
		logger.Error("create API server", "error", err)
		os.Exit(1)
	}

	shortenerServer := &http.Server{
		Addr:    settings.ListenAddress,
		Handler: shortenerHandler,
	}

	defer shutdown(func(ctx context.Context) {
		if err := shortenerServer.Shutdown(ctx); err != nil {
			logger.Error("shutdown HTTP server", "error", err)
		}
	}, settings.ServerTimeout.TimeDuration())()

	logger.Info("HTTP server started", "addr", settings.ListenAddress)
	if err := shortenerServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("HTTP server stopped", "error", err)
		os.Exit(1)
	}

}
