package services

import (
	"drugstore-backend/dto"
	"drugstore-backend/models"
	"errors"
	"time"
)

func ConvertDetailDTOToSaleDetail(dto dto.DetailDTO) models.SaleDetail {
	return models.SaleDetail{
		ID:        dto.ID,
		DrugID:    dto.DrugID,
		Quantity:  dto.Quantity,
		UnitPrice: dto.UnitPrice,
	}
}

func ConvertDetailsDTOToSaleDetails(details []dto.DetailDTO) []models.SaleDetail {
	saleDetails := make([]models.SaleDetail, len(details))
	for i, detail := range details {
		saleDetails[i] = ConvertDetailDTOToSaleDetail(detail)
	}
	return saleDetails
}

func MakeSale(customerID int, details []dto.DetailDTO) (models.Sale, error) {
	total := 0.0
	for _, detail := range details {
		drugDTO, err := GetDrugAsDTO(detail.DrugID)
		if err != nil {
			return models.Sale{}, err
		}
		if detail.Quantity > drugDTO.Stock {
			return models.Sale{}, errors.New("insufficient stock")
		}
		total += float64(detail.Quantity) * detail.UnitPrice
		drug := ConvertDrugDTOToDrug(drugDTO)
		drug.Stock -= detail.Quantity
		drugs[detail.DrugID] = drug
	}

	sale := models.Sale{
		ID:         len(drugs) + 1,
		CustomerID: customerID,
		Date:       time.Now(),
		Total:      total,
		Details:    ConvertDetailsDTOToSaleDetails(details),
	}

	return sale, nil
}
