package services

import (
	"drugstore-backend/dto"
	"drugstore-backend/models"
	"errors"
)

var drugs = make(map[int]models.Drug)

func ConvertDrugDTOToDrug(dto dto.DrugDTO) models.Drug {
	return models.Drug{
		ID:                 dto.ID,
		GenericName:        dto.GenericName,
		DrugTypeID:         dto.DrugTypeID,
		Price:              dto.Price,
		AdministrationForm: dto.AdministrationForm,
		Dosage:             dto.Dosage,
		Description:        dto.Description,
		ExpirationDate:     dto.ExpirationDate,
		Stock:              dto.Stock,
	}
}

func ConvertDrugToDrugDTO(drug models.Drug) dto.DrugDTO {
	return dto.DrugDTO{
		ID:                 drug.ID,
		GenericName:        drug.GenericName,
		DrugTypeID:         drug.DrugTypeID,
		Price:              drug.Price,
		AdministrationForm: drug.AdministrationForm,
		Dosage:             drug.Dosage,
		Description:        drug.Description,
		ExpirationDate:     drug.ExpirationDate,
		Stock:              drug.Stock,
	}
}

func GetDrugAsDTO(ID int) (dto.DrugDTO, error) {
	drug, found := drugs[ID]
	if !found {
		return dto.DrugDTO{}, errors.New("drug not found")
	}
	return ConvertDrugToDrugDTO(drug), nil
}
