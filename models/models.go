package models

import "time"

type Training struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	HeldAt       Schedule  `json:"heldAt"`
	Organizer    Organizer `json:"organizer"`
	Participants []ASN     `json:"participants"`
}

type Schedule struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	Finish      time.Time `json:"finish"`
}

type Organizer struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Longname string  `json:"longname"`
	Address  Address `json:"address"`
}

type ASN struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Nip           string `json:"nip"`
	CurrentJob    string `json:"current_job"`
	CurrentGrade  string `json:"current_grade"`
	CurrentPlaces OPD    `json:"current_places"`
}

type Address struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Roadname string `json:"roadname"`
	Number   string `json:"number"`
	City     string `json:"city"`
	Province string `json:"province"`
}

type OPD struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}
