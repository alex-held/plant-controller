package service

import (
	"context"
	
	"github.com/pkg/errors"
	
	"plant-controller/internal/model"
	"plant-controller/internal/store"
)

// TrayWebService ...
type PlantWebService struct {
	ctx context.Context
	store *store.Store
}

// NewTrayWebService creates a new Tray web service
func NewPlantWebService(ctx context.Context, store *store.Store) *PlantWebService {
	return &PlantWebService{
		ctx:   ctx,
		store: store,
	}
}

func (g PlantWebService) Get(ctx context.Context, id int) (*model.Plant, error) {
	result, err := g.store.Plant.Get(ctx, int64(id))
	if err != nil {
		return nil, errors.Wrapf(err, "[svc.plant.get] error getting plants from repo")
	}
	return result.ToApp(), err
}

func (g PlantWebService) GetAll(ctx context.Context) (res []*model.Plant, err error) {
	all, err := g.store.Plant.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[svc.plant.getall] error loading plants from repo")
	}
	for _, i := range all {
		res = append(res, i.ToApp())
	}
	return res, nil
}

func (g PlantWebService) Create(ctx context.Context, name string) (*model.Plant, error) {
	created, err := g.store.Plant.Create(ctx, name)
	if err != nil {
		return nil, errors.Wrapf(err, "[svc.plant.create] error creating plants in repo")
	}
	return created.ToApp(), err
}

func (g PlantWebService) Update(ctx context.Context, config *model.Plant) (*model.Plant, error) {
	updated, err := g.store.Plant.Update(ctx, config.ToDB())
	if err != nil {
		return nil, errors.Wrapf(err, "[svc.plant.update] error updating plants in repo")
	}
	return updated.ToApp(), err
}

func (g PlantWebService) Delete(ctx context.Context, i int) error {
	panic("implement me")
}
