package models

type Training struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Start       string `json:"start"`
	Finish      string `json:"finish"`
	Organizer    Organisasi `json:"organizer"`
	Location    Organisasi `json:"location"`
	Participants []ASN      `json:"participants"`
}

type Organisasi struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	LongName string  `json:"long_name"`
 Road     string `json:"road"`
	Number   string `json:"number"`
	City     string `json:"city"`
	Province string `json:"province"`
}


type OPD struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LongName string `json:"long_name"`
	Road     string `json:"road"`
	Number   string `json:"number"`
	City     string `json:"city"`
	Province string `json:"province"`
}


type ASN struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Nip           string `json:"nip"`
	CurrentJob    string `json:"current_job"`
	CurrentGrade  string `json:"current_grade"`
	CurrentPlaces OPD    `json:"current_places"`
}
