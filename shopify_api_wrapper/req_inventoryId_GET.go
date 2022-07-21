package shopify_api_wrapper

//TODO: Read up on the documentation for inserting the quantities to find out which fields we need.
type ShopifyApiGetInventoryIdResult struct {
	Locations []struct {
		Id int64 `json:"id"`
	} `json:"locations"`
}

type ShopifyApiGetInventoryIdReturn struct {
	Body *ShopifyApiGetInventoryIdResult
}

// TODO: Figure out how to do pagination in Shopify API?
func ShopifyApiGetInventoryId(params ShopifyApiQueryParams) (ShopifyApiGetInventoryIdReturn, error) {
	resp, err := GetShopifyApiBaseClient().
		R().
		SetQueryParams(params.AsReqParams()).
		SetResult(ShopifyApiGetInventoryIdResult{}).
		Get("locations.json")
	if err != nil {
		return ShopifyApiGetInventoryIdReturn{}, err
	}

	return ShopifyApiGetInventoryIdReturn{
		Body: resp.Result().(*ShopifyApiGetInventoryIdResult),
	}, nil
}
