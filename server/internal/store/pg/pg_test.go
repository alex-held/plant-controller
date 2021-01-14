package pg

import (
	"fmt"
	"testing"
	
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	
	"plant-controller/internal/model"
)

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User `pg:"rel:has-one"`
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

func TestCreatePlantAndGrowConfig(t *testing.T) {
	db := initDB([]interface{}{
		(*model.PlantDB)(nil),
		(*model.GrowConfigDB)(nil),
	})
	defer db.Close()
	
	alfalfa := &model.Plant{
		Name: "Alfalfa",
	}
	_, err := db.Model(alfalfa).Insert()
	if err != nil {
		panic(err)
	}
	alfalfaConfig := &model.GrowConfigDB{
		PlantId:           alfalfa.Id,
		Titel:           "Alfalfa Config 1",
		Description:     "First Config",
		GerminationDays: 3,
		GrowingDays:     5,
	}
	_, err = db.Model(alfalfaConfig).Insert()
	if err != nil {
		panic(err)
	}
	
	mungo := &model.Plant{
		Name: "Mungo",
	}
	_, err = db.Model(mungo).Insert()
	if err != nil {
		panic(err)
	}
	mungoConfig := &model.GrowConfigDB{
		Plant:           mungo,
		Titel:           "Mungo Config 1",
		Description:     "First Config",
		GerminationDays: 4,
		GrowingDays:     4,
	}
	
	_, err = db.Model(mungoConfig).Insert()
	if err != nil {
		panic(err)
	}
	
	var plants []model.Plant
	db.Model(&plants).Select()
	fmt.Printf("%+v\n", plants)
	
	var gcfgs []model.GrowConfig
	db.Model(&gcfgs).Relation("Plant").Select()
	fmt.Printf("\n\n%+v\n", gcfgs)
	fmt.Printf("\n\n%+v\n", gcfgs[0])
	
}


func TestCreateSchema(t *testing.T) {
	
	db := initDB([]interface{}{
		(*User)(nil),
		(*Story)(nil),
	})
	defer db.Close()
	user1 := &User{
		Name:   "admin",
		Emails: []string{"admin1@admin", "admin2@admin"},
	}
	_, err := db.Model(user1).Insert()
	if err != nil {
		panic(err)
	}
	
	_, err = db.Model(&User{
		Name:   "root",
		Emails: []string{"root1@root", "root2@root"},
	}).Insert()
	if err != nil {
		panic(err)
	}
	
	story1 := &Story{
		Title:    "Cool story",
		AuthorId: user1.Id,
	}
	_, err = db.Model(story1).Insert()
	if err != nil {
		panic(err)
	}
	
	// Select user by primary key.
	user := &User{Id: user1.Id}
	err = db.Model(user).WherePK().Select()
	if err != nil {
		panic(err)
	}
	
	// Select all users.
	var users []User
	err = db.Model(&users).Select()
	if err != nil {
		panic(err)
	}
	
	// Select story and associated author in one query.
	story := new(Story)
	err = db.Model(story).
		Relation("Author").
		Where("story.id = ?", story1.Id).
		Select()
	if err != nil {
		panic(err)
	}
	
	fmt.Println(user)
	fmt.Println(users)
	fmt.Println(story)
}

func initDB(models []interface{}) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5433",
		User:     "postgres",
		Password: "postgres",
		Database: "users",
	})
	
	
	if err := createSchema(db, models); err != nil {
		panic(err)
	}
	return db
}

func createSchema(db *pg.DB, models []interface{}) error  {
	
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
