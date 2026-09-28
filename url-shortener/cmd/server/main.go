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
	"github.com/jackc/pgx/v5/pgxpool"
)

func readConfigSettings(fileName string) settings.Settings {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var s settings.Settings

	if err := json.Unmarshal(data, &s); err != nil {
		panic(err)
	}

	return s
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

	s := readSettings(*configFile)

	dbPool, err := pgxpool.New(context.Background(), s.DbConnectionString)
	if err != nil {
		panic(err)
	}

	defer dbPool.Close()
	shortenerService, err := NewShortenerService(PostgresRepository{pool: dbPool}, s.RedirectUrl)
	if err != nil {
		logger.Error("create shortener server", "error", err)
		os.Exit(1)
	}

	shortenerHandler, err := api.NewServer(shortenerApi{shortenerService})
	if err != nil {
		logger.Error("create API server", "error", err)
		os.Exit(1)
	}

	httpServer := &http.Server{
		Addr:    s.ListenAddress,
		Handler: shortenerHandler,
	}

	defer shutdown(func(ctx context.Context) {
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.Error("shutdown HTTP server", "error", err)
		}
	}, s.ServerTimeout.TimeDuration())()

	logger.Info("HTTP server started", "addr", s.ListenAddress)
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("HTTP server stopped", "error", err)
		os.Exit(1)
	}

}
