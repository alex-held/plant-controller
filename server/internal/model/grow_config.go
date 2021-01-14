package model

type GrowConfig struct {
	Id int `json:"id"`
	PlantType PlantType `json:"plant_type"`
	Name string      `json:"name"`
	GerminationDays int `json:"germination_days"`
	GrowingDays int `json:"growing_days"`
}
