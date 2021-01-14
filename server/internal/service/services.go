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
	Get(context.Context, int) (*model.GrowConfig, error)
	GetAll(context.Context) ([]*model.GrowConfig, error)
	Create(context.Context, *model.GrowConfig) (*model.GrowConfig, error)
	Update(context.Context, *model.GrowConfig) (*model.GrowConfig, error)
	Delete(context.Context, int) error
}

//go:generate mockery --dir . --name TrayConfigService --output ./mocks
type TrayConfigService interface {
	Get(context.Context, int) (*model.TrayConfig, error)
	GetAll(context.Context) ([]*model.TrayConfig, error)
	Create(context.Context, *model.TrayConfig) (*model.TrayConfig, error)
	Update(context.Context, *model.TrayConfig) (*model.TrayConfig, error)
	Delete(context.Context, int) error
}

//go:generate mockery --dir . --name PlantService --output ./mocks
type PlantService interface {
	Get(context.Context, int) (*model.Plant, error)
	GetAll(context.Context) ([]*model.Plant, error)
	Create(context.Context, string) (*model.Plant, error)
	Update(context.Context, *model.Plant) (*model.Plant, error)
	Delete(context.Context, int) error
}
