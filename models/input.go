package models

type ASNInput struct {
	Name          string   `json:"name"`
	Nip           string   `json:"nip"`
	CurrentJob    string   `json:"current_job"`
	CurrentGrade  string   `json:"current_grade"`
	CurrentPlaces OPDInput `json:"current_places"`
}

type OrgInput struct {
	Name       string `json:"name"`
	LongName   string `json:"long_name"`
	RoadNumber string `json:"road_number"`
	City       string `json:"city"`
	Province   string `json:"province"`
}

type OPDInput struct {
	Name       string `json:"name"`
	LongName   string `json:"long_name"`
	RoadNumber string `json:"road_number"`
	City       string `json:"city"`
	Province   string `json:"province"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type TrainingInput struct {
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Organizer    OrgInput   `json:"organizer"`
	Participants []ASNInput `json:"participants"`
}
