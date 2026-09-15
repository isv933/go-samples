package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/isv933/go-samples/hello-rest-api/gen/api"
)

type handler struct{}

func (handler) GetUserById(_ context.Context, params api.GetUserByIdParams) (*api.User, error) {
	// Пример данных. Здесь можно подключить хранилище пользователей.
	user := &api.User{Name: "User " + strconv.Itoa(params.ID)}
	if params.ID == 1 {
		user.Name = "Alice"
		user.SetEmail(api.NewOptString("alice@example.com"))
	}
	return user, nil
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":18080"
	}

	apiServer, err := api.NewServer(handler{})
	if err != nil {
		logger.Error("create API server", "error", err)
		os.Exit(1)
	}

	server := &http.Server{
		Addr:              addr,
		Handler:           apiServer,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("shutdown HTTP server", "error", err)
		}
	}()

	logger.Info("HTTP server started", "addr", addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("HTTP server stopped", "error", err)
		os.Exit(1)
	}
}
