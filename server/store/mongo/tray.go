package mongo

import (
	"context"
	
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	
	"plant-controller/logger"
	"plant-controller/model"
)

type TrayMongoRepo struct {
	db    *MongoDB
	trays *mongo.Collection
}

// NewTrayRepo
func NewTrayRepo(db *MongoDB) *TrayMongoRepo {
	return &TrayMongoRepo{
		db:    db,
		trays: db.Database.Collection("trays"),
	}
}

func (t *TrayMongoRepo) GetAllTrays(ctx context.Context) ([]*model.Tray, error) {
	res, err := t.trays.Find(ctx, bson.M{})
	logger := logger.Get()
	if err != nil {
		logger.Err(errors.Wrap(err, "Error while fetching trays."))
		return nil, err
	}
	var trays []*model.Tray
	err = res.All(ctx, &trays)
	if err != nil {
		logger.Err(errors.Wrap(err, "Error while decoding trays."))
		return nil, err
	}
	return trays, err
}

func (t *TrayMongoRepo) GetTray(ctx context.Context, slot int) (*model.Tray, error) {
	var res model.Tray
	err := t.trays.
		FindOne(ctx, bson.M{"slot": slot}).
		Decode(&res)
	
	logger := logger.Get()
	if err != nil {
		logger.Err(errors.Wrap(err, "Error while finding / decoding tray."))
		return nil, err
	}
	
	logger.Debug().Msgf("Found Tray for slot %d\n%n", slot, res)
	return &res, err
}

func (t *TrayMongoRepo) CreateTray(ctx context.Context, tray *model.Tray) (*model.Tray, error) {
	_, err := t.trays.InsertOne(ctx, tray)
	logger := logger.Get()
	if err != nil {
		logger.Err(errors.Wrap(err, "Error while inserting tray"))
	}
	logger.Debug().Msgf("Created Tray for Slot %d", tray.Slot)
	return tray, err
}

func (t *TrayMongoRepo) UpdateTray(ctx context.Context, tray *model.Tray) (*model.Tray, error) {
	panic("implement me")
}

func (t *TrayMongoRepo) DeleteTray(ctx context.Context, slot int) error {
	panic("implement me")
}
