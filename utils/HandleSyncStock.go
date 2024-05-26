package utils

import (
	"fmt"
	"pcn_stock_syncer/shopify_api_wrapper"
	"pcn_stock_syncer/teams_notifier"

	"time"
)

// TODO: Take a look at all the endpoints again and double check if everything is done correctly.
func HandleSyncStock() error {

	ShopifyProducts, err := shopify_api_wrapper.ShopifyApiGetProducts_AllPages(shopify_api_wrapper.ShopifyApiQueryParams{
		Fields: []string{"id, title, variants"},
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

	successProducts := []teams_notifier.ProductAmounts{}

	// Iterate over all products from shopify and all variants from each product.
	for _, product := range ShopifyProducts.Body.Products {
		fmt.Println("Product: ", product.Id)
		for _, variant := range product.Variants {
			time.Sleep(2 * time.Second)

			isTracked, err := ShopifyApiGetInventoryItem(variant.InventoryItemId)
			if err != nil {

				teams_notifier.SendInventoryItemErrorToTeams(variant.Barcode, product.Id, err)
				return err
				//continue
			}
			if !isTracked {
				DevLog(fmt.Sprintf("Barcode: %v not tracked in Shopify: ", variant.Barcode))
				// TODO: vi skal lave et eller andet tjek her på om det burde være tracked eller ej.
				continue
			}

			pcnAvailableQuantity, exists := PcnProducts[variant.Barcode]
			if !exists {
				DevLog(fmt.Sprintf("Barcode: %v not found in PCN: ", variant.Barcode))
				continue

			} else {
				if pcnAvailableQuantity == variant.InventoryQuantity {
					DevLog(fmt.Sprintf("Barcode %v is already up to date", variant.Barcode))
					continue

				}
				if err := shopify_api_wrapper.SetInventoryLevel(&shopify_api_wrapper.SetInventoryLevelBody{
					Location_id:       ShopifyInventoryId.Body.Locations[0].LocationId,
					Inventory_item_id: variant.InventoryItemId,
					Available:         pcnAvailableQuantity,
				}); err != nil {
					DevLog(fmt.Sprintf("Error updating inventory level for barcode %v: %v", variant.Barcode, err))
					teams_notifier.SendUpdateInventoryLevelErrorToTeams(variant.Barcode, product.Id, pcnAvailableQuantity, err)
					continue

				} else {
					DevLog(fmt.Sprintf("Barcode %v updated to %v", variant.Barcode, pcnAvailableQuantity))
					var productAmounts teams_notifier.ProductAmounts
					productAmounts.Barcode = variant.Barcode
					productAmounts.ProductName = product.ProductName
					productAmounts.QuantityPCN = pcnAvailableQuantity
					productAmounts.QuantityShopify = variant.InventoryQuantity
					successProducts = append(successProducts, productAmounts)
				}

			}
		}
	}
	teams_notifier.NotifyTeamsSuccesCount(successProducts)
	return nil
}
