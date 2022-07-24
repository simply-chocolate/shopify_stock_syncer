package utils

import (
	"fmt"
	"pcn_stock_syncer/pcn_api_wrapper"
	"pcn_stock_syncer/shopify_api_wrapper"
)

func HandleSyncStock() error {
	PcnProducts, err := pcn_api_wrapper.PcnApiGetStockData()
	if err != nil {
		return err
	}

	ShopifyInventoryId, err := shopify_api_wrapper.ShopifyApiGetInventoryId(shopify_api_wrapper.ShopifyApiQueryParams{
		Fields: []string{"id,name,address1"},
	})
	if err != nil {
		return err
	}

	ShopifyProduts, err := shopify_api_wrapper.ShopifyApiGetProducts(shopify_api_wrapper.ShopifyApiQueryParams{
		Fields:  []string{"variants"},
		Filters: []string{"limit=250"},
	})
	if err != nil {
		return err
	}

	for _, product := range ShopifyProduts.Body.Products {
		for _, variant := range product.Variants {
			for _, pcnProduct := range PcnProducts.Body.Results {
				//TODO: More checks
				//TODO: lav et udtræk omkring hvorvidt inventory bliver tracked på items
				// https://shopify.dev/api/admin-rest/2022-07/resources/inventoryitem

				if variant.Barcode == "" {
					continue
				}
				if len(variant.Barcode) == 4 {
					continue
				}

				if variant.Barcode == pcnProduct.Barcode {
					if err := shopify_api_wrapper.SetInventoryLevel(&shopify_api_wrapper.SetInventoryLevelBody{
						Location_id:       ShopifyInventoryId.Body.Locations[0].LocationId,
						Inventory_item_id: variant.InventoryItemId,
						Available:         pcnProduct.Available,
					}); err != nil {
						fmt.Printf("error setting inventory lvl for item: %s\n", variant.Barcode)
						return err
					}
				}
			}
		}
	}
	return nil
}
