package store

import (
	"context"
	
	"plant-controller/internal/model"
)

// UserRepo is a store for trays
//go:generate mockery --dir . --name UserRepo --output ./mocks
type TrayRepo interface {
	Repo
	Get(context.Context, int64) (*model.TrayDB, error)
	GetAll(context.Context) ([]*model.TrayDB, error)
	Create(context.Context, *model.TrayDB) (*model.TrayDB, error)
	Update(context.Context, *model.TrayDB) (*model.TrayDB, error)
	Delete(context.Context, int64) error
}

// GrowConfigRepo is a store for GrowConfig
type GrowConfigRepo interface {
	Repo
	Get(context.Context, int64) (*model.GrowConfigDB, error)
	GetAll(context.Context) ([]*model.GrowConfigDB, error)
	Create(context.Context, *model.GrowConfigDB) (*model.GrowConfigDB, error)
	Update(context.Context, *model.GrowConfigDB) (*model.GrowConfigDB, error)
	Delete(context.Context, int64) error
}

type Repo interface {
	GetNextId() int64
}

// GrowConfigRepo is a store for GrowConfig
type TrayConfigRepo interface {
	Repo
	Get(context.Context, int64) (*model.TrayConfigDB, error)
	GetAll(context.Context) ([]*model.TrayConfigDB, error)
	Create(context.Context, *model.TrayConfigDB) (*model.TrayConfigDB, error)
	Update(context.Context, *model.TrayConfigDB) (*model.TrayConfigDB, error)
	Delete(context.Context, int64) error
}

// GrowConfigRepo is a store for GrowConfig
type PlantRepo interface {
	Repo
	Get(context.Context, int64) (*model.PlantDB, error)
	GetAll(context.Context) ([]*model.PlantDB, error)
	Create(context.Context, string) (*model.PlantDB, error)
	Update(context.Context, *model.PlantDB) (*model.PlantDB, error)
	Delete(context.Context, int64) error
}
