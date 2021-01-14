package service

import (
	"context"
	"errors"
	
	"plant-controller/store"
)

// Manager is just a collection of all services we have in the project
type Manager struct {
	Tray        TrayService
}

// NewManager creates new service manager
func NewManager(ctx context.Context, store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}
	return &Manager{
		Tray: NewTrayWebService(ctx, store),
	}, nil
}
