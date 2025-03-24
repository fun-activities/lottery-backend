// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fun-activities/lottery-backend/config"
	"github.com/fun-activities/lottery-backend/internal/controller/http"
	"github.com/fun-activities/lottery-backend/internal/repo"
	"github.com/fun-activities/lottery-backend/internal/service"
	"github.com/fun-activities/lottery-backend/pkg/httpserver"
	"github.com/fun-activities/lottery-backend/pkg/logger"
	"github.com/fun-activities/lottery-backend/pkg/mysql"
	"github.com/gin-gonic/gin"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	dsn := "root:@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlClient, err := mysql.Open(dsn)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - mysql.Open %w", err))
	}
	defer mysqlClient.Close()

	// Use case
	service := service.New(
		service.Dependent{
			Repo: repo.New(mysqlClient),
		},
	)

	// HTTP Server
	handler := gin.New()
	http.NewRouter(handler, service, l)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
