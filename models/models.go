package models

type Training struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Start        string `json:"start"`
	Finish       string `json:"finish"`
	Organizer    Org    `json:"organizer"`
	Location     Org    `json:"location"`
	Participants []ASN  `json:"participants"`
}

type Org struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LongName   string `json:"long_name"`
	RoadNumber string `json:"road_number"`
	City       string `json:"city"`
	Province   string `json:"province"`
}

type OPD struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LongName   string `json:"long_name"`
	RoadNumber string `json:"road_number"`
	City       string `json:"city"`
	Province   string `json:"province"`
}

type ASN struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Nip           string `json:"nip"`
	CurrentJob    string `json:"current_job"`
	CurrentGrade  string `json:"current_grade"`
	CurrentPlaces OPD    `json:"current_places"`
}

type ASNInput struct {
	Name          string   `json:"name"`
	Nip           string   `json:"nip"`
	CurrentJob    string   `json:"current_job"`
	CurrentGrade  string   `json:"current_grade"`
	CurrentPlaces OPDInput `json:"current_places"`
}

type OPDInput struct {
	Name       string `json:"name"`
	LongName   string `json:"long_name"`
	RoadNumber string `json:"road_number"`
	City       string `json:"city"`
	Province   string `json:"province"`
}

type OrgInput struct {
	Name       string `json:"name"`
	LongName   string `json:"long_name"`
	RoadNumber string `json:"road_number"`
	City       string `json:"city"`
	Province   string `json:"province"`
}

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type TrainingInput struct {
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Organizer    OrgInput   `json:"organizer"`
	Participants []ASNInput `json:"participants"`
}
