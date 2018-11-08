// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type ASNInput struct {
	Name          string   `json:"name"`
	Nip           string   `json:"nip"`
	CurrentJob    string   `json:"current_job"`
	CurrentGrade  string   `json:"current_grade"`
	CurrentPlaces OPDInput `json:"current_places"`
}

type OPDInput struct {
	Name     string `json:"name"`
	LongName string `json:"long_name"`
	Road     string `json:"road"`
	Number   string `json:"number"`
	City     string `json:"city"`
	Province string `json:"province"`
}

type OrgzInput struct {
	Name     string `json:"name"`
	LongName string `json:"long_name"`
	Road     string `json:"road"`
	Number   string `json:"number"`
	City     string `json:"city"`
	Province string `json:"province"`
}

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type TrainingInput struct {
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Organizer    OrgzInput  `json:"organizer"`
	Participants []ASNInput `json:"participants"`
}
