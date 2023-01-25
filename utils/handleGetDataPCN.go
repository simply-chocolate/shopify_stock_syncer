package utils

import "pcn_stock_syncer/pcn_api_wrapper"

// Get all the stockdata from PCN and return a map of barcodes to available stock.
func PcnApiGetStockData() (map[string]int, error) {
	PcnProducts, err := pcn_api_wrapper.PcnApiGetStockData()
	if err != nil {
		return map[string]int{}, err
	}

	// Create a map of barcodes to available stock.
	BarcodeToAvailableStock := make(map[string]int)
	for _, product := range PcnProducts.Body.Results {
		BarcodeToAvailableStock[product.Barcode] = product.Available
	}

	return BarcodeToAvailableStock, nil
}
