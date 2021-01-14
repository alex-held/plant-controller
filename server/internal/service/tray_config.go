package service

import (
	"context"
	
	"github.com/pkg/errors"
	
	"plant-controller/internal/model"
	"plant-controller/internal/store"
)

// TrayWebService ...
type TrayConfigWebService struct {
	ctx   context.Context
	store *store.Store
}

// NewTrayWebService creates a new Tray web service
func NewTrayConfigWebService(ctx context.Context, store *store.Store) *TrayConfigWebService {
	return &TrayConfigWebService{
		ctx:   ctx,
		store: store,
	}
}

func (g TrayConfigWebService) Get(ctx context.Context, id int) (*model.TrayConfig, error) {
	result, err := g.store.TrayConfig.Get(ctx, int64(id))
	if err != nil {
		return nil, errors.Wrapf(err, "[svc.trayconfig.get] error getting trayconfigs from repo")
	}
	return result.ToApp(), err
}

func (g TrayConfigWebService) GetAll(ctx context.Context) (res []*model.TrayConfig, err error) {
	all, err := g.store.TrayConfig.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[svc.trayconfig.getall] error loading trayconfigs from repo")
	}
	for _, i := range all {
		res = append(res, i.ToApp())
	}
	return res, nil
}

func (g TrayConfigWebService) Create(ctx context.Context, config *model.TrayConfig) (*model.TrayConfig, error) {
	created, err := g.store.TrayConfig.Create(ctx, config.ToDB())
	if err != nil {
		return nil, errors.Wrapf(err, "[svc.trayconfig.create] error creating trayconfigs in repo")
	}
	return created.ToApp(), err
}

func (g TrayConfigWebService) Update(ctx context.Context, config *model.TrayConfig) (*model.TrayConfig, error) {
	updated, err := g.store.TrayConfig.Update(ctx, config.ToDB())
	if err != nil {
		return nil, errors.Wrapf(err, "[svc.trayconfig.update] error updating trayconfigs in repo")
	}
	return updated.ToApp(), err
}

func (g TrayConfigWebService) Delete(ctx context.Context, i int) error {
	panic("implement me")
}
