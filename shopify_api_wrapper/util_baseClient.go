package shopify_api_wrapper

import (
	"fmt"
	"os"

	"github.com/imroc/req/v3"
)

func GetShopifyApiBaseClient() *req.Client {
	return req.C().
		SetBaseURL(fmt.Sprintf("%s/%s", os.Getenv("SHOPIFY_ADDRESS"), os.Getenv("SHOPIFY_API_VERSION"))).
		SetCommonBasicAuth(os.Getenv("SHOPIFY_KEY"), os.Getenv("SHOPIFY_PASS")).
		SetCommonHeader("Content-Type", "application/json")
}
