package shopify_api_wrapper


//TODO: Read up on the documentation for inserting the quantities to find out which fields we need.
type ShopifyApiGetProductsResult {
	Products []struct{
		Variants []struct{
			Barcode string `json:"barcode"`
			InventoryItemId string `json:"inventory_item_id"`
			InventoryQuantity int `json:"inventory_quantity"`
		}`json:"variants"`
	} `json:"products"`
}


func ShopifyApiGetProducts(params ShopifyApiQueryParams) {
	resp, err := GetShopifyApiBaseClient().
		R().
		SetQueryParams(params.AsReqParams()).
		Get("products")
}
