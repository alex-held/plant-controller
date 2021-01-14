package model

import (
	"fmt"
)

type PlantType int

const (
	Empty PlantType = iota
	Mungo
	Alfafa
	Ruccola
)


func (p PlantType) Stringer() string {
	switch p {
	case Empty:
		return "Empty"
	case Alfafa:
		return "Alfafa Samen"
	case Mungo:
		return "Mungo Bohnen"
	case Ruccola:
		return "Ruccola Samen"
	default:
		panic(fmt.Errorf("Unknown Plant Type. "))
	}
}
