package shopify_api_wrapper

import (
	"net/http"
	"net/url"
	"strings"
)

type ShopifyApiGetProductsResult struct {
	Products []struct {
		Variants []struct {
			Barcode           string `json:"barcode"`
			Sku               string `json:"sku"`
			InventoryItemId   int    `json:"inventory_item_id"`
			InventoryQuantity int    `json:"inventory_quantity"`
		} `json:"variants"`
	} `json:"products"`
}

type ShopifyApiGetProductsReturn struct {
	Body            *ShopifyApiGetProductsResult
	ResponseHeaders http.Header
}

func ShopifyApiGetProducts(params ShopifyApiQueryParams) (ShopifyApiGetProductsReturn, error) {
	resp, err := GetShopifyApiBaseClient().
		DevMode().
		R().
		SetQueryParams(params.AsReqParams()).
		SetSuccessResult(ShopifyApiGetProductsResult{}).
		Get("products.json")
	if err != nil {
		return ShopifyApiGetProductsReturn{}, err
	}

	return ShopifyApiGetProductsReturn{
		Body:            resp.SuccessResult().(*ShopifyApiGetProductsResult),
		ResponseHeaders: resp.Header,
	}, nil
}

func ShopifyApiGetProducts_AllPages(params ShopifyApiQueryParams) (ShopifyApiGetProductsReturn, error) {
	params.Limit = 250
	res := ShopifyApiGetProductsResult{}
	var nextLink *url.URL
	for page := 0; ; page++ {
		if nextLink != nil {
			qp, _ := url.ParseQuery(nextLink.RawQuery)
			params.PageInfo = qp.Get("page_info")
		}

		getItemsRes, err := ShopifyApiGetProducts(params)
		if err != nil {
			return ShopifyApiGetProductsReturn{}, err
		}

		res.Products = append(res.Products, getItemsRes.Body.Products...)

		if getItemsRes.ResponseHeaders.Get("Link") == "" {
			break
		}

		linkHeaderRaw := getItemsRes.ResponseHeaders.Get("Link")
		linkHeaderEntries := strings.Split(linkHeaderRaw, ", ")
		for _, linkHeaderEntry := range linkHeaderEntries {
			linkHeaderEntryParts := strings.Split(linkHeaderEntry, "; ")
			link := linkHeaderEntryParts[0]
			rel := linkHeaderEntryParts[1]

			if rel == "rel=\"next\"" {
				nextLink, _ = url.Parse(link[1 : len(link)-1])
				break
			}
		}

		if nextLink == nil {
			break
		}
	}

	return ShopifyApiGetProductsReturn{
		Body: &res,
	}, nil
}
