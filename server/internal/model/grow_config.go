package model

type GrowConfig struct {
	Id              int64
	Titel           string
	Description     string
	GerminationDays int
	GrowingDays     int
	PlantId         int64
}

type GrowConfigDB struct {
	Id              int64
	Titel           string
	Description     string
	GerminationDays int
	GrowingDays     int
	PlantId         int64
	Plant           *PlantDB `pg:"rel:has-one"`
}

func (g *GrowConfig) ToDB() *GrowConfigDB {
	return &GrowConfigDB{
		Id:              g.Id,
		Titel:           g.Titel,
		Description:     g.Description,
		GerminationDays: g.GerminationDays,
		GrowingDays:     g.GrowingDays,
		PlantId:         g.PlantId,
	}
}

func (g *GrowConfigDB) ToApp() *GrowConfig {
	return &GrowConfig{
		Id:              g.Id,
		Titel:           g.Titel,
		Description:     g.Description,
		GerminationDays: g.GerminationDays,
		GrowingDays:     g.GrowingDays,
		PlantId:         g.PlantId,
	}
}
