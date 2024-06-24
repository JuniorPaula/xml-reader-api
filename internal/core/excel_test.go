package core

import (
	"sync"
	"testing"
	"xml-reader-api/internal/models"
)

const testFilePath = "./datatest/test_Reconfile_fornecedores.xlsx"
const testSheetName = "Planilha1"

func TestReadExelFile(t *testing.T) {
	dataChannel := make(chan models.Supplier)
	go ReadExcel(dataChannel, testFilePath, testSheetName)

	receivedSuppliers := []models.Supplier{}
	for supplier := range dataChannel {
		receivedSuppliers = append(receivedSuppliers, supplier)
	}

	expectedSuppliers := models.Supplier{
		ParterID: "PartnerId",
	}

	if len(receivedSuppliers) != 1 {
		t.Errorf("Expected 1 supplier, got %d", len(receivedSuppliers))
	}

	receivedSupplier := receivedSuppliers[0]
	if receivedSupplier.ParterID != expectedSuppliers.ParterID {
		t.Errorf("Expected %v, got %v", expectedSuppliers.ParterID, receivedSupplier.ParterID)
	}
}

func TestLoadSupplierList(t *testing.T) {
	dataChannel := make(chan models.Supplier)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		ReadExcel(dataChannel, testFilePath, testSheetName)
	}()

	receivedSuppliers := []models.Supplier{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		LoadSupplierList(dataChannel, &receivedSuppliers)
	}()

	wg.Wait()

	expectedSuppliers := models.Supplier{
		ParterID: "PartnerId",
	}

	if len(receivedSuppliers) != 1 {
		t.Errorf("Expected 1 supplier, got %d", len(receivedSuppliers))
	}

	receivedSupplier := receivedSuppliers[0]
	if receivedSupplier.ParterID != expectedSuppliers.ParterID {
		t.Errorf("Expected %v, got %v", expectedSuppliers.ParterID, receivedSupplier.ParterID)
	}
}
