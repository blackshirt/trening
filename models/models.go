package models

import "time"

type Training struct {
	ID           int    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	HeldAt       Schedule  `json:"heldAt"`
	Organizer    Organisasi `json:"organizer"`
	Participants []ASN     `json:"participants"`
}

type Schedule struct {
	ID          int    `json:"id"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	Finish      time.Time `json:"finish"`
}

type Organisasi struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	LongName string  `json:"long_name"`
	Address  Address `json:"address"`
}

type ASN struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Nip           string `json:"nip"`
	CurrentJob    string `json:"current_job"`
	CurrentGrade  string `json:"current_grade"`
	CurrentPlaces OPD    `json:"current_places"`
}

type Address struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Road     string `json:"road"`
	Number   string `json:"number"`
	City     string `json:"city"`
	Province string `json:"province"`
}

type OPD struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}
