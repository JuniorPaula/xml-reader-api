package repository

import (
	"xml-reader-api/internal/core"
	"xml-reader-api/internal/models"
)

type SupplierRepository interface {
	GetSuppliers(limit, offset int) ([]models.Supplier, error)
}

type supplierRepository struct {
	FilePath  string
	SheetName string
}

var (
	dataChannel  = make(chan models.Supplier)
	supplierList []models.Supplier
)

// NewSupplierRepository creates a new instance of SupplierRepository.
// It also starts reading the excel file and loading the supplier list.
func NewSupplierRepository(filePath, sheetName string) SupplierRepository {
	go core.ReadExcel(dataChannel, filePath, sheetName)
	go core.LoadSupplierList(dataChannel, &supplierList)

	return &supplierRepository{}
}

// GetSuppliers implements SupplierRepository.
func (s *supplierRepository) GetSuppliers(limit, offset int) ([]models.Supplier, error) {
	start := offset
	end := offset + limit
	if start > len(supplierList) {
		start = len(supplierList)
	}
	if end > len(supplierList) {
		end = len(supplierList)
	}

	response := (supplierList)[start:end]

	return response, nil
}
