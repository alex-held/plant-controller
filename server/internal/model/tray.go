package model

type Tray struct {
	Id           int64 `json:"id"`
	IsActive     bool  `json:"is_active"`
	Slot         int   `json:"slot"`
	GrowConfigId int64 `json:"grow_config_id"`
	TrayConfigId int64 `json:"tray_config_id"`
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
	GrowConfig   *GrowConfigDB `pg:"rel:has-one"`
	TrayConfigId int64
	TrayConfig   *TrayConfigDB `pg:"rel:has-one"`
}

func (d *TrayDB) ToApp() *Tray {
	return &Tray{
		Id:           d.Id,
		IsActive:     d.IsActive,
		Slot:         d.Slot,
		GrowConfigId: d.GrowConfigId,
		TrayConfigId: d.TrayConfigId,
	}
}
