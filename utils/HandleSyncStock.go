package utils

import (
	"pcn_stock_syncer/shopify_api_wrapper"
	"pcn_stock_syncer/teams_notifier"
	"time"
)

// TODO: Take a look at all the endpoints again and double check if everything is done correctly.
func HandleSyncStock() error {

	ShopifyProducts, err := shopify_api_wrapper.ShopifyApiGetProducts_AllPages(shopify_api_wrapper.ShopifyApiQueryParams{
		Fields: []string{"id, variants"},
		//Ids:    []string{"6748027912399"},
		Status: "active",
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

	successBarcodes := []string{}

	// Iterate over all products from shopify and all variants from each product.
	for _, product := range ShopifyProducts.Body.Products {
		for _, variant := range product.Variants {
			time.Sleep(2 * time.Second)

			isTracked, err := ShopifyApiGetInventoryItem(variant.InventoryItemId)
			if err != nil {

				teams_notifier.SendInventoryItemErrorToTeams(variant.Barcode, product.Id, err)
				return err
				//continue
			}
			if !isTracked {
				// TODO: vi skal lave et eller andet tjek her på om det burde være tracked eller ej.
				continue
			}

			pcnAvailableQuantity, exists := PcnProducts[variant.Barcode]
			if !exists {
				teams_notifier.SendNotInPCNErrorToTeams(variant.Barcode, product.Id)
				continue
			} else {
				if pcnAvailableQuantity == variant.InventoryQuantity {
					continue
				}

				if err := shopify_api_wrapper.SetInventoryLevel(&shopify_api_wrapper.SetInventoryLevelBody{
					Location_id:       ShopifyInventoryId.Body.Locations[0].LocationId,
					Inventory_item_id: variant.InventoryItemId,
					Available:         pcnAvailableQuantity,
				}); err != nil {

					teams_notifier.SendUpdateInventoryLevelErrorToTeams(variant.Barcode, product.Id, pcnAvailableQuantity, err)
					continue
				} else {
					successBarcodes = append(successBarcodes, variant.Barcode)
				}

			}
		}
	}
	teams_notifier.NotifyTeamsSuccesCount(successBarcodes)
	return nil
}
