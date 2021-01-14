package service

import (
	"context"
	
	"github.com/pkg/errors"
	
	"plant-controller/internal/model"
	"plant-controller/internal/store"
)

// TrayWebService ...
type TrayWebService struct {
	ctx   context.Context
	store *store.Store
}

// NewTrayWebService creates a new Tray web service
func NewTrayWebService(ctx context.Context, store *store.Store) *TrayWebService {
	return &TrayWebService{
		ctx:   ctx,
		store: store,
	}
}

// GetTray ...
func (svc *TrayWebService) GetTray(ctx context.Context, id int) (*model.Tray, error) {
	tray, err := svc.store.Tray.Get(ctx, int64(id))
	if err != nil {
		return nil, errors.Wrap(err, "svc.tray.GetTray")
	}
	
	return tray.ToApp(), nil
}

// GetAllTrays returns all trays stored in the database
func (svc *TrayWebService) GetAllTrays(ctx context.Context) (res []*model.Tray, err error) {
	all, err := svc.store.Tray.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[svc.tray.getall] error loading tray from repo")
	}
	for _, i := range all {
		res = append(res, i.ToApp())
	}
	return res, nil
}

// CreateTray ...
func (svc TrayWebService) CreateTray(ctx context.Context, reqTray *model.Tray) (*model.Tray, error) {
	created, err := svc.store.Tray.Create(ctx, reqTray.ToDB())
	if err != nil {
		return nil, errors.Wrap(err, "[svc.tray.create] error while creating tray")
	}
	return created.ToApp(), nil
}

// UpdateTray ...
func (svc *TrayWebService) UpdateTray(ctx context.Context, reqTray *model.Tray) (*model.Tray, error) {
	updated, err := svc.store.Tray.Update(ctx, reqTray.ToDB())
	if err != nil {
		return nil, errors.Wrap(err, "[svc.tray.update] error while updating tray")
	}
	return updated.ToApp(), err
}

// DeleteTray ...
func (svc *TrayWebService) DeleteTray(ctx context.Context, id int) error {
	// Check if tray exists
	err := svc.store.Tray.Delete(ctx, int64(id))
	if err != nil {
		return errors.Wrap(err, "[svc.tray.delete] error")
	}
	return nil
}
