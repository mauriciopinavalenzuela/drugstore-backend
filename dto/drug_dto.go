package dto

import "time"

type DrugDTO struct {
	ID                 int       `json:"id"`
	GenericName        string    `json:"generic_name"`
	DrugTypeID         int       `json:"drug_type_id"`
	Price              float64   `json:"price"`
	AdministrationForm string    `json:"administration_form"`
	Dosage             string    `json:"dosage"`
	Description        string    `json:"description"`
	ExpirationDate     time.Time `json:"expiration_date"`
	Stock              int       `json:"stock"`
}
