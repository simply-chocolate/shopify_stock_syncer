package pcn_api_wrapper

import (
	"errors"
	"fmt"
	"os"
)

type PcnApiBundleProductListResult struct {
	Results []struct {
		BundleCode  string                     `json:"bundle_articleno"`
		BundleLines []PcnApiBundleProductLines `json:"bundlelines"`
	} `json:"results"`
}

type PcnApiBundleProductLines struct {
	BarCode  string  `json:"articleno"`
	Quantity float64 `json:"amount"`
}

type PcnApiBundleProductListReturn struct {
	Body       *PcnApiBundleProductListResult
	StatusCode int
}

// Calls the PCN API at endpoint stocklist to return Barcodes and available units in stock.
func PcnApiBundleProductList() (PcnApiBundleProductListReturn, int, error) {
	resp, err := GetPcnApiBaseClient().
		//DevMode().
		R().
		SetSuccessResult(PcnApiBundleProductListResult{}).
		SetBody(map[string]interface{}{
			"cid":     os.Getenv("PCN_CID"),
			"olsuser": os.Getenv("PCN_OLSUSER"),
			"olspass": os.Getenv("PCN_OLSPASS"),
		}).Post("bundleproductlist")
	if err != nil {
		fmt.Println(err)
		return PcnApiBundleProductListReturn{}, resp.StatusCode, errors.New("error getting the stock data from Pcn API")
	}
	if resp.IsError() {
		return PcnApiBundleProductListReturn{}, resp.StatusCode, fmt.Errorf("error contacting PCN API. Status code: %v, errorMsg: %v ", resp.StatusCode, resp.Response)
	}

	return PcnApiBundleProductListReturn{
		Body: resp.SuccessResult().(*PcnApiBundleProductListResult),
	}, resp.StatusCode, nil

}
