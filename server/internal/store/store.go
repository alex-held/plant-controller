package store

import (
	"context"
	"time"
	
	"github.com/pkg/errors"
	
	"plant-controller/internal/logger"
	"plant-controller/internal/store/pg"
)

// Store contains all repositories
type Store struct {
	//	Mongo      *mongo.MongoDB
	PG         *pg.DB
	Tray       TrayRepo
	TrayConfig TrayConfigRepo
	GrowConfig GrowConfigRepo
	Plant      PlantRepo
}

// New creates new store
func New(ctx context.Context) (*Store, error) {
	
	postgres, err := pg.Dial()
	if err != nil {
		return nil, errors.Wrap(err, "postgresDB.Dial failed")
	}
	var store Store
	
	// Init PostgresDB repositories
	if postgres != nil {
		store.PG = postgres
		go store.KeepAlive()
		store.GrowConfig = pg.NewGrowConfigRepo(postgres)
		store.TrayConfig = pg.NewTrayConfigRepo(postgres)
		store.Tray = pg.NewTrayRepo(postgres)
		store.Plant = pg.NewPlantRepo(postgres)
	}
	
	return &store, nil
}

// KeepAlivePollPeriod is a MongoDB keepalive check time period
const KeepAlivePollPeriod = 3

// KeepAlive makes sure MongoDB is alive and reconnects if needed
func (store *Store) KeepAlive() {
	logger := logger.Get()
	var err error
	for {
		// Check if PostgreSQL is alive every 3 seconds
		time.Sleep(time.Second * KeepAlivePollPeriod)
		lostConnect := false
		if store == nil {
			lostConnect = true
		} else if _, err := store.PG.Exec("SELECT 1"); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		logger.Debug().Msg("[store.KeepAlive] Lost Postgres connection. Restoring...")
		store.PG, err = pg.Dial()
		if err != nil {
			logger.Err(err)
			continue
		}
		logger.Debug().Msg("[store.KeepAlive] Postgres reconnected")
	}
}
