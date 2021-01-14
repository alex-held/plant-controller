package store

import (
	"context"
	
	"plant-controller/internal/model"
)

// UserRepo is a store for trays
//go:generate mockery --dir . --name UserRepo --output ./mocks
type TrayRepo interface {
	GetTray(context.Context, int) (*model.Tray, error)
	GetAllTrays(context.Context) ([]*model.Tray, error)
	CreateTray(context.Context, *model.Tray) (*model.Tray, error)
	UpdateTray(context.Context, *model.Tray) (*model.Tray, error)
	DeleteTray(context.Context, int) error
}

// GrowConfigRepo is a store for GrowConfig
type GrowConfigRepo interface {
	Get(context.Context, int64) (*model.GrowConfigDB, error)
	GetAll(context.Context) ([]*model.GrowConfigDB, error)
	Create(context.Context, *model.GrowConfigDB) (*model.GrowConfigDB, error)
	Update(context.Context, *model.GrowConfigDB) (*model.GrowConfigDB, error)
	Delete(context.Context, int64) error
}

/*
// FileMetaRepo is a store for files
//go:generate mockery --dir . --name FileMetaRepo --output ./mocks
type FileMetaRepo interface {
	GetFileMeta(context.Context, uuid.UUID) (*model.DBFile, error)
	CreateFileMeta(context.Context, *model.DBFile) (*model.DBFile, error)
	UpdateFileMeta(context.Context, *model.DBFile) (*model.DBFile, error)
	DeleteFileMeta(context.Context, uuid.UUID) error
}

// FileContentRepo is a store for file content
//go:generate mockery --dir . --name FileContentRepo --output ./mocks
type FileContentRepo interface {
	Upload(context.Context, *model.DBFile, []byte) error
	Download(context.Context, *model.DBFile) ([]byte, error)
}
*/
