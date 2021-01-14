package service

import (
	"context"
	
	"github.com/pkg/errors"
	
	"plant-controller/internal/model"
	"plant-controller/internal/store"
)

// TrayWebService ...
type GrowConfigWebService struct {
	ctx   context.Context
	store *store.Store
}

func (g GrowConfigWebService) GetGrowConfig(ctx context.Context, i int) (*model.GrowConfig, error) {
	panic("implement me")
}

func (g GrowConfigWebService) GetAll(ctx context.Context) (res []*model.GrowConfig, err error) {
	all, err := g.store.GrowConfig.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[svc.growconfig.getall] error loading growconfigs from repo")
	}
	for _, i := range all {
		res = append(res, i.ToApp())
	}
	return res, nil
}

func (g GrowConfigWebService) CreateGrowConfig(ctx context.Context, config *model.GrowConfig) (*model.GrowConfig, error) {
	panic("implement me")
}

func (g GrowConfigWebService) UpdateGrowConfig(ctx context.Context, config *model.GrowConfig) (*model.GrowConfig, error) {
	panic("implement me")
}

func (g GrowConfigWebService) DeleteGrowConfig(ctx context.Context, i int) error {
	panic("implement me")
}

// NewTrayWebService creates a new Tray web service
func NewGrowConfigWebService(ctx context.Context, store *store.Store) *GrowConfigWebService {
	return &GrowConfigWebService{
		ctx:   ctx,
		store: store,
	}
}
