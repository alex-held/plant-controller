package model

import (
	"fmt"
)

type Tray struct {
	Id           int64
	IsActive     bool
	Slot         int    `json:"slot"`
	GrowConfigId int64  `json:"grow_config_id"`
	TrayConfigId int64  `json:"tray_config_id"`
	Plant        string `json:"plant_type"`
	StartDate    string `json:"start_date"`
}

func (t *Tray) ToDB() *TrayDB {
	return &TrayDB{
		Id:           t.Id,
		Slot:         t.Slot,
		IsActive:     t.IsActive,
		GrowConfigId: t.GrowConfigId,
		TrayConfigId: t.TrayConfigId,
	}
}

type TrayDB struct {
	Id           int64
	Slot         int
	IsActive     bool
	GrowConfigId int64
	GrowConfig   *GrowConfig `pg:"rel:has-one"`
	TrayConfigId int64
	TrayConfig   *TrayConfig `pg:"rel:has-one"`
}

func (t Tray) String() string {
	return fmt.Sprintf("Id=%d; IsActive=%t; Slot=%d; Plant=%s", t.Id, t.IsActive, t.Slot, t.Plant)
}
