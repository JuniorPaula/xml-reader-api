package core

import (
	"fmt"
	"strconv"
	"time"
	"xml-reader-api/internal/models"

	"github.com/xuri/excelize/v2"
)

func LoadSupplierList(dataChannel chan models.Supplier, supplierList *[]models.Supplier) {
	fmt.Println("Loading data, please wait...")
	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				fmt.Println("Processing completed")
				return
			}
			*supplierList = append(*supplierList, data)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}

func ReadExcel(dataChannel chan models.Supplier, filePath, sheetName string) {
	defer close(dataChannel)

	file, err := excelize.OpenFile(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rows, err := file.Rows(sheetName)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		_, err = rows.Columns()
		if err != nil {
			panic(err)
		}
	}

	for rows.Next() {
		cells, err := rows.Columns()
		if err != nil {
			panic(err)
		}
		if isEmptyRow(cells) {
			continue
		}

		mpnID, _ := strconv.Atoi(getCellValue(cells, 6))
		tier2MpnID, _ := strconv.Atoi(getCellValue(cells, 7))
		unitPrice, _ := strconv.ParseFloat(getCellValue(cells, 33), 64)
		quantity, _ := strconv.ParseFloat(getCellValue(cells, 34), 64)
		billingPreTaxTotal, _ := strconv.ParseFloat(getCellValue(cells, 36), 64)
		pricingPreTaxTotal, _ := strconv.ParseFloat(getCellValue(cells, 38), 64)
		effectiveUnitPrice, _ := strconv.ParseFloat(getCellValue(cells, 44), 64)
		pcToBCExchangeRate, _ := strconv.Atoi(getCellValue(cells, 45))
		partnerEarnedCreditPercentage, _ := strconv.Atoi(getCellValue(cells, 49))
		creditPercentage, _ := strconv.Atoi(getCellValue(cells, 50))

		supplier := models.Supplier{
			ParterID:                      getCellValue(cells, 0),
			PartnerName:                   getCellValue(cells, 1),
			CustomerID:                    getCellValue(cells, 2),
			CustomerName:                  getCellValue(cells, 3),
			CustomerDomainName:            getCellValue(cells, 4),
			CustomerCountry:               getCellValue(cells, 5),
			MpnID:                         mpnID,
			Tier2MpnID:                    tier2MpnID,
			InvoiceNumber:                 getCellValue(cells, 8),
			ProductID:                     getCellValue(cells, 9),
			SKUID:                         getCellValue(cells, 10),
			AvailabilityID:                getCellValue(cells, 11),
			SKUName:                       getCellValue(cells, 12),
			ProductName:                   getCellValue(cells, 13),
			PublisherName:                 getCellValue(cells, 14),
			PublisherID:                   getCellValue(cells, 15),
			SubscriptionDescription:       getCellValue(cells, 16),
			SubscriptionID:                getCellValue(cells, 17),
			ChargeStartDate:               getCellValue(cells, 18),
			ChargeEndDate:                 getCellValue(cells, 19),
			UsageDate:                     getCellValue(cells, 20),
			MeterType:                     getCellValue(cells, 21),
			MeterCategory:                 getCellValue(cells, 22),
			MeterID:                       getCellValue(cells, 23),
			MeterSubCategory:              getCellValue(cells, 24),
			MeterName:                     getCellValue(cells, 25),
			MeterRegion:                   getCellValue(cells, 26),
			Unit:                          getCellValue(cells, 27),
			ResourceLocation:              getCellValue(cells, 28),
			CostomerService:               getCellValue(cells, 29),
			ResourceGroup:                 getCellValue(cells, 30),
			ResourceURI:                   getCellValue(cells, 31),
			ChargeType:                    getCellValue(cells, 32),
			UnitPrice:                     unitPrice, // 33
			Quantity:                      quantity,  // 34
			UnitType:                      getCellValue(cells, 35),
			BillingPreTaxTotal:            billingPreTaxTotal, // 36
			BillingCurrency:               getCellValue(cells, 37),
			PricingPreTaxTotal:            pricingPreTaxTotal, // 38
			PricingCurrency:               getCellValue(cells, 39),
			ServiceInfo1:                  getCellValue(cells, 40),
			ServiceInfo2:                  getCellValue(cells, 41),
			Tags:                          getCellValue(cells, 42),
			AdditionalInfo:                getCellValue(cells, 43),
			EffectiveUnitPrice:            effectiveUnitPrice, // 44
			PCToBCExchangeRate:            pcToBCExchangeRate, // 45
			PCToBCExchangeRateDate:        getCellValue(cells, 46),
			EntitlementId:                 getCellValue(cells, 47),
			EntitlementDescription:        getCellValue(cells, 48),
			PartnerEarnedCreditPercentage: partnerEarnedCreditPercentage, //49
			CreditPercentage:              creditPercentage,              // 50
			CreditType:                    getCellValue(cells, 51),
			BenefitOrderId:                getCellValue(cells, 52),
			BenefitId:                     getCellValue(cells, 53),
		}

		dataChannel <- supplier
	}
}

func getCellValue(cells []string, index int) string {
	if index < len(cells) {
		return cells[index]
	}
	return ""
}

func isEmptyRow(cells []string) bool {
	for _, cell := range cells {
		if cell != "" {
			return false
		}
	}
	return true
}
