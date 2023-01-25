package shopify_api_wrapper

import (
	"strconv"
	"strings"
)

type ShopifyApiQueryParams struct {
	// Select which field you want from the API
	Fields []string
	// Select which filters you want to apply on the result
	Filters []string
	// Sets the limit for the page
	Limit int
	// Holds the predefined pageinfo from Shopify
	PageInfo string
	// Checks the status of the product (active, draft)
	Status string
	// Returns only the products with the given ids
	Ids []string
}

func (p *ShopifyApiQueryParams) AsReqParams() map[string]string {
	queryParams := make(map[string]string)
	if p.Fields != nil {
		queryParams["fields"] = strings.Join(p.Fields, ",")
	}
	if p.Filters != nil {
		queryParams["filters"] = strings.Join(p.Filters, "&")
	}
	if p.Limit != 0 {
		queryParams["limit"] = strconv.Itoa(p.Limit)
	}
	if p.PageInfo != "" {
		queryParams["page_info"] = p.PageInfo
	}
	if p.Status != "" {
		queryParams["status"] = p.Status
	}
	if p.Ids != nil {
		queryParams["ids"] = strings.Join(p.Ids, ",")
	}

	return queryParams
}
