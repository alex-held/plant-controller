package model

import (
	"time"
)

type TrayConfig struct {
	Id        int64
	PlantId   int
	Plant     *Plant `pg:"rel:has-one"`
	StartedAt time.Time
}
