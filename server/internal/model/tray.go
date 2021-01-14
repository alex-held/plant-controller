package model

import (
	"fmt"
	"time"
)

type Tray struct {
	Slot      int       `json:"slot"`
	PlantType PlantType `json:"plant_type"`
	StartDate time.Time `json:"start_date"`
}


func (t Tray) String() string {
	return fmt.Sprintf("Slot=%d; Plant=%s", t.Slot, t.PlantType.Stringer())
}


func (t Tray) isEmpty() bool {
	return t.PlantType == Empty
}

