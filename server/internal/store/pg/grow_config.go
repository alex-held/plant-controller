package pg

import (
	"context"
	
	"github.com/pkg/errors"
	
	"plant-controller/internal/model"
)

type GrowConfigPGRepo struct {
	db *DB
}

func (g GrowConfigPGRepo) GetNextId() int64 {
	res := &model.GrowConfigDB{}
	g.db.Model(res).Last()
	return res.Id + 1
}

// NewGrowConfigRepo
func NewGrowConfigRepo(db *DB) *GrowConfigPGRepo {
	return &GrowConfigPGRepo{
		db: db,
	}
}

func (g GrowConfigPGRepo) Get(ctx context.Context, id int64) (*model.GrowConfigDB, error) {
	result := &model.GrowConfigDB{Id: id}
	
	if err := g.db.WithContext(ctx).
		Model(result).
		Relation("Plant").
		WherePK().
		Select(); err != nil {
		return nil, errors.Wrapf(err, "[repo.growconfig.getall] error loading growconfigs")
	}
	
	return result, nil
}

func (g GrowConfigPGRepo) GetAll(ctx context.Context) ([]*model.GrowConfigDB, error) {
	var result []*model.GrowConfigDB
	
	if err := g.db.WithContext(ctx).
		Model(&result).
		Relation("Plant").
		Select(); err != nil {
		return nil, errors.Wrapf(err, "[repo.growconfig.getall] error loading growconfigs")
	}
	
	return result, nil
}

func (g GrowConfigPGRepo) Create(ctx context.Context, gc *model.GrowConfigDB) (*model.GrowConfigDB, error) {
	gc.Id = g.GetNextId()
	
	_, err := g.db.WithContext(ctx).
		Model(gc).
		Relation("Plant").
		Insert()
	
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.growconfig.create] error creating growconfig")
	}
	
	created := &model.GrowConfigDB{Id: gc.Id}
	g.db.WithContext(ctx).
		Model(created).
		WherePK().
		Select()
	
	return created, nil
}

func (g GrowConfigPGRepo) Update(ctx context.Context, m *model.GrowConfigDB) (*model.GrowConfigDB, error) {
	_, err := g.db.WithContext(ctx).
		Model(m).
		WherePK().
		Update()
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.growconfig.update] error updating growconfig")
	}
	
	updated := &model.GrowConfigDB{Id: m.Id}
	g.db.WithContext(ctx).
		Model(updated).
		WherePK().
		Select()
	
	return updated, nil
}

func (g GrowConfigPGRepo) Delete(ctx context.Context, id int64) error {
	_, err := g.db.WithContext(ctx).
		Model(&model.GrowConfigDB{
			Id: id,
		}).
		WherePK().
		Delete()
	
	if err != nil {
		return errors.Wrapf(err, "[repo.growconfig.delete] error deleting growconfig")
	}
	
	return nil
}
