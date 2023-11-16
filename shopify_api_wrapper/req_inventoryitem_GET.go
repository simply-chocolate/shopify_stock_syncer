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

// TODO: Handle the error that this sometimes throws. Most likely its when the product is not tracked.
/*
panic: interface conversion: interface {} is nil, not *shopify_api_wrapper.ShopifyApiInventoryItemResult

goroutine 1318 [running]:
pcn_stock_syncer/shopify_api_wrapper.ShopifyApiInventoryItem(0x8514a0?)
        C:/Projects/shopify_stock_syncer/shopify_api_wrapper/req_inventoryitem_GET.go:35 +0x525
pcn_stock_syncer/utils.ShopifyApiGetInventoryItem({0xc000025ad0, 0xe})
        C:/Projects/shopify_stock_syncer/utils/handleGetDataShopify.go:11 +0x125
pcn_stock_syncer/utils.HandleSyncStock()
        C:/Projects/shopify_stock_syncer/utils/HandleSyncStock.go:42 +0x370
main.main.func1()
        C:/Projects/shopify_stock_syncer/main.go:35 +0x8a
reflect.Value.call({0x835660?, 0xc0003c5bf0?, 0x488653?}, {0x8c854a, 0x4}, {0xd31288, 0x0, 0x489514?})
        C:/Users/ChristianKHasselstee/scoop/apps/go/current/src/reflect/value.go:584 +0x8c5
reflect.Value.Call({0x835660?, 0xc0003c5bf0?, 0x60?}, {0xd31288?, 0x8?, 0x294e1ce0108?})
        C:/Users/ChristianKHasselstee/scoop/apps/go/current/src/reflect/value.go:368 +0xbc
github.com/go-co-op/gocron.callJobFuncWithParams({0x835660?, 0xc0003c5bf0?}, {0x0, 0x0, 0xc0007cabe8?})
        C:/Users/ChristianKHasselstee/go/pkg/mod/github.com/go-co-op/gocron@v1.30.1/gocron.go:116 +0x1bb
github.com/go-co-op/gocron.runJob({0xc00069bc20, {{0x0, 0x0}, {0x0, 0x0}, 0x0, 0x0, 0x0, 0x0}, {0x835660, ...}, ...})
        C:/Users/ChristianKHasselstee/go/pkg/mod/github.com/go-co-op/gocron@v1.30.1/executor.go:77 +0xe5
github.com/go-co-op/gocron.(*jobFunction).singletonRunner(0xc000818280)
        C:/Users/ChristianKHasselstee/go/pkg/mod/github.com/go-co-op/gocron@v1.30.1/executor.go:106 +0x1b8
created by github.com/go-co-op/gocron.(*executor).runJob
        C:/Users/ChristianKHasselstee/go/pkg/mod/github.com/go-co-op/gocron@v1.30.1/executor.go:180 +0x2b0
exit status 2
*/

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
