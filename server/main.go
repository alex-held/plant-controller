package main

import (
	"context"
	"log"
	"net/http"
	"time"
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	
	"plant-controller/internal/config"
	"plant-controller/internal/controller"
	"plant-controller/internal/logger"
	"plant-controller/internal/service"
	"plant-controller/internal/store"
	"plant-controller/internal/validator"
	
	echoLog "github.com/labstack/gommon/log"
	
	libError "plant-controller/internal/error"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	
	// config
	cfg := config.Get()
	
	// logger
	l := logger.Get()
	
	// Init repository store (with mysql/postgresql inside)
	store, err := store.New(ctx)
	if err != nil {
		return errors.Wrap(err, "store.New failed")
	}
	
	// Init service manager
	serviceManager, err := service.NewManager(ctx, store)
	if err != nil {
		return errors.Wrap(err, "manager.New failed")
	}
	
	// Init controllers
	userController := controller.NewTrays(ctx, serviceManager, l)
	
	// Initialize Echo instance
	e := echo.New()
	e.Validator = validator.NewValidator()
	e.HTTPErrorHandler = libError.Error
	
	// Disable Echo JSON logger in debug mode
	if cfg.LogLevel == "debug" {
		if l, ok := e.Logger.(*echoLog.Logger); ok {
			l.SetHeader("${time_rfc3339} | ${level} | ${short_file}:${line}")
		}
	}
	
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// API V1
	v1 := e.Group("/v1")
	
	// User routes
	userRoutes := v1.Group("/tray")
	userRoutes.GET("/", userController.GetAll)
	userRoutes.POST("/", userController.Create)
	userRoutes.GET("/:slot", userController.Get)
	// TODO: userRoutes.DELETE("/:slot", userController.Delete)
	// TODO: userRoutes.PUT("/:id", userController.Update)
/*
	// File routes
	fileRoutes := v1.Group("/file")
	fileRoutes.POST("/", fileController.Create)
	fileRoutes.GET("/:id", fileController.Get)
	fileRoutes.DELETE("/:id", fileController.Delete)
	fileRoutes.PUT("/:id/content", fileController.Upload)
	fileRoutes.GET("/:id/content", fileController.Download)*/
	
	// Start server
	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
	
	return nil
}
