package service

import (
	"context"
	"fmt"
	
	"github.com/pkg/errors"
	
	"plant-controller/lib/types"
	"plant-controller/model"
	"plant-controller/store"
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
func (svc *TrayWebService) GetTray(ctx context.Context, slot int) (*model.Tray, error) {
	tray, err := svc.store.Tray.GetTray(ctx, slot)
	if err != nil {
		return nil, errors.Wrap(err, "svc.tray.GetTray")
	}
	if tray == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("Tray with slot '%d' not found", slot))
	}
	
	return tray, nil
}

// GetAllTrays returns all trays stored in the database
func (svc *TrayWebService) GetAllTrays(ctx context.Context) ([]*model.Tray, error) {
	tray, err := svc.store.Tray.GetAllTrays(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "svc.tray.GetAllTrays")
	}
	if tray == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprint("No trays found"))
	}
	
	return tray, nil
}

// CreateTray ...
func (svc TrayWebService) CreateTray(ctx context.Context, reqTray *model.Tray) (*model.Tray, error) {
	_, err := svc.store.Tray.CreateTray(ctx, reqTray)
	if err != nil {
		return nil, errors.Wrap(err, "svc.tray.CreateTray error")
	}
	
	// get created tray by ID
	createdTray, err := svc.store.Tray.GetTray(ctx, reqTray.Slot)
	if err != nil {
		return nil, errors.Wrap(err, "svc.tray.GetTray error")
	}
	
	return createdTray, nil
}

// UpdateTray ...
func (svc *TrayWebService) UpdateTray(ctx context.Context, reqTray *model.Tray) (*model.Tray, error) {
	trayDB, err := svc.store.Tray.GetTray(ctx, reqTray.Slot)
	if err != nil {
		return nil, errors.Wrap(err, "svc.tray.GetTray error")
	}
	if trayDB == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("Tray '%d' not found", reqTray.Slot))
	}
	
	// update tray
	_, err = svc.store.Tray.UpdateTray(ctx, reqTray)
	if err != nil {
		return nil, errors.Wrap(err, "svc.tray.UpdateTray error")
	}
	
	// get updated tray by ID
	updatedTray, err := svc.store.Tray.GetTray(ctx, reqTray.Slot)
	if err != nil {
		return nil, errors.Wrap(err, "svc.tray.GetTray error")
	}
	
	return updatedTray, nil
}

// DeleteTray ...
func (svc *TrayWebService) DeleteTray(ctx context.Context, slot int) error {
	// Check if tray exists
	trayDB, err := svc.store.Tray.GetTray(ctx, slot)
	if err != nil {
		return errors.Wrap(err, "svc.tray.GetTray error")
	}
	if trayDB == nil {
		return errors.Wrap(types.ErrNotFound, fmt.Sprintf("Tray with slot '%d' not found", slot))
	}
	
	err = svc.store.Tray.DeleteTray(ctx, slot)
	if err != nil {
		return errors.Wrap(err, "svc.tray.DeleteTray error")
	}
	
	return nil
}
