package model

type Plant struct {
	Id        int64
	Name      string
}


type PlantDB struct {
	Id        int64
	Name      string
}


func (p *Plant) ToDB() *PlantDB {
	return &PlantDB{
		Id:   p.Id,
		Name: p.Name,
	}
}

func (p *PlantDB) ToApp() *Plant {
	 return &Plant{
		 Id:   p.Id,
		 Name: p.Name,
	 }
}
