package store

import (
	"context"
	"time"
	
	"github.com/pkg/errors"
	
	"plant-controller/internal/logger"
	"plant-controller/internal/store/mongo"
	"plant-controller/internal/store/pg"
)

// Store contains all repositories
type Store struct {
	Mongo *mongo.MongoDB
	PG *pg.DB
	Tray  TrayRepo
	GrowConfig GrowConfigRepo
}

// New creates new store
func New(ctx context.Context) (*Store, error) {

	// connect to MongoDB
	mongoDB, err := mongo.Dial()
	if err != nil {
		return nil, errors.Wrap(err, "mongoDB.Dial failed")
	}
	
	postgres, err := pg.Dial()
	if err != nil {
		return nil, errors.Wrap(err, "postgresDB.Dial failed")
	}
	var store Store
	
	// Init MongoDB repositories
	if mongoDB != nil {
		store.Mongo = mongoDB
		go store.KeepAlive()
		store.Tray = mongo.NewTrayRepo(mongoDB)
	}
	
	// Init PostgresDB repositories
	if postgres != nil {
		store.PG = postgres
		go store.KeepAlive()
		store.GrowConfig = pg.NewGrowConfigRepo(postgres)
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
		} else if err = store.Mongo.Client.Ping(context.TODO(), nil); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		logger.Debug().Msg("[store.KeepAlive] Lost MongoDB connection. Restoring...")
		store.Mongo, err = mongo.Dial()
		if err != nil {
			logger.Err(err)
			continue
		}
		logger.Debug().Msg("[store.KeepAlive] MongoDB reconnected")
	}
}
