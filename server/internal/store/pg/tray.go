package pg

import (
	"context"
	
	"github.com/pkg/errors"
	
	"plant-controller/internal/model"
)

type TrayPGRepo struct {
	db *DB
}

// NewTrayRepo
func NewTrayRepo(db *DB) *TrayPGRepo {
	return &TrayPGRepo{
		db: db,
	}
}

func (g TrayPGRepo) GetNextId() int64 {
	res := &model.TrayDB{}
	g.db.Model(res).Last()
	return res.Id + 1
}

func (g TrayPGRepo) Get(ctx context.Context, id int64) (*model.TrayDB, error) {
	result := &model.TrayDB{Id: id}
	
	if err := g.db.WithContext(ctx).
		Model(result).
		Relation("TrayConfig").
		Relation("GrowConfig").
		WherePK().
		Select(); err != nil {
		return nil, errors.Wrapf(err, "[repo.tray.getall] error loading trays")
	}
	
	return result, nil
}

func (g TrayPGRepo) GetAll(ctx context.Context) ([]*model.TrayDB, error) {
	var result []*model.TrayDB
	
	if err := g.db.WithContext(ctx).
		Model(&result).
		Relation("TrayConfig").
		Relation("GrowConfig").
		Select(); err != nil {
		return nil, errors.Wrapf(err, "[repo.tray.getall] error loading trays")
	}
	
	return result, nil
}

func (g TrayPGRepo) Create(ctx context.Context, gc *model.TrayDB) (*model.TrayDB, error) {
	_, err := g.db.WithContext(ctx).
		Model(gc).
		Relation("TrayConfig").
		Relation("GrowConfig").
		Insert()
	
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.tray.create] error creating tray")
	}
	
	created := &model.TrayDB{Id: gc.Id}
	g.db.WithContext(ctx).
		Model(created).
		Relation("TrayConfig").
		Relation("GrowConfig").
		WherePK().
		Select()
	
	return created, nil
}

func (g TrayPGRepo) Update(ctx context.Context, m *model.TrayDB) (*model.TrayDB, error) {
	_, err := g.db.WithContext(ctx).
		Model(m).
		Relation("TrayConfig").
		Relation("GrowConfig").
		Update()
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.tray.update] error updating tray")
	}
	
	updated := &model.TrayDB{Id: m.Id}
	g.db.WithContext(ctx).
		Model(updated).
		Relation("TrayConfig").
		Relation("GrowConfig").
		WherePK().
		Select()
	
	return updated, nil
}

func (g TrayPGRepo) Delete(ctx context.Context, id int64) error {
	_, err := g.db.WithContext(ctx).
		Model(&model.TrayDB{
			Id: id,
		}).
		Relation("TrayConfig").
		Relation("GrowConfig").
		Delete()
	
	if err != nil {
		return errors.Wrapf(err, "[repo.tray.delete] error deleting tray")
	}
	
	return nil
}
