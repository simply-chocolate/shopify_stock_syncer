package shopify_api_wrapper

import (
	"strings"
)

type ShopifyApiQueryParams struct {
	// Select which field you want from the API
	Fields []string
	// Select which filters you want to apply on the result
	Filters []string
}

func (p *ShopifyApiQueryParams) AsReqParams() map[string]string {
	queryParams := make(map[string]string)
	if p.Fields != nil {
		queryParams["fields"] = strings.Join(p.Fields, ",")
	}
	if p.Filters != nil {
		queryParams["filters"] = strings.Join(p.Filters, ",")
	}

	return queryParams
}
