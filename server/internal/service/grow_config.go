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

func (g GrowConfigWebService) Get(ctx context.Context, i int) (*model.GrowConfig, error) {
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

func (g GrowConfigWebService) Create(ctx context.Context, config *model.GrowConfig) (*model.GrowConfig, error) {
	created, err := g.store.GrowConfig.Create(ctx, config.ToDB())
	if err != nil {
		return nil, errors.Wrapf(err, "[svc.growconfig.create] error creating plants in repo")
	}
	return created.ToApp(), err
}

func (g GrowConfigWebService) Update(ctx context.Context, config *model.GrowConfig) (*model.GrowConfig, error) {
	updated, err := g.store.GrowConfig.Update(ctx, config.ToDB())
	if err != nil {
		return nil, err
	}
	return updated.ToApp(), nil
}

func (g GrowConfigWebService) Delete(ctx context.Context, id int) error {
	return g.store.GrowConfig.Delete(ctx, int64(id))
}

// NewTrayWebService creates a new Tray web service
func NewGrowConfigWebService(ctx context.Context, store *store.Store) *GrowConfigWebService {
	return &GrowConfigWebService{
		ctx:   ctx,
		store: store,
	}
}
