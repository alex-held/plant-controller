package service

import (
	"context"
	
	"plant-controller/model"
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
