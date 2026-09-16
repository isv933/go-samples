package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/isv933/go-samples/url-shortener/gen/api"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	listenAddress := flag.String("listen-address", ":18080", "адрес для прослушивания")
	serverTimeoutValue := flag.String("server-timeout", "PT5S", "таймаут graceful shutdown в формате ISO 8601")
	flag.Parse()

	serverTimeout, err := parseISO8601Duration(*serverTimeoutValue)
	if err != nil {
		logger.Error("invalid server timeout", "value", *serverTimeoutValue, "error", err)
		os.Exit(2)
	}

	shortenerHandler, err := api.NewServer(shortenerApi{})
	if err != nil {
		logger.Error("create shortener API server", "error", err)
		os.Exit(1)
	}

	shortenerServer := &http.Server{
		Addr:    *listenAddress,
		Handler: shortenerHandler,
	}

	defer shutdown(func(ctx context.Context) {
		if err := shortenerServer.Shutdown(ctx); err != nil {
			logger.Error("shutdown HTTP server", "error", err)
		}
	}, serverTimeout)()

	logger.Info("HTTP server started", "addr", *listenAddress)
	if err := shortenerServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("HTTP server stopped", "error", err)
		os.Exit(1)
	}

}
