package models

import "time"

type DrugType struct {
	ID          int
	Name        string
	Description string
}

type Drug struct {
	ID                 int
	GenericName        string
	DrugTypeID         int
	Price              float64
	AdministrationForm string
	Dosage             string
	Description        string
	ExpirationDate     time.Time
	Stock              int
}

type DrugCategory struct {
	ID          int
	Name        string
	Description string
}

type Brand struct {
	ID          int
	Name        string
	Description string
}

type Inventory struct {
	ID            int
	DrugID        int
	Quantity      int
	Lot           string
	ReceptionDate time.Time
}

type Sale struct {
	ID         int
	CustomerID int
	Date       time.Time
	Total      float64
	Details    []SaleDetail
}

type SaleDetail struct {
	ID        int
	SaleID    int
	DrugID    int
	Quantity  int
	UnitPrice float64
}

type Customer struct {
	ID      int
	Name    string
	Address string
	Phone   string
	Email   string
}

type Prescription struct {
	ID         int
	CustomerID int
	Date       time.Time
	Details    []PrescriptionDetail
}

type PrescriptionDetail struct {
	ID             int
	PrescriptionID int
	DrugID         int
	Quantity       int
	Instructions   string
}

type Supplier struct {
	ID      int
	Name    string
	Address string
	Phone   string
	Email   string
}

type Promotion struct {
	ID          int
	Description string
	Discount    float64
	StartDate   time.Time
	EndDate     time.Time
	Drugs       []Drug
}
