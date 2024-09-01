package app

import (
	"annotator-backend/config"
	"context"
	"database/sql"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	MaxHeaderBytes = 1 << 20
	CtxTimeout     = 5
)

type AnnotatorApp struct {
	echo   *echo.Echo
	config *config.Config
	db     *sql.DB
}

func NewAnnotatorApp(config *config.Config, db *sql.DB) *AnnotatorApp {
	return &AnnotatorApp{echo: echo.New(), config: config, db: db}
}

func (annotatorApp *AnnotatorApp) Run() error {
	server := &http.Server{
		Addr:           annotatorApp.config.Server.Port,
		ReadTimeout:    time.Second * annotatorApp.config.Server.ReadTimeout,
		WriteTimeout:   time.Second * annotatorApp.config.Server.WriteTimeout,
		MaxHeaderBytes: MaxHeaderBytes,
	}

	go func() {
		log.Infof("Server is listening on PORT: %s", annotatorApp.config.Server.Port)
		if err := annotatorApp.echo.StartServer(server); err != nil {
			log.Fatal("Error starting Server: ", err)
		}
	}()

	if err := annotatorApp.MapHandlers(annotatorApp.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), CtxTimeout*time.Second)
	defer shutdown()

	log.Info("Server Exited Properly")
	return annotatorApp.echo.Server.Shutdown(ctx)
}
