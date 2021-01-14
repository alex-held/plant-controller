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
	trayController := controller.NewTrays(ctx, serviceManager, l)
	growConfigController := controller.NewGrowConfig(ctx, serviceManager, l)
	plantController := controller.NewPlant(ctx, serviceManager, l)
	trayConfigController := controller.NewTrayConfig(ctx, serviceManager, l)
	
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
	
//	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	
	// API V1
	v1 := e.Group("/v1")
	
	// Tray routes
	trayRoutes := v1.Group("/tray")
	trayRoutes.GET("/", trayController.GetAll)
	trayRoutes.POST("/", trayController.Create)
	trayRoutes.GET("/:id", trayController.Get)
	
	// GrowConfig routes
	growConfigRoutes := v1.Group("/growconfig")
	growConfigRoutes.GET("/", growConfigController.GetAll)
	growConfigRoutes.POST("/", growConfigController.Create)
	growConfigRoutes.GET("/:id", growConfigController.Get)
	
	// growConfigRoutes.DELETE("/:id", growConfigController.Delete)
	
	// GrowConfig routes
	trayConfigRoutes := v1.Group("/trayconfig")
	trayConfigRoutes.GET("/", trayConfigController.GetAll)
	trayConfigRoutes.POST("/", trayConfigController.Create)
	trayConfigRoutes.GET("/:id", trayConfigController.Get)
	
	// GrowConfig routes
	plantRoutes := v1.Group("/plant")
	plantRoutes.GET("/", plantController.GetAll)
	plantRoutes.POST("/", plantController.Create)
	plantRoutes.GET("/:id", plantController.Get)
	plantRoutes.PUT("/:id", plantController.Update)
	plantRoutes.DELETE("/:id", plantController.Delete)
	
	// TODO: userRoutes.DELETE("/:slot", trayController.Delete)
	// TODO: userRoutes.PUT("/:id", trayController.Update)
	
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
