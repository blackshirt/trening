package models

type Asn struct {
	ID            *int    `json:"id"`
	ObjID         *string `json:"obj_id"`
	Name          *string `json:"name"`
	Nip           *string `json:"nip"`
	CurrentJob    *string `json:"current_job"`
	CurrentGrade  *string `json:"current_grade"`
	CurrentPlaces *Opd    `json:"current_places"`
}

type Opd struct {
	ID         *int    `json:"id"`
	ObjID      *string `json:"obj_id"`
	Name       *string `json:"name"`
	LongName   *string `json:"long_name"`
	RoadNumber *string `json:"road_number"`
	City       *string `json:"city"`
	Province   *string `json:"province"`
}

type Org struct {
	ID         *int    `json:"id"`
	ObjID      *string `json:"obj_id"`
	Name       *string `json:"name"`
	LongName   *string `json:"long_name"`
	RoadNumber *string `json:"road_number"`
	City       *string `json:"city"`
	Province   *string `json:"province"`
}

type Trx struct {
	ID          *int     `json:"id"`
	ObjID       *string  `json:"obj_id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Category    *TrxCat  `json:"category"`
	Type        *TrxType `json:"type"`
}

type TrxDetail struct {
	ID           *int    `json:"id"`
	ObjID        *string `json:"obj_id"`
	Trx          Trx     `json:"trx"`
	Start        *string `json:"start"`
	Finish       *string `json:"finish"`
	Organizer    *Org    `json:"organizer"`
	Location     *Org    `json:"location"`
	Participants []Asn   `json:"participants"`
}
