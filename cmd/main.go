package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rusystem/cache"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/internal/repository"
	http_server "github.com/rusystem/web-api-gateway/internal/server/http"
	"github.com/rusystem/web-api-gateway/internal/service"
	http_handler "github.com/rusystem/web-api-gateway/internal/transport/http"
	"github.com/rusystem/web-api-gateway/pkg/auth"
	grpc "github.com/rusystem/web-api-gateway/pkg/client/grpc/accounts"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc/warehouse"
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
	// 91.243.71.100:8080

	// init configs
	cfg, err := config.New(false)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to initialize config, err: %v", err))
	}

	// init in-memory cache
	memCache := cache.New()
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to initialize cache, err: %v", err))
	}

	// init token manager
	tokenManager, err := auth.NewManager(cfg.Auth.SigningKey)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to initialize token manager, err: %v", err))
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
	splClient, err := warehouse.NewSuppliersClient(cfg.Url.Warehouse)
	if err != nil {
		logger.Fatal(fmt.Sprintf("can`t connect to chat grpc service, err - %v\n", err))
	}
	defer func(splClient *warehouse.SuppliersClient) {
		if err = splClient.Close(); err != nil {
			logger.Error(fmt.Sprintf("the error occurred while closing supplier grpc connection, err - %v", err))
		}
	}(splClient)

	// init grpc warehouse client
	whClient, err := warehouse.NewWarehouseClient(cfg.Url.Warehouse)
	if err != nil {
		logger.Fatal(fmt.Sprintf("can`t connect to chat grpc service, err - %v\n", err))
	}
	defer func(whClient *warehouse.WarehouseClient) {
		if err = whClient.Close(); err != nil {
			logger.Error(fmt.Sprintf("the error occurred while closing warehouse grpc connection, err - %v", err))
		}
	}(whClient)

	// init grpc company accounts client
	compClient, err := grpc.NewCompanyAccountsClient(cfg.Url.Accounts)
	if err != nil {
		logger.Fatal(fmt.Sprintf("can`t connect to accounts grpc service, err - %v\n", err))
	}
	defer func(cc *grpc.CompanyAccountsClient) {
		if err = cc.Close(); err != nil {
			logger.Error(fmt.Sprintf("the error occurred while closing accounts grpc connection, err - %v", err))
		}
	}(compClient)

	// init grpc user accounts client
	userClient, err := grpc.NewUserAccountsClient(cfg.Url.Accounts)
	if err != nil {
		logger.Fatal(fmt.Sprintf("can`t connect to accounts grpc service, err - %v\n", err))
	}
	defer func(uc *grpc.UserAccountsClient) {
		if err = uc.Close(); err != nil {
			logger.Error(fmt.Sprintf("the error occurred while closing accounts grpc connection, err - %v", err))
		}
	}(userClient)

	// init grpc sections accounts client
	sectionClient, err := grpc.NewSectionsAccountsClient(cfg.Url.Accounts)
	if err != nil {
		logger.Fatal(fmt.Sprintf("can`t connect to accounts grpc service, err - %v\n", err))
	}
	defer func(sc *grpc.SectionsAccountsClient) {
		if err = sc.Close(); err != nil {
			logger.Error(fmt.Sprintf("the error occurred while closing accounts grpc connection, err - %v", err))
		}
	}(sectionClient)

	// init dep-s
	repo := repository.New(cfg, memCache, pc)
	srv := service.New(service.Config{
		Config:          cfg,
		Repo:            repo,
		TokenManager:    tokenManager,
		SuppliersClient: splClient,
		WarehouseClient: whClient,
		CompanyClient:   compClient,
		UserClient:      userClient,
		SectionsClient:  sectionClient,
	})
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
