package utils

import (
	"fmt"
	"pcn_stock_syncer/pcn_api_wrapper"
	"time"
)

// Get all the stockdata from PCN and return a map of barcodes to available stock.
func PcnApiGetStockData() (map[string]int, error) {
	PcnProducts, err := pcn_api_wrapper.PcnApiGetStockData()
	if err != nil {
		return map[string]int{}, err
	}

	// Create a map of barcodes to available stock.
	BarcodeToAvailableStock := make(map[string]int)
	for _, product := range PcnProducts.Body.Results {
		BarcodeToAvailableStock[product.Barcode] = product.Available - product.OnOrder
	}

	return BarcodeToAvailableStock, nil
}

func PcnApiGetBundleProducts() (map[string]string, error) {
	bundleProducts, statusCode, err := pcn_api_wrapper.PcnApiBundleProductList()

	if err != nil {
		fmt.Println(err)

		if statusCode == 403 {
			time.Sleep(6 * time.Minute)
			bundleProducts, _, err = pcn_api_wrapper.PcnApiBundleProductList()
			if err != nil {
				fmt.Println(err)
				return map[string]string{}, err
			}
		}
	}
	BundleProductList := make(map[string]string)

	for _, bundleProduct := range bundleProducts.Body.Results {
		BundleProductList[bundleProduct.BundleCode] = ""
	}

	return BundleProductList, nil
}
