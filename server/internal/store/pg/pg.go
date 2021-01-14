package pg

import (
	"time"
	
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	
	"plant-controller/internal/config"
	 "plant-controller/internal/model"
)

// Timeout is a Postgres timeout
const Timeout = 5

// DB is a shortcut structure to a Postgres DB
type DB struct {
	*pg.DB
}

// Dial creates new database connection to postgres
func Dial() (*DB, error) {
	cfg := config.Get()
	if cfg.PgURL == "" {
		return nil, nil
	}
	pgOpts, err := pg.ParseURL(cfg.PgURL)
	if err != nil {
		return nil, err
	}
	
	pgDB := pg.Connect(pgOpts)
	
	_, err = pgDB.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	
	err = createSchema(pgDB)
	if err != nil {
		return nil, err
	}
	
	
	pgDB.WithTimeout(time.Second * time.Duration(Timeout))
	
	return &DB{pgDB}, nil
}



func createSchema(db *pg.DB) error  {
	
	models := []interface{}{
		(*model.PlantDB) (nil),
		(*model.GrowConfigDB) (nil),
		(*model.TrayConfigDB) (nil),
		(*model.TrayDB) (nil),
	}
	
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:          false,
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
