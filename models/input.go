package models

type AsnInput struct {
	Name          string    `json:"name"`
	Nip           *string   `json:"nip"`
	CurrentJob    *string   `json:"current_job"`
	CurrentGrade  *string   `json:"current_grade"`
	CurrentPlaces *OpdInput `json:"current_places"`
}

type OpdInput struct {
	Name       string  `json:"name"`
	LongName   string  `json:"long_name"`
	RoadNumber *string `json:"road_number"`
	City       string  `json:"city"`
	Province   *string `json:"province"`
}

type OrgInput struct {
	Name       string  `json:"name"`
	LongName   string  `json:"long_name"`
	RoadNumber *string `json:"road_number"`
	City       string  `json:"city"`
	Province   *string `json:"province"`
}

type Pagination struct {
	Offset *int `json:"offset"`
	Limit  *int `json:"limit"`
}

type Trx struct {
	ID          *int     `json:"id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Category    *TrxCat  `json:"category"`
	Type        *TrxType `json:"type"`
}

type TrxCat struct {
	ID          *int    `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type TrxHistory struct {
	TrxID        *int    `json:"trxId"`
	Start        *string `json:"start"`
	Finish       *string `json:"finish"`
	Organizer    *Org    `json:"organizer"`
	Location     *Org    `json:"location"`
	Participants []*Asn  `json:"participants"`
}

type TrxInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type TrxType struct {
	ID          *int    `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type TrxTypeInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type TrxCatInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
