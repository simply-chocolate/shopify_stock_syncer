package utils

import (
	"fmt"
	"pcn_stock_syncer/pcn_api_wrapper"
	"pcn_stock_syncer/shopify_api_wrapper"
)

// Handles the collection of all the queries.
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

	ShopifyProducts, err := shopify_api_wrapper.ShopifyApiGetProducts(shopify_api_wrapper.ShopifyApiQueryParams{
		Fields:  []string{"variants"},
		Filters: []string{"limit=10"},
	})
	if err != nil {
		return err
	}

	return nil

	ShopifyInventoryItems, err := shopify_api_wrapper.ShopifyApiInventoryItem_AllItems(ShopifyProducts)
	if err != nil {
		return err
	}

	for _, product := range ShopifyProducts.Body.Products {
		for _, variant := range product.Variants {

			isTracked := true

			for _, InventoryItem := range ShopifyInventoryItems.Body.InventoryItems {
				if InventoryItem.InventoryItemId == variant.InventoryItemId {

					isTracked = InventoryItem.IsTracked
				}
			}
			if !isTracked {

				continue
			}

			for _, pcnProduct := range PcnProducts.Body.Results {
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
