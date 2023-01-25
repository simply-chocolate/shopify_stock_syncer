package utils

import (
	"fmt"
	"pcn_stock_syncer/shopify_api_wrapper"
)

// TODO: Take a look at all the endpoints again and double check if everything is done correctly.
func HandleSyncStock() error {

	ShopifyProducts, err := shopify_api_wrapper.ShopifyApiGetProducts_AllPages(shopify_api_wrapper.ShopifyApiQueryParams{
		Fields: []string{"variants"},
		Status: "active",
		//Ids:    []string{"6748027912399"},
		Limit: 20,
	})
	if err != nil {
		return err
	}

	PcnProducts, err := PcnApiGetStockData()
	if err != nil {
		return err
	}

	ShopifyInventoryId, err := shopify_api_wrapper.ShopifyApiGetInventoryId(shopify_api_wrapper.ShopifyApiQueryParams{
		Fields: []string{"id,name,address1"},
	})
	if err != nil {
		return err
	}

	ShopifyInventoryItems, err := shopify_api_wrapper.ShopifyApiInventoryItem_AllItems(ShopifyProducts)
	if err != nil {
		return err
	}

	// Iterate over all products from shopify and all variants from each product.
	for _, product := range ShopifyProducts.Body.Products {
		for _, variant := range product.Variants {
			var isTracked bool

			// Check if the inventory of this variant is tracked in shopify.
			for _, InventoryItem := range ShopifyInventoryItems.Body.InventoryItems {
				if InventoryItem.InventoryItemId == variant.InventoryItemId {

					isTracked = InventoryItem.IsTracked
				}
			}
			if !isTracked {
				continue
			}

			product, exists := PcnProducts[variant.Barcode]
			if !exists {
				fmt.Printf("product with barcode %s does not exist in pcn\n", variant.Barcode)
				// TODO: Implement the teams error handling. and make this send an error to teams.
				continue

			} else {
				if err := shopify_api_wrapper.SetInventoryLevel(&shopify_api_wrapper.SetInventoryLevelBody{
					Location_id:       ShopifyInventoryId.Body.Locations[0].LocationId,
					Inventory_item_id: variant.InventoryItemId,
					Available:         product,
				}); err != nil {
					fmt.Printf("error setting inventory lvl for item: %s\n", variant.Barcode)
					return err
				}
			}
		}
	}

	return nil
}
