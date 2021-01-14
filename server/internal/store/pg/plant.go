package pg

import (
	"context"
	
	"github.com/pkg/errors"
	
	"plant-controller/internal/model"
)

type PlantPGRepo struct {
	db *DB
}

// NewPlantRepo
func NewPlantRepo(db *DB) *PlantPGRepo {
	return &PlantPGRepo{
		db: db,
	}
}

func (g PlantPGRepo) GetNextId() int64 {
	res := &model.PlantDB{}
	g.db.Model(res).Last()
	return res.Id + 1
}

func (g PlantPGRepo) Get(ctx context.Context, id int64) (*model.PlantDB, error) {
	result := &model.PlantDB{Id: id}
	
	if err := g.db.WithContext(ctx).
		Model(result).
		WherePK().
		Select(); err != nil {
		return nil, errors.Wrapf(err, "[repo.plant.getall] error loading plants")
	}
	
	return result, nil
}

func (g PlantPGRepo) GetAll(ctx context.Context) ([]*model.PlantDB, error) {
	var result []*model.PlantDB
	
	if err := g.db.WithContext(ctx).
		Model(&result).
		Select(); err != nil {
		return nil, errors.Wrapf(err, "[repo.plant.getall] error loading plants")
	}
	
	return result, nil
}

func (g PlantPGRepo) Create(ctx context.Context, name string) (*model.PlantDB, error) {
	m := &model.PlantDB{
		Id:   g.GetNextId(),
		Name: name,
	}
	_, err := g.db.	WithContext(ctx).
		Model(m).
		Insert()
	
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.plant.create] error creating plant")
	}
	
	created := &model.PlantDB{Id: m.Id}
	g.db.WithContext(ctx).
		Model(created).
		WherePK().
		Select()
	
	return created, nil
}

func (g PlantPGRepo) Update(ctx context.Context, m *model.PlantDB) (*model.PlantDB, error) {
	_, err := g.db.Model(m).Update()
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.plant.update] error updating plant")
	}
	
	updated := &model.PlantDB{Id: m.Id}
	g.db.WithContext(ctx).
		Model(updated).
		WherePK().
		Select()
	
	return updated, nil
}

func (g PlantPGRepo) Delete(ctx context.Context, id int64) error {
	_, err := g.db.
		WithContext(ctx).
		Model(&model.PlantDB{
		Id: id,
	}).Delete()
	
	if err != nil {
		return errors.Wrapf(err, "[repo.plant.delete] error deleting plant")
	}
	
	return nil
}
