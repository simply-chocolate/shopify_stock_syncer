package shopify_api_wrapper

import "fmt"

type ShopifyApiGetProductsResult struct {
	Products []struct {
		Variants []struct {
			Barcode           string `json:"barcode"`
			InventoryItemId   int    `json:"inventory_item_id"`
			InventoryQuantity int    `json:"inventory_quantity"`
		} `json:"variants"`
	} `json:"products"`
}

type ShopifyApiGetProductsReturn struct {
	Body *ShopifyApiGetProductsResult
}

// TODO: Figure out how to do pagination in Shopify API?
// https://shopify.dev/api/usage/pagination-rest
// Right now we can get a maximum of 250 products out.
// We could order them by product ID and then do calls with the filter "since_id" until the call is empty
func ShopifyApiGetProducts(params ShopifyApiQueryParams) (ShopifyApiGetProductsReturn, error) {
	resp, err := GetShopifyApiBaseClient().
		//DevMode().
		R().
		SetQueryParams(params.AsReqParams()).
		SetResult(ShopifyApiGetProductsResult{}).
		Get("products.json")
	if err != nil {
		return ShopifyApiGetProductsReturn{}, err
	}
	fmt.Println(resp.Header["Link"])

	return ShopifyApiGetProductsReturn{
		Body: resp.Result().(*ShopifyApiGetProductsResult),
	}, nil
}
