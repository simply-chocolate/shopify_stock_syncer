package shopify_api_wrapper

type ShopifyApiGetInventoryIdResult struct {
	Locations []struct {
		LocationId int64 `json:"id"`
	} `json:"locations"`
}

type ShopifyApiGetInventoryIdReturn struct {
	Body *ShopifyApiGetInventoryIdResult
}

func ShopifyApiGetInventoryId(params ShopifyApiQueryParams) (ShopifyApiGetInventoryIdReturn, error) {
	resp, err := GetShopifyApiBaseClient().
		R().
		SetQueryParams(params.AsReqParams()).
		SetSuccessResult(ShopifyApiGetInventoryIdResult{}).
		Get("locations.json")
	if err != nil {
		return ShopifyApiGetInventoryIdReturn{}, err
	}

	return ShopifyApiGetInventoryIdReturn{
		Body: resp.SuccessResult().(*ShopifyApiGetInventoryIdResult),
	}, nil
}
