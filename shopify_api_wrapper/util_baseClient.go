package shopify_api_wrapper

import (
	"os"

	"github.com/imroc/req/v3"
)

func GetShopifyApiBaseClient() *req.Client {
	return req.C().SetBaseURL(os.Getenv("SHOPIFY_ADRESS"))
}
