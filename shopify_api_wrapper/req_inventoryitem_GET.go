package shopify_api_wrapper

import (
	"encoding/json"
	"fmt"
)

type ShopifyApiInventoryItemResult struct {
	InventoryItems []struct {
		InventoryItemId json.Number `json:"id"`
		Barcode         string      `json:"sku"`
		IsTracked       bool        `json:"tracked"`
	} `json:"inventory_items"`
}

type ShopifyApiInventoryItemReturn struct {
	Body *ShopifyApiInventoryItemResult
}

func ShopifyApiInventoryItem(params map[string]string) (ShopifyApiInventoryItemReturn, error) {

	resp, err := GetShopifyApiBaseClient().
		R().
		SetQueryParams(params).
		SetSuccessResult(ShopifyApiInventoryItemResult{}).
		Get("inventory_items.json")
	if err != nil {
		return ShopifyApiInventoryItemReturn{}, err
	}
	if resp == nil {
		return ShopifyApiInventoryItemReturn{}, fmt.Errorf("resp is nil")
	}

	return ShopifyApiInventoryItemReturn{
		Body: resp.SuccessResult().(*ShopifyApiInventoryItemResult),
	}, nil

}

// Calls the function ShopifyApiInventoryItem repeatedly until we have gone through all products
//
// Returns a ShopifyApiInventoryItemReturn containing the data for all the products we've gotten out in our ShopifyApiGetProducts call.
func ShopifyApiInventoryItem_AllItems(ShopifyProducts ShopifyApiGetProductsReturn) (ShopifyApiInventoryItemReturn, error) {
	res := ShopifyApiInventoryItemResult{} // Create the string og Ids and use I to see if you have reached 100
	ProductIdStringSlice := make(map[int]string)
	ProductIdString := ""
	for _, product := range ShopifyProducts.Body.Products {
		for i, variant := range product.Variants {
			// If the index is 99 it will be the 100th entry. Thus if the index is 100 and 100 / 100 = 1 then the length of ProductIdStringSlice will still be 0 and it will be greater.
			if i/100 > len(ProductIdStringSlice) {
				// Remove last element from the string, which should be a comma
				if len(ProductIdString) > 0 {
					ProductIdString = ProductIdString[:len(ProductIdString)-1]
				}
				// Then we set the index to be equal to our string
				ProductIdStringSlice[i/100] = ProductIdString
				// And lastly we erase the string.
				ProductIdString = ""
			}

			ProductIdString += fmt.Sprintf("%v,", variant.InventoryItemId)
		}
	}

	// If we have less than 100 elements, we need to make sure the map is still used.
	if len(ProductIdStringSlice) == 0 {
		if len(ProductIdString) > 0 {
			ProductIdString = ProductIdString[:len(ProductIdString)-1]
		}
		ProductIdStringSlice[0] = ProductIdString
	}

	for _, productIdStringValue := range ProductIdStringSlice {

		ShopifyInventoryItemsRes, err := ShopifyApiInventoryItem(map[string]string{
			"ids": productIdStringValue,
		})
		if err != nil {
			return ShopifyApiInventoryItemReturn{}, err
		}

		res.InventoryItems = append(res.InventoryItems, ShopifyInventoryItemsRes.Body.InventoryItems...)
	}
	return ShopifyApiInventoryItemReturn{
		Body: &res,
	}, nil
}
