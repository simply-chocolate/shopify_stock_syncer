package shopify_api_wrapper

//TODO: Read up on the documentation for inserting the quantities to find out which fields we need.
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
	Body *ShopifyApiGetInventoryIdResult
}

// TODO: Figure out how to do pagination in Shopify API?
func ShopifyApiGetProducts(params ShopifyApiQueryParams) (ShopifyApiGetInventoryIdReturn, error) {
	resp, err := GetShopifyApiBaseClient().
		R().
		SetQueryParams(params.AsReqParams()).
		SetResult(ShopifyApiGetInventoryIdResult{}).
		Get("products.json")
	if err != nil {
		return ShopifyApiGetInventoryIdReturn{}, err
	}

	return ShopifyApiGetInventoryIdReturn{
		Body: resp.Result().(*ShopifyApiGetInventoryIdResult),
	}, nil
}
