package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/internal/repository"
	http_server "github.com/rusystem/web-api-gateway/internal/server/http"
	"github.com/rusystem/web-api-gateway/internal/service"
	http_handler "github.com/rusystem/web-api-gateway/internal/transport/http"
	"github.com/rusystem/web-api-gateway/pkg/auth"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc"
	"github.com/rusystem/web-api-gateway/pkg/database"
	"github.com/rusystem/web-api-gateway/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// init logger
func init() {
	logger.ZapLoggerInit()
}

// @title Web api gateway API
// @version 1.0
// @description API gateway

// @contact.name ru.system.ru@gmail.com
// @contact.email ru.system.ru@gmail.com

// @host localhost:8080
// @BasePath /api/web-api-gateway/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// init configs
	cfg, err := config.New(false)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to initialize config, err: %v", err))
	}

	// init token manager
	tokenManager, err := auth.NewManager(cfg.Auth.SigningKey)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to initialize token manager, err: %v", err))
	}

	// init memcache
	mc, err := database.NewMemcache(cfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to initialize memcache, err: %v", err))
	}

	// init postgres connection
	pc, err := database.NewPostgresConnection(database.PostgresConfig{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	defer func(pc *sql.DB) {
		if err = pc.Close(); err != nil {
			logger.Error(fmt.Sprintf("postgres: failed to close connection, err: %v", err.Error()))
		}
	}(pc)

	// init grpc supplier client
	splClient, err := grpc.NewSuppliersClient(cfg.Url.Warehouse)
	if err != nil {
		logger.Fatal(fmt.Sprintf("can`t connect to chat grpc service, err - %v\n", err))
	}
	defer func(splClient *grpc.SuppliersClient) {
		if err = splClient.Close(); err != nil {
			logger.Error(fmt.Sprintf("the error occurred while closing supplier grpc connection, err - %v", err))
		}
	}(splClient)

	whClient, err := grpc.NewWarehouseClient(cfg.Url.Warehouse)
	if err != nil {
		logger.Fatal(fmt.Sprintf("can`t connect to chat grpc service, err - %v\n", err))
	}
	defer func(whClient *grpc.WarehouseClient) {
		if err = whClient.Close(); err != nil {
			logger.Error(fmt.Sprintf("the error occurred while closing warehouse grpc connection, err - %v", err))
		}
	}(whClient)

	// init dep-s
	repo := repository.New(cfg, mc, pc)
	srv := service.New(cfg, repo, tokenManager, splClient, whClient)
	hh := http_handler.NewHandler(srv, tokenManager, cfg)

	// HTTP Server
	server := http_server.New(cfg, hh.Init())

	go func() {
		if err = server.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(fmt.Sprintf("error occurred while running http server: %s", err))
		}
	}()

	logger.Info("server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err = server.Stop(ctx); err != nil {
		logger.Error(fmt.Sprintf("failed to stop server: %v", err))
	}
}
