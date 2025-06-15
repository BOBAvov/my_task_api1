package main

import (
	"fmt"
	"log/slog"
	"my_tastAPI/internal/config"
	"my_tastAPI/internal/models"
	"my_tastAPI/internal/transport"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	cfg := config.ConfigInit()

	log := setupLogger()

	mu := &sync.Mutex{}

	data := models.NewDataModel()

	router := transport.NewRouter(data, mu, log)

	log.Info(fmt.Sprintf("server is starting for: %s", cfg.Add))

	go http.ListenAndServe(cfg.Add, router)

	<-signalChan

	// задел на будущее под полноценный grace fullShutdown

	log.Info("Server stopped")

}
func setupLogger() *slog.Logger {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	return log
}
