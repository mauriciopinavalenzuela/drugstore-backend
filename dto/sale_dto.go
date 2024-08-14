package dto

import "time"

type SaleDTO struct {
	ID         int         `json:"id"`
	CustomerID int         `json:"customer_id"`
	Date       time.Time   `json:"date"`
	Total      float64     `json:"total"`
	Details    []DetailDTO `json:"details"`
}

type DetailDTO struct {
	ID        int     `json:"id"`
	DrugID    int     `json:"drug_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}
