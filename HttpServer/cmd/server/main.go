package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"httpserver/internal/config"
	"httpserver/internal/handlers"
	"httpserver/internal/middleware"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.Health) // no auth — health checks shouldn't need a key
	mux.Handle("POST /webhook", middleware.Chain(
		http.HandlerFunc(handlers.Webhook),
		middleware.APIKeyAuth(cfg.APIKey),
	))

	// Recovery listed first so it wraps EVERYTHING, including Logging —
	// a panic anywhere still gets logged and converted to a 500.
	handler := middleware.Chain(mux,
		middleware.Recovery(logger),
		middleware.Logging(logger),
	)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler,
	}

	// wire real OS signals into a context — this is the actual production
	// pattern from Context/graceful_shutdown.go, now driving a real server.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		logger.Info("server starting", "port", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done() // blocks here until SIGINT/SIGTERM arrives
	logger.Info("shutdown signal received, draining in-flight requests...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("forced shutdown after timeout", "error", err)
	} else {
		logger.Info("server stopped cleanly")
	}
}
