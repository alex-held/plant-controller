package service

import (
	"context"
	
	"plant-controller/internal/model"
)

// UserService is a service for users
//go:generate mockery --dir . --name TrayService --output ./mocks
type TrayService interface {
	GetTray(context.Context, int) (*model.Tray, error)
	GetAllTrays(context.Context) ([]*model.Tray, error)
	CreateTray(context.Context, *model.Tray) (*model.Tray, error)
	UpdateTray(context.Context, *model.Tray) (*model.Tray, error)
	DeleteTray(context.Context, int) error
}

//go:generate mockery --dir . --name GrowConfigService --output ./mocks
type GrowConfigService interface {
	GetGrowConfig(context.Context, int) (*model.GrowConfig, error)
	GetAll(context.Context) ([]*model.GrowConfig, error)
	CreateGrowConfig(context.Context, *model.GrowConfig) (*model.GrowConfig, error)
	UpdateGrowConfig(context.Context, *model.GrowConfig) (*model.GrowConfig, error)
	DeleteGrowConfig(context.Context, int) error
}
