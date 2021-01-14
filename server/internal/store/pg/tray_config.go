package pg

import (
	"context"
	
	"github.com/pkg/errors"
	
	"plant-controller/internal/model"
)

type TrayConfigPGRepo struct {
	db *DB
}

func (g TrayConfigPGRepo) GetNextId() int64 {
	var res model.TrayConfigDB
	_ = g.db.Model(&res).Last()
	return res.Id + 1
}

// NewTrayConfigRepo
func NewTrayConfigRepo(db *DB) *TrayConfigPGRepo {
	return &TrayConfigPGRepo{
		db: db,
	}
}

func (g TrayConfigPGRepo) Get(ctx context.Context, id int64) (*model.TrayConfigDB, error) {
	result := &model.TrayConfigDB{Id: id}
	
	if err := g.db.Model(result).
		Relation("Plant").
		WherePK().
		Select(); err != nil {
		return nil, errors.Wrapf(err, "[repo.trayconfig.get] error loading trayconfigs")
	}
	
	return result, nil
}

func (g TrayConfigPGRepo) GetAll(ctx context.Context) ([]*model.TrayConfigDB, error) {
	var result []*model.TrayConfigDB
	
	if err := g.db.
		WithContext(ctx).
		Model(&result).
		Relation("Plant").
		Select(); err != nil {
		return nil, errors.Wrapf(err, "[repo.trayconfig.getall] error loading trayconfigs")
	}
	
	return result, nil
}

func (g TrayConfigPGRepo) Create(ctx context.Context, m *model.TrayConfigDB) (*model.TrayConfigDB, error) {
	m.Id = g.GetNextId()
	_, err := g.db.
		WithContext(ctx).
		Model(m).
		Relation("Plant").
		Insert()
	
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.trayconfig.create] error creating trayconfig")
	}
	
	created := &model.TrayConfigDB{Id: m.Id}
	g.db.Model(created).
		WherePK().
		Select()
	
	return created, nil
}

func (g TrayConfigPGRepo) Update(ctx context.Context, m *model.TrayConfigDB) (*model.TrayConfigDB, error) {
	_, err := g.db.Model(m).Update()
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.trayconfig.update] error updating trayconfig")
	}
	
	updated := &model.TrayConfigDB{Id: m.Id}
	g.db.Model(updated).
		WherePK().
		Select()
	
	return updated, nil
}

func (g TrayConfigPGRepo) Delete(ctx context.Context, id int64) error {
	_, err := g.db.Model(&model.TrayConfigDB{
		Id: id,
	}).Delete()
	
	if err != nil {
		return errors.Wrapf(err, "[repo.trayconfig.delete] error deleting trayconfig")
	}
	
	return nil
}
