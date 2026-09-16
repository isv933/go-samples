package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"
)

type shutdownOperation func(context.Context)

func shutdown(shutdownFunc shutdownOperation, timeout time.Duration) func() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-ctx.Done()
		defer stop()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		shutdownFunc(shutdownCtx)
	}()

	return stop
}
