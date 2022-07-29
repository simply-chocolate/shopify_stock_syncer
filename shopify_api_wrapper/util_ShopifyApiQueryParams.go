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

	return queryParams
}
