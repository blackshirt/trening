package models

type Org struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LongName   string `json:"long_name"`
	RoadNumber string `json:"road_number"`
	City       string `json:"city"`
	Province   string `json:"province"`
}

type Opd struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LongName   string `json:"long_name"`
	RoadNumber string `json:"road_number"`
	City       string `json:"city"`
	Province   string `json:"province"`
}

type Asn struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Nip           string `json:"nip"`
	CurrentJob    string `json:"current_job"`
	CurrentGrade  string `json:"current_grade"`
	CurrentPlaces Opd    `json:"current_places"`
}
