package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/ninja-dark/echelon_task/internal/command"

	"github.com/ninja-dark/echelon_task/internal/infrastructure/api/handler"
	"github.com/ninja-dark/echelon_task/internal/infrastructure/api/router"
	"github.com/ninja-dark/echelon_task/internal/infrastructure/api/server"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = logger.Sync() }()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	//cfg, err := config.Read()
	//if err != nil {
	//	logger.Sugar().Fatalf("Cannot load config, due to error: %s", err.Error())
	//}

	l := command.NewExecutor()
	hs := handler.NewHandler(l)
	rt := router.NewRouter(hs)
	logger.Sugar().Infof("Starting Gateway server on port:%s", ":8080")
	srv := server.NewServer(":8080", rt, logger)

	if err := srv.Start(ctx); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Sugar().Fatalf("Failed to start server: %s", err.Error())
		}
	}
}
