package utils

import (
	"encoding/json"
	"pcn_stock_syncer/shopify_api_wrapper"
)

// Gets the inventory from the variant id.
func ShopifyApiGetInventoryItem(variantInventoryItemId json.Number) (bool, error) {

	ShopifyInventoryItem, err := shopify_api_wrapper.ShopifyApiInventoryItem(variantInventoryItemId)
	if err != nil {
		return false, err
	}

	return ShopifyInventoryItem.Body.InventoryItem.IsTracked, nil
}
