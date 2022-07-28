package pcn_api_wrapper

import (
	"os"

	"github.com/imroc/req/v3"
)

// Returns a base client that has already logged in and been authenticated at the base url for the PCN API
func GetPcnApiBaseClient() *req.Client {
	return req.C().
		SetBaseURL(os.Getenv("PCN_ADRESS")).
		SetCommonBasicAuth(os.Getenv("PCN_AUTH_UN"), os.Getenv("PCN_AUTH_PW")).
		SetCommonHeader("Content-Type", "application/json")
}
