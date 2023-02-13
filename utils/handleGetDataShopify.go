package utils

import (
	"fmt"
	"pcn_stock_syncer/shopify_api_wrapper"
)

// Gets the inventory from the variant id.
func ShopifyApiGetInventoryItem(variantInventoryItemId int) (bool, error) {
	ShopifyInventoryItem, err := shopify_api_wrapper.ShopifyApiInventoryItem(map[string]string{
		"ids": fmt.Sprint(variantInventoryItemId),
	})
	if err != nil {
		return false, err
	}
	return ShopifyInventoryItem.Body.InventoryItems[0].IsTracked, nil
}
