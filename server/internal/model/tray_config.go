package model

import (
	"time"
)

type TrayConfig struct {
	Id        int64
	PlantId   int
	StartedAt string
}

func (m *TrayConfig) ToDB() *TrayConfigDB {
	startedAt, err := time.Parse(time.RFC3339, m.StartedAt)
	if err != nil || m.StartedAt == "" {
		startedAt = time.Now()
	}
	return &TrayConfigDB{
		Id:        m.Id,
		PlantId:   m.PlantId,
		StartedAt: startedAt,
	}
}

type TrayConfigDB struct {
	Id        int64 `json:"id,pk,unique"`
	PlantId   int
	Plant     *PlantDB `pg:"rel:has-one"`
	StartedAt time.Time
}

func (t *TrayConfigDB) ToApp() *TrayConfig {
	return &TrayConfig{
		Id:        t.Id,
		PlantId:   t.PlantId,
		StartedAt: t.StartedAt.Format(time.RFC3339),
	}
}
